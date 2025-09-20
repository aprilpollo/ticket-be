package port

import (
	"task-management/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type OrganizationRepository interface {
	GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error)
}

type OrganizationService interface {
	GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error)
}
