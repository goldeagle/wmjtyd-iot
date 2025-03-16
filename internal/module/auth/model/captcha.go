package model

import (
	"gorm.io/gorm"
)

// Captcha 验证码表
// @Description 验证码表
type Captcha struct {
	gorm.Model
	ID         int    `gorm:"primaryKey;column:id" json:"id"`        // ID
	Phone      string `gorm:"column:phone" json:"phone"`             // 手机号
	Code       string `gorm:"column:code" json:"code"`               // 验证码
	ExpireTime int    `gorm:"column:expire_time" json:"expire_time"` // 过期时间
	CreateTime int    `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime int    `gorm:"column:update_time" json:"update_time"` // 更新时间
	Status     int    `gorm:"column:status" json:"status"`           // 状态
}

// TableName 指定表名
func (Captcha) TableName() string {
	return "bsd_captcha"
}

// Create 创建验证码
func (c *Captcha) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

// Update 更新验证码
func (c *Captcha) Update(db *gorm.DB) error {
	return db.Save(c).Error
}

// Delete 删除验证码
func (c *Captcha) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}

// GetByID 根据ID获取验证码
func (c *Captcha) GetByID(db *gorm.DB, id int) error {
	return db.First(c, id).Error
}

// List 获取验证码列表
func (c *Captcha) List(db *gorm.DB, page, pageSize int) ([]Captcha, int64, error) {
	var captchas []Captcha
	var count int64

	if err := db.Model(c).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&captchas).Error; err != nil {
		return nil, 0, err
	}

	return captchas, count, nil
}
