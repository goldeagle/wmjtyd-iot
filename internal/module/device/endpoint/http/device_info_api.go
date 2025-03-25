package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterInfoRoutes(router fiber.Router, db *gorm.DB) {
	deviceGroup := router.Group("/info")

	// @Summary 创建设备信息
	// @Description 创建新的设备信息记录
	// @Tags 设备信息
	// @Accept json
	// @Produce json
	// @Param deviceInfo body model.DeviceInfo true "设备信息"
	// @Success 200 {object} response.Response{data=model.DeviceInfo}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /info [post]
	deviceGroup.Post("", func(c *fiber.Ctx) error {
		var deviceInfo model.DeviceInfo
		if err := c.BodyParser(&deviceInfo); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := deviceInfo.Create(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, deviceInfo)
	})

	// @Summary 获取设备信息
	// @Description 根据ID获取设备信息
	// @Tags 设备信息
	// @Produce json
	// @Param id path int true "设备ID"
	// @Success 200 {object} response.Response{data=model.DeviceInfo}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /info/{id} [get]
	deviceGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var deviceInfo model.DeviceInfo
		if err := deviceInfo.GetByID(db, uint(id)); err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device not found")
		}

		return response.Success(c, deviceInfo)
	})

	// @Summary 更新设备信息
	// @Description 根据ID更新设备信息
	// @Tags 设备信息
	// @Accept json
	// @Produce json
	// @Param id path int true "设备ID"
	// @Param deviceInfo body model.DeviceInfo true "设备信息"
	// @Success 200 {object} response.Response{data=model.DeviceInfo}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /info/{id} [put]
	deviceGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var deviceInfo model.DeviceInfo
		if err := c.BodyParser(&deviceInfo); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		deviceInfo.ID = uint(id)
		if err := deviceInfo.Update(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, deviceInfo)
	})

	// @Summary 删除设备信息
	// @Description 根据ID删除设备信息
	// @Tags 设备信息
	// @Produce json
	// @Param id path int true "设备ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /info/{id} [delete]
	deviceGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var deviceInfo model.DeviceInfo
		deviceInfo.ID = uint(id)
		if err := deviceInfo.Delete(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备信息列表
	// @Description 分页获取设备信息列表
	// @Tags 设备信息
	// @Produce json
	// @Param page query int false "页码" default(1)
	// @Param pageSize query int false "每页数量" default(10)
	// @Success 200 {object} response.Response{data=[]model.DeviceInfo,total=int}
	// @Failure 500 {object} response.Response
	// @Router /info [get]
	deviceGroup.Get("", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 10)

		var deviceInfo model.DeviceInfo
		deviceInfos, count, err := deviceInfo.List(db, page, pageSize)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.SuccessWithTotal(c, deviceInfos, count)
	})
}
