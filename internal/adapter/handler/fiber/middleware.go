package httpfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"task-management/internal/util"
)

type MiddlewareHandler struct {
	app *fiber.App
}

func NewMiddlewareHandler(app *fiber.App) *MiddlewareHandler {
	return &MiddlewareHandler{
		app: app,
	}
}

// SetupGlobalMiddleware sets up all global middleware
func (m *MiddlewareHandler) SetupGlobalMiddleware() {
	// Recovery middleware - recovers from panics
	m.app.Use(recover.New())

	// Request ID middleware - adds unique ID to each request
	m.app.Use(requestid.New())

	// Logger middleware - logs HTTP requests
	m.app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${method} ${path} (${ip}) ${latency}\n",
	}))
	m.app.Use(util.ZapLogger())

	// CORS middleware - handles Cross-Origin Resource Sharing
	m.app.Use(cors.New(cors.Config{
		AllowOrigins:     "*", // In production, specify exact origins
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false, // Set to false when using wildcard origins
	}))
}

// AuthMiddleware - authentication middleware
func (m *MiddlewareHandler) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Implement JWT token validation
		// For now, just pass through
		return c.Next()
	}
}

// RateLimitMiddleware - rate limiting middleware
func (m *MiddlewareHandler) RateLimitMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO: Implement rate limiting logic
		// For now, just pass through
		return c.Next()
	}
}
