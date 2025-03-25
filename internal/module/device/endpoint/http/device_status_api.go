package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterStatusRoutes(router fiber.Router, db *gorm.DB) {
	statusGroup := router.Group("/status")
	repo := model.NewDeviceStatusRepo(db)

	// @Summary 创建设备状态信息
	// @Description 创建新的设备状态信息记录
	// @Tags 设备状态
	// @Accept json
	// @Produce json
	// @Param status body model.DeviceStatus true "设备状态信息"
	// @Success 200 {object} response.Response{data=model.DeviceStatus}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /status [post]
	statusGroup.Post("", func(c *fiber.Ctx) error {
		var status model.DeviceStatus
		if err := c.BodyParser(&status); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := repo.Create(&status); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, status)
	})

	// @Summary 获取设备状态信息
	// @Description 根据ID获取设备状态信息
	// @Tags 设备状态
	// @Produce json
	// @Param id path int true "状态信息ID"
	// @Success 200 {object} response.Response{data=model.DeviceStatus}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /status/{id} [get]
	statusGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		status, err := repo.GetByID(uint(id))
		if err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device status not found")
		}

		return response.Success(c, status)
	})

	// @Summary 更新设备状态信息
	// @Description 根据ID更新设备状态信息
	// @Tags 设备状态
	// @Accept json
	// @Produce json
	// @Param id path int true "状态信息ID"
	// @Param status body model.DeviceStatus true "设备状态信息"
	// @Success 200 {object} response.Response{data=model.DeviceStatus}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /status/{id} [put]
	statusGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var status model.DeviceStatus
		if err := c.BodyParser(&status); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		status.ID = uint(id)
		if err := repo.Update(&status); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, status)
	})

	// @Summary 删除设备状态信息
	// @Description 根据ID删除设备状态信息
	// @Tags 设备状态
	// @Produce json
	// @Param id path int true "状态信息ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /status/{id} [delete]
	statusGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var status model.DeviceStatus
		status.ID = uint(id)
		if err := repo.Delete(&status); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备状态信息列表
	// @Description 根据设备ID获取关联的状态信息列表
	// @Tags 设备状态
	// @Produce json
	// @Param device_id path int true "设备ID"
	// @Success 200 {object} response.Response{data=[]model.DeviceStatus}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /status/device/{device_id} [get]
	statusGroup.Get("/device/:device_id", func(c *fiber.Ctx) error {
		deviceID, err := c.ParamsInt("device_id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid device ID")
		}

		statuses, err := repo.ListByDeviceID(deviceID)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, statuses)
	})
}
