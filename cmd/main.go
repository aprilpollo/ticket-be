package main

import (
	"fmt"
	"log"
	config "task-management/internal/adapter/config"
	"task-management/internal/adapter/handler/fiber"
	"task-management/internal/adapter/handler/fiber/middleware"
	"task-management/internal/adapter/handler/fiber/routes"
	"task-management/internal/adapter/storage/gorm"
	"task-management/internal/adapter/storage/gorm/repository"
	"task-management/internal/core/service"
)

func main() {
	config.LoadConfig()

	fmt.Println("[INFO] Initializing database connection...")
	err := gormOrm.Init(
		config.Env.Postgre.URI,
		config.Env.Postgre.MaxIdleConns,
		config.Env.Postgre.MaxOpenConns,
		config.Env.Postgre.ConnMaxLifetime,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("[INFO] Database connection initialized successfully")

	// Initialize repositories
	userRepo := repository.NewUserRepository(gormOrm.Trx)
	authRepo := repository.NewAuthRepository(gormOrm.Trx)
	organizationRepo := repository.NewOrganizationRepository(gormOrm.Trx)

	// Initialize services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo)
	organizationService := service.NewOrganizationService(organizationRepo)

	// Initialize handlers
	userHandler := routes.NewUserHandler(userService)
	authHandler := routes.NewAuthHandler(authService)
	organizationHandler := routes.NewOrganizationHandler(organizationService)

	// Initialize middleware
	mOrganization := middleware.NewOrganizationMiddleware(organizationService)

	// Initialize App routes
	app := httpfiber.NewApp()
	app.MainRoutes()
	app.AuthRoutes(authHandler)
	app.UserRoutes(userHandler, mOrganization)
	app.OrganizationRoutes(organizationHandler, mOrganization)

	fmt.Println("[INFO] Starting server...")
	app.Serve(fmt.Sprintf(":%s", config.Env.ApiPort))
}
