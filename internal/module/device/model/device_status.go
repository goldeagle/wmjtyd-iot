package model

import (
	"time"

	"gorm.io/gorm"
)

// DeviceStatus 设备状态信息
type DeviceStatus struct {
	ID        uint      `gorm:"primaryKey;comment:ID"`
	DeviceID  int       `gorm:"not null;comment:设备"`
	Status    int       `gorm:"not null;comment:状态:0-离线,1-在线,2-故障,3-维护"`
	Ts        int       `gorm:"not null;comment:状态时间"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

// TableName 返回表名
func (DeviceStatus) TableName() string {
	return "sac_device_status"
}

// DeviceStatusRepo 设备状态信息仓库
type DeviceStatusRepo struct {
	db *gorm.DB
}

// NewDeviceStatusRepo 创建新的设备状态信息仓库
func NewDeviceStatusRepo(db *gorm.DB) *DeviceStatusRepo {
	return &DeviceStatusRepo{db: db}
}

// Create 创建记录
func (r *DeviceStatusRepo) Create(d *DeviceStatus) error {
	return r.db.Create(d).Error
}

// Update 更新记录
func (r *DeviceStatusRepo) Update(d *DeviceStatus) error {
	return r.db.Save(d).Error
}

// Delete 删除记录
func (r *DeviceStatusRepo) Delete(d *DeviceStatus) error {
	return r.db.Delete(d).Error
}

// GetByID 根据ID获取记录
func (r *DeviceStatusRepo) GetByID(id uint) (*DeviceStatus, error) {
	var d DeviceStatus
	err := r.db.First(&d, id).Error
	return &d, err
}

// ListByDeviceID 根据设备ID获取记录列表
func (r *DeviceStatusRepo) ListByDeviceID(deviceID int) ([]DeviceStatus, error) {
	var statuses []DeviceStatus
	err := r.db.Where("device_id = ?", deviceID).Find(&statuses).Error
	return statuses, err
}
