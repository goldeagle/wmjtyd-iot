package model

import (
	"gorm.io/gorm"
)

// File 文件表
// @Description 文件表
type File struct {
	gorm.Model
	FileID      int    `gorm:"primaryKey;column:file_id" json:"file_id"` // 文件ID
	FileName    string `gorm:"column:file_name" json:"file_name"`        // 文件名
	FilePath    string `gorm:"column:file_path" json:"file_path"`        // 文件路径
	FileSize    int64  `gorm:"column:file_size" json:"file_size"`        // 文件大小
	FileType    string `gorm:"column:file_type" json:"file_type"`        // 文件类型
	UploadTime  int64  `gorm:"column:upload_time" json:"upload_time"`    // 上传时间
	UploaderID  int    `gorm:"column:uploader_id" json:"uploader_id"`    // 上传者ID
	Status      int    `gorm:"column:status" json:"status"`              // 文件状态
	Description string `gorm:"column:description" json:"description"`    // 文件描述
}

// TableName 指定表名
func (File) TableName() string {
	return "cd_file"
}

// Create 创建文件记录
func (f *File) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

// Update 更新文件记录
func (f *File) Update(db *gorm.DB) error {
	return db.Save(f).Error
}

// Delete 删除文件记录
func (f *File) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

// GetByID 根据ID获取文件记录
func (f *File) GetByID(db *gorm.DB, id int) error {
	return db.First(f, id).Error
}

// List 获取文件列表
func (f *File) List(db *gorm.DB, page, pageSize int) ([]File, int64, error) {
	var files []File
	var count int64

	if err := db.Model(f).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&files).Error; err != nil {
		return nil, 0, err
	}

	return files, count, nil
}
