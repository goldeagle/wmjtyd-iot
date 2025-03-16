package model

import (
	"gorm.io/gorm"
)

// Token 用户Token表
// @Description 用户Token表
type Token struct {
	gorm.Model
	ID         int    `gorm:"primaryKey;column:id" json:"id"`        // ID
	UserID     int    `gorm:"column:user_id" json:"user_id"`         // 用户ID
	Token      string `gorm:"column:token" json:"token"`             // Token
	ExpireTime int    `gorm:"column:expire_time" json:"expire_time"` // 过期时间
	CreateTime int    `gorm:"column:create_time" json:"create_time"` // 创建时间
	UpdateTime int    `gorm:"column:update_time" json:"update_time"` // 更新时间
	Status     int    `gorm:"column:status" json:"status"`           // 状态
}

// TableName 指定表名
func (Token) TableName() string {
	return "bsd_token"
}

// Create 创建Token
func (t *Token) Create(db *gorm.DB) error {
	return db.Create(t).Error
}

// Update 更新Token
func (t *Token) Update(db *gorm.DB) error {
	return db.Save(t).Error
}

// Delete 删除Token
func (t *Token) Delete(db *gorm.DB) error {
	return db.Delete(t).Error
}

// GetByID 根据ID获取Token
func (t *Token) GetByID(db *gorm.DB, id int) error {
	return db.First(t, id).Error
}

// List 获取Token列表
func (t *Token) List(db *gorm.DB, page, pageSize int) ([]Token, int64, error) {
	var tokens []Token
	var count int64

	if err := db.Model(t).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&tokens).Error; err != nil {
		return nil, 0, err
	}

	return tokens, count, nil
}
