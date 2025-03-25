package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterLogRoutes(router fiber.Router, db *gorm.DB) {
	logGroup := router.Group("/api/device/log")

	// @Summary 创建设备日志
	// @Description 创建新的设备日志记录
	// @Tags 设备日志
	// @Accept json
	// @Produce json
	// @Param log body model.DeviceLog true "设备日志信息"
	// @Success 200 {object} response.Response{data=model.DeviceLog}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/log [post]
	logGroup.Post("", func(c *fiber.Ctx) error {
		var log model.DeviceLog
		if err := c.BodyParser(&log); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := log.Create(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, log)
	})

	// @Summary 获取设备日志
	// @Description 根据ID获取设备日志详情
	// @Tags 设备日志
	// @Produce json
	// @Param id path int true "日志ID"
	// @Success 200 {object} response.Response{data=model.DeviceLog}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /api/device/log/{id} [get]
	logGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var log model.DeviceLog
		if err := log.GetByID(db, uint(id)); err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device log not found")
		}

		return response.Success(c, log)
	})

	// @Summary 更新设备日志
	// @Description 根据ID更新设备日志
	// @Tags 设备日志
	// @Accept json
	// @Produce json
	// @Param id path int true "日志ID"
	// @Param log body model.DeviceLog true "设备日志信息"
	// @Success 200 {object} response.Response{data=model.DeviceLog}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/log/{id} [put]
	logGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var log model.DeviceLog
		if err := c.BodyParser(&log); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		log.ID = uint(id)
		if err := log.Update(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, log)
	})

	// @Summary 删除设备日志
	// @Description 根据ID删除设备日志
	// @Tags 设备日志
	// @Produce json
	// @Param id path int true "日志ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/log/{id} [delete]
	logGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var log model.DeviceLog
		log.ID = uint(id)
		if err := log.Delete(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备日志列表
	// @Description 分页获取设备日志列表
	// @Tags 设备日志
	// @Produce json
	// @Param page query int false "页码" default(1)
	// @Param pageSize query int false "每页数量" default(10)
	// @Success 200 {object} response.Response{data=[]model.DeviceLog,total=int}
	// @Failure 500 {object} response.Response
	// @Router /api/device/log [get]
	logGroup.Get("", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 10)

		var log model.DeviceLog
		logs, count, err := log.List(db, page, pageSize)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.SuccessWithTotal(c, logs, count)
	})
}
