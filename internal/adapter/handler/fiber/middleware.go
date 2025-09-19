package httpfiber

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/golang-jwt/jwt/v5"
	"task-management/internal/core/domain"
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
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Authorization header is required", nil)
		}

		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Invalid authorization header format", nil)
		}

		claims := &domain.JWTClaims{}
		secretKey := []byte("your-secret-key") 

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Invalid token", nil)
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Invalid token", nil)
		}

		c.Locals("user_id", claims.UserID)
		c.Locals("email", claims.Email)

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
