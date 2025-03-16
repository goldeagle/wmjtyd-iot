package model

import (
	"time"

	"gorm.io/gorm"
)

// DeviceNet 设备网络信息
type DeviceNet struct {
	ID         uint      `gorm:"primaryKey;comment:ID"`
	DeviceID   int       `gorm:"not null;comment:所属设备"`
	McuID      string    `gorm:"size:64;not null;comment:设备ID/IMEI"`
	PositionID int       `gorm:"comment:设备当前位置id，从1~12"`
	Ts         int       `gorm:"not null;comment:上报时间"`
	State      int       `gorm:"comment:当前设备状态：0-无应答（备用）；1-正常工作；2-有警报；3-有故障；4-无法工作"`
	IP         string    `gorm:"size:32;comment:当前连接网络IP"`
	Mask       string    `gorm:"size:32;comment:当前连接网络掩码"`
	Gw         string    `gorm:"size:64;comment:当前连接网关"`
	Mac        string    `gorm:"size:64;comment:当前连接所用网络设备的MAC地址"`
	CreatedAt  time.Time `gorm:"autoCreateTime;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime;comment:更新时间"`
}

// TableName 返回表名
func (DeviceNet) TableName() string {
	return "sac_device_net"
}

// DeviceNetRepo 设备网络信息仓库
type DeviceNetRepo struct {
	db *gorm.DB
}

// NewDeviceNetRepo 创建新的设备网络信息仓库
func NewDeviceNetRepo(db *gorm.DB) *DeviceNetRepo {
	return &DeviceNetRepo{db: db}
}

// Create 创建记录
func (r *DeviceNetRepo) Create(d *DeviceNet) error {
	return r.db.Create(d).Error
}

// Update 更新记录
func (r *DeviceNetRepo) Update(d *DeviceNet) error {
	return r.db.Save(d).Error
}

// Delete 删除记录
func (r *DeviceNetRepo) Delete(d *DeviceNet) error {
	return r.db.Delete(d).Error
}

// GetByID 根据ID获取记录
func (r *DeviceNetRepo) GetByID(id uint) (*DeviceNet, error) {
	var d DeviceNet
	err := r.db.First(&d, id).Error
	return &d, err
}

// ListByDeviceID 根据设备ID获取记录列表
func (r *DeviceNetRepo) ListByDeviceID(deviceID int) ([]DeviceNet, error) {
	var nets []DeviceNet
	err := r.db.Where("device_id = ?", deviceID).Find(&nets).Error
	return nets, err
}
