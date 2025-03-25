package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterPositionRoutes(router fiber.Router, db *gorm.DB) {
	positionGroup := router.Group("/position")
	repo := model.NewDevicePositionRepo(db)

	// @Summary 创建设备位置信息
	// @Description 创建新的设备位置信息记录
	// @Tags 设备位置
	// @Accept json
	// @Produce json
	// @Param position body model.DevicePosition true "设备位置信息"
	// @Success 200 {object} response.Response{data=model.DevicePosition}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /position [post]
	positionGroup.Post("", func(c *fiber.Ctx) error {
		var position model.DevicePosition
		if err := c.BodyParser(&position); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := repo.Create(&position); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, position)
	})

	// @Summary 获取设备位置信息
	// @Description 根据ID获取设备位置信息
	// @Tags 设备位置
	// @Produce json
	// @Param id path int true "位置信息ID"
	// @Success 200 {object} response.Response{data=model.DevicePosition}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /position/{id} [get]
	positionGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		position, err := repo.GetByID(uint(id))
		if err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device position not found")
		}

		return response.Success(c, position)
	})

	// @Summary 更新设备位置信息
	// @Description 根据ID更新设备位置信息
	// @Tags 设备位置
	// @Accept json
	// @Produce json
	// @Param id path int true "位置信息ID"
	// @Param position body model.DevicePosition true "设备位置信息"
	// @Success 200 {object} response.Response{data=model.DevicePosition}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /position/{id} [put]
	positionGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var position model.DevicePosition
		if err := c.BodyParser(&position); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		position.ID = uint(id)
		if err := repo.Update(&position); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, position)
	})

	// @Summary 删除设备位置信息
	// @Description 根据ID删除设备位置信息
	// @Tags 设备位置
	// @Produce json
	// @Param id path int true "位置信息ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /position/{id} [delete]
	positionGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var position model.DevicePosition
		position.ID = uint(id)
		if err := repo.Delete(&position); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备位置信息列表
	// @Description 根据设备ID获取关联的位置信息列表
	// @Tags 设备位置
	// @Produce json
	// @Param device_id path int true "设备ID"
	// @Success 200 {object} response.Response{data=[]model.DevicePosition}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /position/device/{device_id} [get]
	positionGroup.Get("/device/:device_id", func(c *fiber.Ctx) error {
		deviceID, err := c.ParamsInt("device_id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid device ID")
		}

		positions, err := repo.ListByDeviceID(deviceID)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, positions)
	})
}
