package model

import (
	"gorm.io/gorm"
)

// User 用户表
// @Description 用户表
type User struct {
	gorm.Model
	ID            int    `gorm:"primaryKey;column:id" json:"id"`            // ID
	GroupID       int    `gorm:"column:group_id" json:"group_id"`           // 分组ID
	Username      string `gorm:"column:username" json:"username"`           // 用户名
	Nickname      string `gorm:"column:nickname" json:"nickname"`           // 昵称
	Email         string `gorm:"column:email" json:"email"`                 // 邮箱地址
	Mobile        string `gorm:"column:mobile" json:"mobile"`               // 手机号
	Avatar        string `gorm:"column:avatar" json:"avatar"`               // 头像
	Gender        int    `gorm:"column:gender" json:"gender"`               // 性别:0=未知,1=男,2=女
	Birthday      string `gorm:"column:birthday" json:"birthday"`           // 生日
	Money         int    `gorm:"column:money" json:"money"`                 // 余额
	Score         int    `gorm:"column:score" json:"score"`                 // 积分
	LastLoginTime int    `gorm:"column:lastlogintime" json:"lastlogintime"` // 上次登录时间
	LastLoginIP   string `gorm:"column:lastloginip" json:"lastloginip"`     // 登录IP
	LoginFailure  int    `gorm:"column:loginfailure" json:"loginfailure"`   // 失败次数
	JoinIP        string `gorm:"column:joinip" json:"joinip"`               // 加入IP
	JoinTime      int    `gorm:"column:jointime" json:"jointime"`           // 加入时间
	Motto         string `gorm:"column:motto" json:"motto"`                 // 签名
	Password      string `gorm:"column:password" json:"password"`           // 密码
	Salt          string `gorm:"column:salt" json:"salt"`                   // 密码盐
	Status        string `gorm:"column:status" json:"status"`               // 状态
	UpdateTime    int    `gorm:"column:updatetime" json:"updatetime"`       // 更新时间
	CreateTime    int    `gorm:"column:createtime" json:"createtime"`       // 创建时间
}

// TableName 指定表名
func (User) TableName() string {
	return "bsd_user"
}

// Create 创建用户
func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

// Update 更新用户
func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

// Delete 删除用户
func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

// GetByID 根据ID获取用户
func (u *User) GetByID(db *gorm.DB, id int) error {
	return db.First(u, id).Error
}

// List 获取用户列表
func (u *User) List(db *gorm.DB, page, pageSize int) ([]User, int64, error) {
	var users []User
	var count int64

	if err := db.Model(u).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, count, nil
}
