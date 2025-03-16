package model

import (
	"time"

	"gorm.io/gorm"
)

// DeviceWarning 设备告警信息
type DeviceWarning struct {
	ID        uint      `gorm:"primaryKey;comment:ID"`
	DeviceID  int       `gorm:"not null;comment:设备"`
	Level     int       `gorm:"not null;comment:告警级别:1-一般,2-严重,3-紧急"`
	Code      string    `gorm:"size:32;not null;comment:告警代码"`
	Message   string    `gorm:"type:text;not null;comment:告警信息"`
	Ts        int       `gorm:"not null;comment:告警时间"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

// TableName 返回表名
func (DeviceWarning) TableName() string {
	return "sac_device_warning"
}

// DeviceWarningRepo 设备告警信息仓库
type DeviceWarningRepo struct {
	db *gorm.DB
}

// NewDeviceWarningRepo 创建新的设备告警信息仓库
func NewDeviceWarningRepo(db *gorm.DB) *DeviceWarningRepo {
	return &DeviceWarningRepo{db: db}
}

// Create 创建记录
func (r *DeviceWarningRepo) Create(d *DeviceWarning) error {
	return r.db.Create(d).Error
}

// Update 更新记录
func (r *DeviceWarningRepo) Update(d *DeviceWarning) error {
	return r.db.Save(d).Error
}

// Delete 删除记录
func (r *DeviceWarningRepo) Delete(d *DeviceWarning) error {
	return r.db.Delete(d).Error
}

// GetByID 根据ID获取记录
func (r *DeviceWarningRepo) GetByID(id uint) (*DeviceWarning, error) {
	var d DeviceWarning
	err := r.db.First(&d, id).Error
	return &d, err
}

// ListByDeviceID 根据设备ID获取记录列表
func (r *DeviceWarningRepo) ListByDeviceID(deviceID int) ([]DeviceWarning, error) {
	var warnings []DeviceWarning
	err := r.db.Where("device_id = ?", deviceID).Find(&warnings).Error
	return warnings, err
}
