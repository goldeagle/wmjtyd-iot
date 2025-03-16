package model

import (
	"gorm.io/gorm"
)

// DeviceCmd 设备指令模型
// @Description 设备指令模型
type DeviceCmd struct {
	gorm.Model
	DeviceID uint   `gorm:"comment:设备ID" json:"device_id"`
	CmdType  string `gorm:"type:varchar(50);comment:指令类型" json:"cmd_type"`
	CmdData  string `gorm:"type:text;comment:指令数据" json:"cmd_data"`
	Status   string `gorm:"type:varchar(20);comment:指令状态" json:"status"`
	ExecTime int64  `gorm:"comment:执行时间" json:"exec_time"`
	Result   string `gorm:"type:text;comment:执行结果" json:"result"`
	ErrorMsg string `gorm:"type:text;comment:错误信息" json:"error_msg"`
}

// Create 创建设备指令记录
func (c *DeviceCmd) Create(db *gorm.DB) error {
	return db.Create(c).Error
}

// GetByID 根据ID获取设备指令记录
func (c *DeviceCmd) GetByID(db *gorm.DB, id uint) error {
	return db.First(c, id).Error
}

// Update 更新设备指令记录
func (c *DeviceCmd) Update(db *gorm.DB) error {
	return db.Save(c).Error
}

// Delete 删除设备指令记录
func (c *DeviceCmd) Delete(db *gorm.DB) error {
	return db.Delete(c).Error
}

// List 获取设备指令列表
func (c *DeviceCmd) List(db *gorm.DB, page, pageSize int) ([]DeviceCmd, int64, error) {
	var cmds []DeviceCmd
	var count int64

	err := db.Model(c).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&cmds).Error
	if err != nil {
		return nil, 0, err
	}

	return cmds, count, nil
}
