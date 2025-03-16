package model

import (
	"time"

	"gorm.io/gorm"
)

// RemoteHeartBeat 远程心跳信息
// @Description 远程心跳信息
type RemoteHeartBeat struct {
	gorm.Model
	DeviceID    string    `gorm:"type:varchar(50);not null;comment:设备ID" json:"device_id"`
	Timestamp   time.Time `gorm:"comment:时间戳" json:"timestamp"`
	Status      int       `gorm:"type:tinyint(4);default:1;comment:状态" json:"status"`
	CreatedTime int       `gorm:"comment:创建时间" json:"created_time"`
	UpdatedTime int       `gorm:"comment:更新时间" json:"updated_time"`
}

// Create 创建远程心跳信息
func (r *RemoteHeartBeat) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

// GetByID 根据ID获取远程心跳信息
func (r *RemoteHeartBeat) GetByID(db *gorm.DB, id uint) error {
	return db.First(r, id).Error
}

// Update 更新远程心跳信息
func (r *RemoteHeartBeat) Update(db *gorm.DB) error {
	return db.Save(r).Error
}

// Delete 删除远程心跳信息
func (r *RemoteHeartBeat) Delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

// List 获取远程心跳信息列表
func (r *RemoteHeartBeat) List(db *gorm.DB, page, pageSize int) ([]RemoteHeartBeat, int64, error) {
	var remoteHeartBeats []RemoteHeartBeat
	var count int64

	err := db.Model(r).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&remoteHeartBeats).Error
	if err != nil {
		return nil, 0, err
	}

	return remoteHeartBeats, count, nil
}
