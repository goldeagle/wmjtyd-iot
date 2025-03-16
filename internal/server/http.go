package server

import (
	"context"
	"wmjtyd-iot/internal/server/route"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	App    *fiber.App
	logger *zap.Logger
}

func NewHTTPServer(addr string, logger *zap.Logger, db *gorm.DB) *HTTPServer {
	app := fiber.New()

	server := &HTTPServer{
		App:    app,
		logger: logger,
	}
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
