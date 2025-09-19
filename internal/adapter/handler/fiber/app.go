package httpfiber

import (
	"github.com/gofiber/fiber/v2"
)

type App struct {
	app        *fiber.App
	middleware *MiddlewareHandler
}

func NewApp() *App {
	app := fiber.New()
	middleware := NewMiddlewareHandler(app)
	middleware.SetupGlobalMiddleware()

	return &App{
		app:        app,
		middleware: middleware,
	}
}

// GetApp returns the fiber app instance
func (r *App) GetApp() *fiber.App {
	return r.app
}

func (r *App) MainRoutes() {
	r.app.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"version": "1.0.0",
		})
	})
	r.app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})
}

func (r *App) UserRoutes(userHandler *UserHandler) {
	api := r.app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Get("/", userHandler.GetAllUsers)
	users.Get("/:id", userHandler.GetUserByID)

}

func (r *App) AuthRoutes(authHandler *AuthHandler) {
	api := r.app.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/signin", authHandler.SignIn)
	auth.Post("/signup", authHandler.SignUp)
	auth.Get("/validate", r.middleware.AuthMiddleware(), authHandler.ValidateToken)
	auth.Get("/validate/user", r.middleware.AuthMiddleware(), authHandler.ValidateUser)
}

func (r *App) Serve(port string) error {
	return r.app.Listen(port)
}
