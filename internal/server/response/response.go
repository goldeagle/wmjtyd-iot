package response

import (
	"github.com/gofiber/fiber/v2"
)

// Response 标准响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success 成功响应
func Success(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Code:    fiber.StatusOK,
		Message: "success",
		Data:    data,
	})
}

// SuccessWithTotal 带总数的成功响应
func SuccessWithTotal(c *fiber.Ctx, data interface{}, total int64) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "success",
		"data":    data,
		"total":   total,
	})
}

// Fail 失败响应
func Fail(c *fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
