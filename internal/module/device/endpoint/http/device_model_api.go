package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterModelRoutes(router fiber.Router, db *gorm.DB) {
	modelGroup := router.Group("/model")

	// @Summary 创建设备型号
	// @Description 创建新的设备型号记录
	// @Tags 设备型号
	// @Accept json
	// @Produce json
	// @Param deviceModel body model.DeviceModel true "设备型号"
	// @Success 200 {object} response.Response{data=model.DeviceModel}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /model [post]
	modelGroup.Post("", func(c *fiber.Ctx) error {
		var deviceModel model.DeviceModel
		if err := c.BodyParser(&deviceModel); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := deviceModel.Create(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, deviceModel)
	})

	// @Summary 获取设备型号
	// @Description 根据ID获取设备型号
	// @Tags 设备型号
	// @Produce json
	// @Param id path int true "设备型号ID"
	// @Success 200 {object} response.Response{data=model.DeviceModel}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /model/{id} [get]
	modelGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var deviceModel model.DeviceModel
		if err := deviceModel.GetByID(db, uint(id)); err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device model not found")
		}

		return response.Success(c, deviceModel)
	})

	// @Summary 更新设备型号
	// @Description 根据ID更新设备型号
	// @Tags 设备型号
	// @Accept json
	// @Produce json
	// @Param id path int true "设备型号ID"
	// @Param deviceModel body model.DeviceModel true "设备型号"
	// @Success 200 {object} response.Response{data=model.DeviceModel}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /model/{id} [put]
	modelGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var deviceModel model.DeviceModel
		if err := c.BodyParser(&deviceModel); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		deviceModel.ID = uint(id)
		if err := deviceModel.Update(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, deviceModel)
	})

	// @Summary 删除设备型号
	// @Description 根据ID删除设备型号
	// @Tags 设备型号
	// @Produce json
	// @Param id path int true "设备型号ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /model/{id} [delete]
	modelGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var deviceModel model.DeviceModel
		deviceModel.ID = uint(id)
		if err := deviceModel.Delete(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备型号列表
	// @Description 分页获取设备型号列表
	// @Tags 设备型号
	// @Produce json
	// @Param page query int false "页码" default(1)
	// @Param pageSize query int false "每页数量" default(10)
	// @Success 200 {object} response.Response{data=[]model.DeviceModel,total=int}
	// @Failure 500 {object} response.Response
	// @Router /model [get]
	modelGroup.Get("", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 10)

		var deviceModel model.DeviceModel
		models, count, err := deviceModel.List(db, page, pageSize)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.SuccessWithTotal(c, models, count)
	})
}
