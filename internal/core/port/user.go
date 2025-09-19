package port

import (
	"task-management/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error)
	GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error)
}

type UserService interface {
	GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error)
	GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error)
}
