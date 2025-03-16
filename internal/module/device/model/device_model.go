package model

import (
	"gorm.io/gorm"
)

// DeviceModel 设备型号模型
// @Description 设备型号模型
type DeviceModel struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100);comment:型号名称" json:"name"`
	Manufacturer string `gorm:"type:varchar(100);comment:生产厂家" json:"manufacturer"`
	Description  string `gorm:"type:varchar(250);comment:型号描述" json:"description"`
	Protocol     string `gorm:"type:varchar(50);comment:通信协议" json:"protocol"`
	FirmwareID   uint   `gorm:"comment:默认固件ID" json:"firmware_id"`
}

// Create 创建设备型号记录
func (m *DeviceModel) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

// GetByID 根据ID获取设备型号记录
func (m *DeviceModel) GetByID(db *gorm.DB, id uint) error {
	return db.First(m, id).Error
}

// Update 更新设备型号记录
func (m *DeviceModel) Update(db *gorm.DB) error {
	return db.Save(m).Error
}

// Delete 删除设备型号记录
func (m *DeviceModel) Delete(db *gorm.DB) error {
	return db.Delete(m).Error
}

// List 获取设备型号列表
func (m *DeviceModel) List(db *gorm.DB, page, pageSize int) ([]DeviceModel, int64, error) {
	var models []DeviceModel
	var count int64

	err := db.Model(m).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&models).Error
	if err != nil {
		return nil, 0, err
	}

	return models, count, nil
}
