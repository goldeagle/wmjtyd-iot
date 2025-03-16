package model

import (
	"gorm.io/gorm"
)

// DeviceFirmware 设备固件模型
// @Description 设备固件模型
type DeviceFirmware struct {
	gorm.Model
	Caption    string `gorm:"type:varchar(100);comment:固件说明" json:"caption"`
	ModelIDs   string `gorm:"type:text;comment:适配设备型号" json:"model_ids"`
	Version    string `gorm:"type:varchar(50);comment:固件版本" json:"version"`
	URL        string `gorm:"type:varchar(250);comment:下载地址" json:"url"`
	CreateTime int64  `gorm:"comment:上传时间" json:"create_time"`
}

// Create 创建固件记录
func (f *DeviceFirmware) Create(db *gorm.DB) error {
	return db.Create(f).Error
}

// GetByID 根据ID获取固件记录
func (f *DeviceFirmware) GetByID(db *gorm.DB, id uint) error {
	return db.First(f, id).Error
}

// Update 更新固件记录
func (f *DeviceFirmware) Update(db *gorm.DB) error {
	return db.Save(f).Error
}

// Delete 删除固件记录
func (f *DeviceFirmware) Delete(db *gorm.DB) error {
	return db.Delete(f).Error
}

// List 获取固件列表
func (f *DeviceFirmware) List(db *gorm.DB, page, pageSize int) ([]DeviceFirmware, int64, error) {
	var DeviceFirmwares []DeviceFirmware
	var count int64

	err := db.Model(f).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&DeviceFirmwares).Error
	if err != nil {
		return nil, 0, err
	}

	return DeviceFirmwares, count, nil
}
