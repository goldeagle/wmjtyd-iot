package model

import (
	"gorm.io/gorm"
)

// Menu 菜单表
// @Description 菜单表
type Menu struct {
	gorm.Model
	MenuID         int    `gorm:"primaryKey;column:menu_id" json:"menu_id"`        // 菜单ID
	PID            int    `gorm:"column:pid" json:"pid"`                           // 父级ID
	ControllerName string `gorm:"column:controller_name" json:"controller_name"`   // 模块名称
	Title          string `gorm:"column:title" json:"title"`                       // 模块标题
	PKID           string `gorm:"column:pk_id" json:"pk_id"`                       // 主键名
	TableNameStr   string `gorm:"column:table_name" json:"table_name"`             // 模块数据库表
	IsCreate       int    `gorm:"column:is_create" json:"is_create"`               // 是否允许生成模块
	Status         int    `gorm:"column:status" json:"status"`                     // 状态
	SortID         int    `gorm:"column:sortid" json:"sortid"`                     // 排序号
	TableStatus    int    `gorm:"column:table_status" json:"table_status"`         // 是否生成数据库表
	IsURL          int    `gorm:"column:is_url" json:"is_url"`                     // 是否只是URL链接
	URL            string `gorm:"column:url" json:"url"`                           // URL
	MenuIcon       string `gorm:"column:menu_icon" json:"menu_icon"`               // icon字体图标
	TabMenu        string `gorm:"column:tab_menu" json:"tab_menu"`                 // tab选项卡菜单配置
	AppID          int    `gorm:"column:app_id" json:"app_id"`                     // 所属模块
	IsSubmit       int    `gorm:"column:is_submit" json:"is_submit"`               // 是否允许投稿
	UploadConfigID int    `gorm:"column:upload_config_id" json:"upload_config_id"` // 上传配置ID
	Connect        string `gorm:"column:connect" json:"connect"`                   // 数据库连接
}

// TableName 指定表名
func (Menu) TableName() string {
	return "cd_menu"
}

// Create 创建菜单
func (m *Menu) Create(db *gorm.DB) error {
	return db.Create(m).Error
}

// Update 更新菜单
func (m *Menu) Update(db *gorm.DB) error {
	return db.Save(m).Error
}

// Delete 删除菜单
func (m *Menu) Delete(db *gorm.DB) error {
	return db.Delete(m).Error
}

// GetByID 根据ID获取菜单
func (m *Menu) GetByID(db *gorm.DB, id int) error {
	return db.First(m, id).Error
}

// List 获取菜单列表
func (m *Menu) List(db *gorm.DB, page, pageSize int) ([]Menu, int64, error) {
	var menus []Menu
	var count int64

	if err := db.Model(m).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&menus).Error; err != nil {
		return nil, 0, err
	}

	return menus, count, nil
}
