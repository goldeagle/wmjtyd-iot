package response

import "github.com/gofiber/fiber/v2"

func Error(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(fiber.Map{"code": code, "msg": msg, "data": nil})
}

func ErrorMap(c *fiber.Ctx, code int, msg map[string]string) error {
	return c.Status(code).JSON(fiber.Map{"code": code, "msg": msg, "data": nil})
}

func Success(c *fiber.Ctx, msg string, data interface{}) error {
	return c.JSON(fiber.Map{"code": 1, "msg": msg, "data": data})
}
