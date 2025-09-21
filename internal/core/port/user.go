package port

import (
	"task-management/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error)
	GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error)
	UpdateUser(ctx *fiber.Ctx, id uint, user *domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(ctx *fiber.Ctx, id uint) error
}

type UserService interface {
	GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error)
	GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error)
	UpdateUser(ctx *fiber.Ctx, id uint, user *domain.UpdateUserRequest) (*domain.User, error)
	DeleteUser(ctx *fiber.Ctx, id uint) error
}
