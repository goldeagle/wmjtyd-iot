package model

import (
	"time"

	"gorm.io/gorm"
)

// RemoteJobs 远程任务信息
// @Description 远程任务信息
type RemoteJobs struct {
	gorm.Model
	JobID       string    `gorm:"type:varchar(50);not null;comment:任务ID" json:"job_id"`
	DeviceID    string    `gorm:"type:varchar(50);not null;comment:设备ID" json:"device_id"`
	JobType     string    `gorm:"type:varchar(50);comment:任务类型" json:"job_type"`
	Status      int       `gorm:"type:tinyint(4);default:1;comment:状态" json:"status"`
	StartTime   time.Time `gorm:"comment:开始时间" json:"start_time"`
	EndTime     time.Time `gorm:"comment:结束时间" json:"end_time"`
	CreatedTime int       `gorm:"comment:创建时间" json:"created_time"`
	UpdatedTime int       `gorm:"comment:更新时间" json:"updated_time"`
}

// Create 创建远程任务信息
func (r *RemoteJobs) Create(db *gorm.DB) error {
	return db.Create(r).Error
}

// GetByID 根据ID获取远程任务信息
func (r *RemoteJobs) GetByID(db *gorm.DB, id uint) error {
	return db.First(r, id).Error
}

// Update 更新远程任务信息
func (r *RemoteJobs) Update(db *gorm.DB) error {
	return db.Save(r).Error
}

// Delete 删除远程任务信息
func (r *RemoteJobs) Delete(db *gorm.DB) error {
	return db.Delete(r).Error
}

// List 获取远程任务信息列表
func (r *RemoteJobs) List(db *gorm.DB, page, pageSize int) ([]RemoteJobs, int64, error) {
	var remoteJobs []RemoteJobs
	var count int64

	err := db.Model(r).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&remoteJobs).Error
	if err != nil {
		return nil, 0, err
	}

	return remoteJobs, count, nil
}
