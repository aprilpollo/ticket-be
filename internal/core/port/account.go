package port

import (
	"task-management/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type AccountRepository interface {
	GetAccount(ctx *fiber.Ctx, accID uint) (*domain.Account, error)
	GetOrganization(ctx *fiber.Ctx, accID uint) (int64, int64, int64, []*domain.Organization, error)
	GetRoleInOrganization(ctx *fiber.Ctx, accID, orgID uint) (*domain.RoleInOrganization, error)
	UpdateAccount(ctx *fiber.Ctx, userID uint, acc *domain.Account) (*domain.Account, error)
}

type AccountService interface {
	GetAccount(ctx *fiber.Ctx, accID uint) (*domain.Account, error)
	GetOrganization(ctx *fiber.Ctx, accID uint) (int64, int64, int64, []*domain.Organization, error)
	GetRoleInOrganization(ctx *fiber.Ctx, accID, orgID uint) (*domain.RoleInOrganization, error)
	UpdateAccount(ctx *fiber.Ctx, accID uint, acc *domain.Account) (*domain.Account, error)
}