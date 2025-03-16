package model

import (
	"gorm.io/gorm"
)

// DeviceConfig 设备配置模型
// @Description 设备配置模型
type DeviceConfig struct {
	gorm.Model
	DeviceID  uint   `gorm:"comment:设备ID" json:"device_id"`
	ConfigKey string `gorm:"type:varchar(100);comment:配置键" json:"config_key"`
	ConfigVal string `gorm:"type:text;comment:配置值" json:"config_val"`
	Version   string `gorm:"type:varchar(50);comment:配置版本" json:"version"`
	Status    string `gorm:"type:varchar(20);comment:配置状态" json:"status"`
}

// Create 创建设备配置记录
func (c *DeviceConfig) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

// GetByID 根据ID获取设备配置记录
func (c *DeviceConfig) GetByID(db *gorm.DB, id uint) error {
	return db.First(c, id).Error
}

// Update 更新设备配置记录
func (c *DeviceConfig) Update(db *gorm.DB) error {
	return db.Save(c).Error
}

// Delete 删除设备配置记录
func (c *DeviceConfig) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}

// List 获取设备配置列表
func (c *DeviceConfig) List(db *gorm.DB, page, pageSize int) ([]DeviceConfig, int64, error) {
	var configs []DeviceConfig
	var count int64

	err := db.Model(c).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&configs).Error
	if err != nil {
		return nil, 0, err
	}

	return configs, count, nil
}

// GetByDeviceID 根据设备ID获取配置
func (c *DeviceConfig) GetByDeviceID(db *gorm.DB, deviceID uint) ([]DeviceConfig, error) {
	var configs []DeviceConfig
	err := db.Where("device_id = ?", deviceID).Find(&configs).Error
	return configs, err
}
