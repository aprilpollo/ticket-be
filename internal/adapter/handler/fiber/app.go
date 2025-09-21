package httpfiber

import (
	"task-management/internal/adapter/handler/fiber/middleware"
	"task-management/internal/adapter/handler/fiber/routes"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	app  *fiber.App
	mApp *middleware.App
}

func NewApp() *App {
	app := fiber.New()

	mApp := middleware.NewMiddlewareHandler(app)
	mApp.SetupGlobalMiddleware()

	return &App{
		app:  app,
		mApp: mApp,
	}
}

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

func (r *App) AuthRoutes(authHandler *routes.AuthHandler) {
	api := r.app.Group("/api/v1")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/signin", authHandler.SignIn)
	auth.Post("/signup", authHandler.SignUp)
	auth.Get("/validate", r.mApp.AuthMiddleware(), authHandler.ValidateToken)
	auth.Get("/validate/user", r.mApp.AuthMiddleware(), authHandler.ValidateUser)
}

func (r *App) UserRoutes(userHandler *routes.UserHandler, mOrganization *middleware.OrganizationMiddleware) {
	api := r.app.Group("/api/v1")

	// User routes
	users := api.Group("/users")
	users.Use(r.mApp.AuthMiddleware())
	users.Use(mOrganization.Middleware())

	users.Get("/", mOrganization.MiddlewareWithPermission("IsPreview"), userHandler.GetAllUsers)
	users.Get("/:id", mOrganization.MiddlewareWithPermission("IsPreview"), userHandler.GetUserByID)
	users.Put("/:id", mOrganization.MiddlewareWithPermission("CanManageMembers"), userHandler.UpdateUser)
	users.Delete("/:id", mOrganization.MiddlewareWithPermission("CanManageMembers"), userHandler.DeleteUser)
}

func (r *App) OrganizationRoutes(organizationHandler *routes.OrganizationHandler, mOrganization *middleware.OrganizationMiddleware) {
	api := r.app.Group("/api/v1")
	organizations := api.Group("/organizations")

	{
		organizations.Use(r.mApp.AuthMiddleware())
		organizations.Use(mOrganization.Middleware())
	}

	{
		organizations.Get("/", mOrganization.MiddlewareWithPermission("CanViewReports"), organizationHandler.GetOrganization)
		organizations.Get("/roles", organizationHandler.GetUserRoleInOrganization)
		//organizations.Put("/:id", organizationHandler.UpdateOrganization)
	}
}

func (r *App) Serve(port string) error {
	return r.app.Listen(port)
}
