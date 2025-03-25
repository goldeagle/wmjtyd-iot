package http

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// RegisterDeviceRoutes 注册所有设备相关路由
func RegisterDeviceRoutes(router fiber.Router, db *gorm.DB) {
	deviceGroup := router.Group("/device")

	// 注册各模块路由
	RegisterCmdRoutes(deviceGroup, db)
	RegisterConfigRoutes(deviceGroup, db)
	RegisterFirmwareRoutes(deviceGroup, db)
	RegisterInfoRoutes(deviceGroup, db)
	RegisterLogRoutes(deviceGroup, db)
	RegisterModelRoutes(deviceGroup, db)
	RegisterNetRoutes(deviceGroup, db)
	RegisterPositionRoutes(deviceGroup, db)
	RegisterStatusRoutes(deviceGroup, db)
}
