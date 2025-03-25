package model

import (
	"gorm.io/gorm"
)

// DeviceInfo 设备信息模型
// @Description 设备信息模型
type DeviceInfo struct {
	gorm.Model
	UUID         string `gorm:"type:char(36);not null;comment:UUID" json:"uuid"`
	Serial       string `gorm:"type:varchar(50);comment:SN序列号" json:"serial"`
	Name         string `gorm:"type:varchar(50);comment:名称" json:"name"`
	Alias        string `gorm:"type:varchar(100);comment:别名" json:"alias"`
	Type         string `gorm:"type:enum('GW','HMI','PLC','CAMERA','SENSOR','MISC');default:'HMI';comment:设备类型" json:"type"`
	State        int    `gorm:"type:tinyint(4);default:1;comment:状态" json:"state"`
	MqttClientID string `gorm:"type:varchar(50);comment:MQTT用户ID" json:"mqtt_client_id"`
	MqttPassword string `gorm:"type:varchar(50);comment:MQTT用户密码" json:"mqtt_password"`
	CreatedTime  int    `gorm:"comment:创建时间" json:"created_time"`
	UpdatedTime  int    `gorm:"comment:更新时间" json:"updated_time"`
	ModelID      int    `gorm:"comment:设备型号" json:"model_id"`
	LocationID   int    `gorm:"comment:投放地址" json:"location_id"`
	UserID       int    `gorm:"comment:绑定用户的id" json:"user_id"`
	GroupID      int    `gorm:"comment:绑定用户的分组" json:"group_id"`
	OnlineTime   int    `gorm:"comment:上线时间" json:"online_time"`
	OfflineTime  int    `gorm:"comment:下线时间" json:"offline_time"`
}

// Create 创建设备信息
func (i *DeviceInfo) Create(db *gorm.DB) error {
	return db.Create(i).Error
}

// GetByID 根据ID获取设备信息
func (i *DeviceInfo) GetByID(db *gorm.DB, id uint) error {
	return db.First(i, id).Error
}

// Update 更新设备信息
func (i *DeviceInfo) Update(db *gorm.DB) error {
	return db.Save(i).Error
}

// Delete 删除设备信息
func (i *DeviceInfo) Delete(db *gorm.DB) error {
	return db.Delete(i).Error
}

// List 获取设备信息列表
func (i *DeviceInfo) List(db *gorm.DB, page, pageSize int) ([]DeviceInfo, int64, error) {
	var DeviceInfos []DeviceInfo
	var count int64

	err := db.Model(i).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&DeviceInfos).Error
	if err != nil {
		return nil, 0, err
	}

	return DeviceInfos, count, nil
}
