package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

const (
	defaultEndpoint  = "localhost:9000"
	defaultAccessKey = "minioadmin"
	defaultSecretKey = "minioadmin"
	defaultUseSSL    = false
)

type MinioClient struct {
	client *minio.Client
	logger *zap.Logger
}

func NewMinioClient(logger *zap.Logger) (*MinioClient, error) {
	client, err := minio.New(defaultEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(defaultAccessKey, defaultSecretKey, ""),
		Secure: defaultUseSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create MinIO client: %w", err)
	}

	return &MinioClient{
		client: client,
		logger: logger,
	}, nil
}

func (m *MinioClient) CreateBucket(ctx context.Context, bucketName string) error {
	exists, err := m.client.BucketExists(ctx, bucketName)
	if err != nil {
		return fmt.Errorf("failed to check bucket existence: %w", err)
	}

	if exists {
		return nil
	}

	err = m.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return fmt.Errorf("failed to create bucket: %w", err)
	}

	m.logger.Info("Bucket created successfully", zap.String("bucket", bucketName))
	return nil
}

func (m *MinioClient) UploadFile(ctx context.Context, bucketName, objectName string, file io.Reader, size int64) error {
	_, err := m.client.PutObject(ctx, bucketName, objectName, file, size, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to upload file: %w", err)
	}

	m.logger.Info("File uploaded successfully",
		zap.String("bucket", bucketName),
		zap.String("object", objectName))
	return nil
}

func (m *MinioClient) DownloadFile(ctx context.Context, bucketName, objectName string) (io.Reader, error) {
	object, err := m.client.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to download file: %w", err)
	}

	return object, nil
}

func (m *MinioClient) DeleteFile(ctx context.Context, bucketName, objectName string) error {
	err := m.client.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	m.logger.Info("File deleted successfully",
		zap.String("bucket", bucketName),
		zap.String("object", objectName))
	return nil
}

func (m *MinioClient) ListObjects(ctx context.Context, bucketName string, prefix string) ([]minio.ObjectInfo, error) {
	objectCh := m.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix:    prefix,
		Recursive: true,
	})

	var objects []minio.ObjectInfo
	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("error listing objects: %w", object.Err)
		}
		objects = append(objects, object)
	}

	return objects, nil
}

func (m *MinioClient) PresignedGetObject(ctx context.Context, bucketName, objectName string, expires time.Duration) (string, error) {
	url, err := m.client.PresignedGetObject(ctx, bucketName, objectName, expires, nil)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url.String(), nil
}

func (m *MinioClient) HealthCheck(ctx context.Context) error {
	_, err := m.client.ListBuckets(ctx)
	if err != nil {
		return fmt.Errorf("MinIO health check failed: %w", err)
	}
	return nil
}
