package http

import (
	"wmjtyd-iot/internal/module/device/model"
	"wmjtyd-iot/internal/server/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterNetRoutes(router fiber.Router, db *gorm.DB) {
	netGroup := router.Group("/net")
	repo := model.NewDeviceNetRepo(db)

	// @Summary 创建设备网络信息
	// @Description 创建新的设备网络信息记录
	// @Tags 设备网络
	// @Accept json
	// @Produce json
	// @Param net body model.DeviceNet true "设备网络信息"
	// @Success 200 {object} response.Response{data=model.DeviceNet}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /net [post]
	netGroup.Post("", func(c *fiber.Ctx) error {
		var net model.DeviceNet
		if err := c.BodyParser(&net); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		if err := repo.Create(&net); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, net)
	})

	// @Summary 获取设备网络信息
	// @Description 根据ID获取设备网络信息
	// @Tags 设备网络
	// @Produce json
	// @Param id path int true "网络信息ID"
	// @Success 200 {object} response.Response{data=model.DeviceNet}
	// @Failure 400 {object} response.Response
	// @Failure 404 {object} response.Response
	// @Router /net/{id} [get]
	netGroup.Get("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		net, err := repo.GetByID(uint(id))
		if err != nil {
			return response.Fail(c, fiber.StatusNotFound, "Device net info not found")
		}

		return response.Success(c, net)
	})

	// @Summary 更新设备网络信息
	// @Description 根据ID更新设备网络信息
	// @Tags 设备网络
	// @Accept json
	// @Produce json
	// @Param id path int true "网络信息ID"
	// @Param net body model.DeviceNet true "设备网络信息"
	// @Success 200 {object} response.Response{data=model.DeviceNet}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /net/{id} [put]
	netGroup.Put("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var net model.DeviceNet
		if err := c.BodyParser(&net); err != nil {
			return response.Fail(c, fiber.StatusBadRequest, err.Error())
		}

		net.ID = uint(id)
		if err := repo.Update(&net); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, net)
	})

	// @Summary 删除设备网络信息
	// @Description 根据ID删除设备网络信息
	// @Tags 设备网络
	// @Produce json
	// @Param id path int true "网络信息ID"
	// @Success 200 {object} response.Response
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /net/{id} [delete]
	netGroup.Delete("/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid ID")
		}

		var net model.DeviceNet
		net.ID = uint(id)
		if err := repo.Delete(&net); err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nil)
	})

	// @Summary 获取设备网络信息列表
	// @Description 根据设备ID获取关联的网络信息列表
	// @Tags 设备网络
	// @Produce json
	// @Param device_id path int true "设备ID"
	// @Success 200 {object} response.Response{data=[]model.DeviceNet}
	// @Failure 400 {object} response.Response
	// @Failure 500 {object} response.Response
	// @Router /net/device/{device_id} [get]
	netGroup.Get("/device/:device_id", func(c *fiber.Ctx) error {
		deviceID, err := c.ParamsInt("device_id")
		if err != nil {
			return response.Fail(c, fiber.StatusBadRequest, "Invalid device ID")
		}

		nets, err := repo.ListByDeviceID(deviceID)
		if err != nil {
			return response.Fail(c, fiber.StatusInternalServerError, err.Error())
		}

		return response.Success(c, nets)
	})
}
