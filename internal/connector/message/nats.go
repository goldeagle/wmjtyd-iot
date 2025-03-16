package message

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
)

// NATSConfig 配置参数
type NATSConfig struct {
	URL         string
	ClusterID   string
	ClientID    string
	ConnectWait time.Duration
}

// NATSClient NATS客户端
type NATSClient struct {
	conn      *nats.Conn
	stream    stan.Conn
	config    NATSConfig
	logger    *zap.Logger
	mu        sync.RWMutex
	connected bool
}

// NewNATSClient 创建新的NATS客户端
func NewNATSClient(config NATSConfig, logger *zap.Logger) *NATSClient {
	return &NATSClient{
		config: config,
		logger: logger,
	}
}

// Connect 连接到NATS服务器
func (c *NATSClient) Connect(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.connected {
		return nil
	}

	// 连接NATS
	opts := []nats.Option{
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			c.logger.Warn("NATS disconnected", zap.Error(err))
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			c.logger.Info("NATS reconnected", zap.String("url", nc.ConnectedUrl()))
		}),
	}

	conn, err := nats.Connect(c.config.URL, opts...)
	if err != nil {
		return fmt.Errorf("failed to connect to NATS: %w", err)
	}

	// 连接NATS Streaming
	stream, err := stan.Connect(
		c.config.ClusterID,
		c.config.ClientID,
		stan.NatsConn(conn),
		stan.ConnectWait(c.config.ConnectWait),
	)
	if err != nil {
		conn.Close()
		return fmt.Errorf("failed to connect to NATS Streaming: %w", err)
	}

	c.conn = conn
	c.stream = stream
	c.connected = true

	c.logger.Info("Connected to NATS",
		zap.String("url", c.config.URL),
		zap.String("clusterID", c.config.ClusterID),
	)

	return nil
}

// Publish 发布消息
func (c *NATSClient) Publish(subject string, data []byte) error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.connected {
		return errors.New("not connected to NATS")
	}

	return c.stream.Publish(subject, data)
}

// Subscribe 订阅消息
func (c *NATSClient) Subscribe(subject string, handler stan.MsgHandler) (stan.Subscription, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if !c.connected {
		return nil, errors.New("not connected to NATS")
	}

	return c.stream.Subscribe(subject, handler)
}

// Close 关闭连接
func (c *NATSClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if !c.connected {
		return nil
	}

	var errs []error

	// 先关闭 Streaming 连接
	if c.stream != nil {
		if err := c.stream.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close NATS Streaming connection: %w", err))
		}
	}

	// 再关闭 NATS 连接
	if c.conn != nil {
		c.conn.Close() // NATS Close() 方法不返回错误
		c.conn = nil
	}

	c.connected = false
	c.logger.Info("Closed NATS connection")

	// 返回所有错误
	if len(errs) > 0 {
		return fmt.Errorf("errors occurred while closing connections: %v", errs)
	}

	return nil
}

// IsConnected 检查是否已连接
func (c *NATSClient) IsConnected() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.connected
}
