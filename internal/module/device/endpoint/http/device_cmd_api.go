package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCmdRoutes(router fiber.Router, db *gorm.DB) {
	cmdGroup := router.Group("/api/device/cmd")

	// @Summary 创建设备指令
	// @Description 创建新的设备指令记录
	// @Tags 设备指令
	// @Accept json
	// @Produce json
	// @Param cmd body model.DeviceCmd true "设备指令信息"
	// @Success 200 {object} response.Response{data=model.DeviceCmd}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/cmd [post]
	cmdGroup.Post("", func(c *fiber.Ctx) error {
		var cmd model.DeviceCmd
		if err := c.BodyParser(&cmd); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := cmd.Create(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, cmd)
	})

	// @Summary 获取设备指令
	// @Description 根据ID获取设备指令详情
	// @Tags 设备指令
	// @Produce json
	// @Param id path int true "指令ID"
	// @Success 200 {object} response.Response{data=model.DeviceCmd}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /api/device/cmd/{id} [get]
	cmdGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var cmd model.DeviceCmd
		if err := cmd.GetByID(db, uint(id)); err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device command not found")
		}

		return response.Success(c, cmd)
	})

	// @Summary 更新设备指令
	// @Description 根据ID更新设备指令
	// @Tags 设备指令
	// @Accept json
	// @Produce json
	// @Param id path int true "指令ID"
	// @Param cmd body model.DeviceCmd true "设备指令信息"
	// @Success 200 {object} response.Response{data=model.DeviceCmd}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/cmd/{id} [put]
	cmdGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var cmd model.DeviceCmd
		if err := c.BodyParser(&cmd); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		cmd.ID = uint(id)
		if err := cmd.Update(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, cmd)
	})

	// @Summary 删除设备指令
	// @Description 根据ID删除设备指令
	// @Tags 设备指令
	// @Produce json
	// @Param id path int true "指令ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /api/device/cmd/{id} [delete]
	cmdGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var cmd model.DeviceCmd
		cmd.ID = uint(id)
		if err := cmd.Delete(db); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备指令列表
	// @Description 分页获取设备指令列表
	// @Tags 设备指令
	// @Produce json
	// @Param page query int false "页码" default(1)
	// @Param pageSize query int false "每页数量" default(10)
	// @Success 200 {object} response.Response{data=[]model.DeviceCmd,total=int}
	// @Failure 500 {object} response.Response
	// @Router /api/device/cmd [get]
	cmdGroup.Get("", func(c *fiber.Ctx) error {
		page := c.QueryInt("page", 1)
		pageSize := c.QueryInt("pageSize", 10)

		var cmd model.DeviceCmd
		cmds, count, err := cmd.List(db, page, pageSize)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.SuccessWithTotal(c, cmds, count)
	})
}
