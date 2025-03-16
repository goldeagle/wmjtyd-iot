package model

import (
	"gorm.io/gorm"
)

// DeviceLog 设备日志模型
// @Description 设备日志模型
type DeviceLog struct {
	gorm.Model
	DeviceID  uint   `gorm:"comment:设备ID" json:"device_id"`
	LogType   string `gorm:"type:varchar(50);comment:日志类型" json:"log_type"`
	LogLevel  string `gorm:"type:varchar(20);comment:日志级别" json:"log_level"`
	Content   string `gorm:"type:text;comment:日志内容" json:"content"`
	Timestamp int64  `gorm:"comment:日志时间戳" json:"timestamp"`
}

// Create 创建设备日志记录
func (l *DeviceLog) Create(db *gorm.DB) error {
	return db.Create(l).Error
}

// GetByID 根据ID获取设备日志记录
func (l *DeviceLog) GetByID(db *gorm.DB, id uint) error {
	return db.First(l, id).Error
}

// Update 更新设备日志记录
func (l *DeviceLog) Update(db *gorm.DB) error {
	return db.Save(l).Error
}

// Delete 删除设备日志记录
func (l *DeviceLog) Delete(db *gorm.DB) error {
	return db.Delete(l).Error
}

// List 获取设备日志列表
func (l *DeviceLog) List(db *gorm.DB, page, pageSize int) ([]DeviceLog, int64, error) {
	var logs []DeviceLog
	var count int64

	err := db.Model(l).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&logs).Error
	if err != nil {
		return nil, 0, err
	}

	return logs, count, nil
}
