package server

import (
	"context"
	deviceHttp "wmjtyd-iot/internal/module/device/endpoint/http"
	"wmjtyd-iot/internal/server/route"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
)

type HTTPServer struct {
	App    *fiber.App
	logger *zap.Logger
}

func NewHTTPServer(addr string, logger *zap.Logger, db *gorm.DB) *HTTPServer {
	app := fiber.New()

	// 根据配置决定是否启用Swagger UI
	if viper.GetBool("service.enable_swagger") {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	server := &HTTPServer{
		App:    app,
		logger: logger,
	}

	// 注册设备路由（通过routes.go统一注册）
	deviceHttp.RegisterDeviceRoutes(app.Group("/api"), db)

	return server
}

func (s *HTTPServer) RegisterRoutes(routes []route.Route) {
	for _, route := range routes {
		switch route.Method {
		case "GET":
			s.App.Get(route.Path, route.Handler)
		case "POST":
			s.App.Post(route.Path, route.Handler)
		case "PUT":
			s.App.Put(route.Path, route.Handler)
		case "DELETE":
			s.App.Delete(route.Path, route.Handler)
		}
		s.logger.Info("Registered HTTP route",
			zap.String("method", route.Method),
			zap.String("path", route.Path))
	}
}

func (s *HTTPServer) Start() error {
	s.logger.Info("Starting HTTP server")
	return s.App.Listen(":3000")
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server")
	return s.App.Shutdown()
}
