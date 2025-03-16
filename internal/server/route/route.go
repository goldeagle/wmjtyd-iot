package route

import "github.com/gofiber/fiber/v2"

// Route 路由结构体
type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}
