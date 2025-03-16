package model

import (
	"time"

	"gorm.io/gorm"
)

// DevicePosition 设备位置信息
type DevicePosition struct {
	ID        uint      `gorm:"primaryKey;comment:ID"`
	DeviceID  int       `gorm:"not null;comment:设备"`
	Name      string    `gorm:"size:64;not null;comment:位置名称"`
	Longitude float64   `gorm:"not null;comment:经度"`
	Latitude  float64   `gorm:"not null;comment:纬度"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

// TableName 返回表名
func (DevicePosition) TableName() string {
	return "sac_device_position"
}

// DevicePositionRepo 设备位置信息仓库
type DevicePositionRepo struct {
	db *gorm.DB
}

// NewDevicePositionRepo 创建新的设备位置信息仓库
func NewDevicePositionRepo(db *gorm.DB) *DevicePositionRepo {
	return &DevicePositionRepo{db: db}
}

// Create 创建记录
func (r *DevicePositionRepo) Create(d *DevicePosition) error {
	return r.db.Create(d).Error
}

// Update 更新记录
func (r *DevicePositionRepo) Update(d *DevicePosition) error {
	return r.db.Save(d).Error
}

// Delete 删除记录
func (r *DevicePositionRepo) Delete(d *DevicePosition) error {
	return r.db.Delete(d).Error
}

// GetByID 根据ID获取记录
func (r *DevicePositionRepo) GetByID(id uint) (*DevicePosition, error) {
	var d DevicePosition
	err := r.db.First(&d, id).Error
	return &d, err
}

// ListByDeviceID 根据设备ID获取记录列表
func (r *DevicePositionRepo) ListByDeviceID(deviceID int) ([]DevicePosition, error) {
	var positions []DevicePosition
	err := r.db.Where("device_id = ?", deviceID).Find(&positions).Error
	return positions, err
}
