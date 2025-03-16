package model

import (
	"gorm.io/gorm"
)

// Log 日志表
// @Description 日志表
type Log struct {
	gorm.Model
	LogID       int    `gorm:"primaryKey;column:log_id" json:"log_id"` // 日志ID
	UserID      int    `gorm:"column:user_id" json:"user_id"`          // 用户ID
	UserName    string `gorm:"column:user_name" json:"user_name"`      // 用户名
	Module      string `gorm:"column:module" json:"module"`            // 操作模块
	Action      string `gorm:"column:action" json:"action"`            // 操作类型
	Description string `gorm:"column:description" json:"description"`  // 操作描述
	IP          string `gorm:"column:ip" json:"ip"`                    // 操作IP
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`  // 操作时间
	Status      int    `gorm:"column:status" json:"status"`            // 操作状态
}

// TableName 指定表名
func (Log) TableName() string {
	return "cd_log"
}

// Create 创建日志记录
func (l *Log) Create(db *gorm.DB) error {
	return db.Create(l).Error
}

// Update 更新日志记录
func (l *Log) Update(db *gorm.DB) error {
	return db.Save(l).Error
}

// Delete 删除日志记录
func (l *Log) Delete(db *gorm.DB) error {
	return db.Delete(l).Error
}

// GetByID 根据ID获取日志记录
func (l *Log) GetByID(db *gorm.DB, id int) error {
	return db.First(l, id).Error
}

// List 获取日志列表
func (l *Log) List(db *gorm.DB, page, pageSize int) ([]Log, int64, error) {
	var logs []Log
	var count int64

	if err := db.Model(l).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&logs).Error; err != nil {
		return nil, 0, err
	}

	return logs, count, nil
}
