package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterFirmwareRoutes(router fiber.Router, db *gorm.DB) {
	firmwareGroup := router.Group("/api/device/firmware")

	// @Summary 创建设备固件
	// @Description 创建新的设备固件记录
	// @Tags 设备固件
	// @Accept json
	// @Produce json
	// @Param firmware body model.DeviceFirmware true "设备固件信息"
	// @Success 200 {object} response.Response{data=model.DeviceFirmware}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/firmware [post]
	firmwareGroup.Post("", func(c *fiber.Ctx) error {
		var firmware model.DeviceFirmware
		if err := c.BodyParser(&firmware); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := firmware.Create(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, firmware)
	})

	// @Summary 获取设备固件
	// @Description 根据ID获取设备固件详情
	// @Tags 设备固件
	// @Produce json
	// @Param id path int true "固件ID"
	// @Success 200 {object} response.Response{data=model.DeviceFirmware}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /api/device/firmware/{id} [get]
	firmwareGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var firmware model.DeviceFirmware
		if err := firmware.GetByID(db, uint(id)); err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device firmware not found")
		}

		return response.Success(c, firmware)
	})

	// @Summary 更新设备固件
	// @Description 根据ID更新设备固件
	// @Tags 设备固件
	// @Accept json
	// @Produce json
	// @Param id path int true "固件ID"
	// @Param firmware body model.DeviceFirmware true "设备固件信息"
	// @Success 200 {object} response.Response{data=model.DeviceFirmware}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/firmware/{id} [put]
	firmwareGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var firmware model.DeviceFirmware
		if err := c.BodyParser(&firmware); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		firmware.ID = uint(id)
		if err := firmware.Update(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, firmware)
	})

	// @Summary 删除设备固件
	// @Description 根据ID删除设备固件
	// @Tags 设备固件
	// @Produce json
	// @Param id path int true "固件ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/firmware/{id} [delete]
	firmwareGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var firmware model.DeviceFirmware
		firmware.ID = uint(id)
		if err := firmware.Delete(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备固件列表
	// @Description 分页获取设备固件列表
	// @Tags 设备固件
	// @Produce json
	// @Param page query int false "页码" default(1)
	// @Param pageSize query int false "每页数量" default(10)
	// @Success 200 {object} response.Response{data=[]model.DeviceFirmware,total=int}
	// @Failure 500 {object} response.Response
	// @Router /api/device/firmware [get]
	firmwareGroup.Get("", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 10)

		var firmware model.DeviceFirmware
		firmwares, count, err := firmware.List(db, page, pageSize)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.SuccessWithTotal(c, firmwares, count)
	})
}
