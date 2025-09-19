package main

import (
	"fmt"
	"log"
	config "task-management/internal/adapter/config"
	httpfiber "task-management/internal/adapter/handler/fiber"
	gormOrm "task-management/internal/adapter/storage/gorm"
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

	// Initialize services
	userService := service.NewUserService(userRepo)
	authService := service.NewAuthService(authRepo)

	// Initialize app with auth service
	app := httpfiber.NewApp()


	// Initialize handlers
	userHandler := httpfiber.NewUserHandler(userService)
	authHandler := httpfiber.NewAuthHandler(authService)

	// Initialize routes
	app.MainRoutes()
	app.UserRoutes(userHandler)
	app.AuthRoutes(authHandler)

	fmt.Println("[INFO] Starting server...")
	app.Serve(fmt.Sprintf(":%s", config.Env.ApiPort))
}
