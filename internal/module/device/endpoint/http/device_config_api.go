package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterConfigRoutes(router fiber.Router, db *gorm.DB) {
	configGroup := router.Group("/api/device/config")

	// @Summary 创建设备配置
	// @Description 创建新的设备配置记录
	// @Tags 设备配置
	// @Accept json
	// @Produce json
	// @Param config body model.DeviceConfig true "设备配置信息"
	// @Success 200 {object} response.Response{data=model.DeviceConfig}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/config [post]
	configGroup.Post("", func(c *fiber.Ctx) error {
		var config model.DeviceConfig
		if err := c.BodyParser(&config); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := config.Create(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, config)
	})

	// @Summary 获取设备配置
	// @Description 根据ID获取设备配置详情
	// @Tags 设备配置
	// @Produce json
	// @Param id path int true "配置ID"
	// @Success 200 {object} response.Response{data=model.DeviceConfig}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /api/device/config/{id} [get]
	configGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var config model.DeviceConfig
		if err := config.GetByID(db, uint(id)); err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device config not found")
		}

		return response.Success(c, config)
	})

	// @Summary 更新设备配置
	// @Description 根据ID更新设备配置
	// @Tags 设备配置
	// @Accept json
	// @Produce json
	// @Param id path int true "配置ID"
	// @Param config body model.DeviceConfig true "设备配置信息"
	// @Success 200 {object} response.Response{data=model.DeviceConfig}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/config/{id} [put]
	configGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var config model.DeviceConfig
		if err := c.BodyParser(&config); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		config.ID = uint(id)
		if err := config.Update(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, config)
	})

	// @Summary 删除设备配置
	// @Description 根据ID删除设备配置
	// @Tags 设备配置
	// @Produce json
	// @Param id path int true "配置ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/config/{id} [delete]
	configGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var config model.DeviceConfig
		config.ID = uint(id)
		if err := config.Delete(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备配置列表
	// @Description 分页获取设备配置列表
	// @Tags 设备配置
	// @Produce json
	// @Param page query int false "页码" default(1)
	// @Param pageSize query int false "每页数量" default(10)
	// @Success 200 {object} response.Response{data=[]model.DeviceConfig,total=int}
	// @Failure 500 {object} response.Response
	// @Router /api/device/config [get]
	configGroup.Get("", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 10)

		var config model.DeviceConfig
		configs, count, err := config.List(db, page, pageSize)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.SuccessWithTotal(c, configs, count)
	})

	// @Summary 根据设备ID获取配置
	// @Description 根据设备ID获取相关配置列表
	// @Tags 设备配置
	// @Produce json
	// @Param device_id path int true "设备ID"
	// @Success 200 {object} response.Response{data=[]model.DeviceConfig}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/config/device/{device_id} [get]
	configGroup.Get("/device/:device_id", func(c *fiber.Ctx) error {
		deviceID, err := c.ParamsInt("device_id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid device ID")
		}

		var config model.DeviceConfig
		configs, err := config.GetByDeviceID(db, uint(deviceID))
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, configs)
	})
}
