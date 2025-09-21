package port

import (
	"task-management/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type OrganizationRepository interface {
	GetOrganization(ctx *fiber.Ctx) (int64, int64, int64, []*domain.Organization, error)
	GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error)
	GetUserRoleInOrganizationByID(ctx *fiber.Ctx, orgId uint, userId uint) (*domain.OrganizationMemberRole, error)
	UpdateOrganization(ctx *fiber.Ctx, id uint, organization *domain.UpdateOrganizationRequest) (*domain.Organization, error)
}

type OrganizationService interface {
	GetOrganization(ctx *fiber.Ctx) (int64, int64, int64, []*domain.Organization, error)
	GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error)
	GetUserRoleInOrganizationByID(ctx *fiber.Ctx, orgId uint, userId uint) (*domain.OrganizationMemberRole, error)
	UpdateOrganization(ctx *fiber.Ctx, id uint, organization *domain.UpdateOrganizationRequest) (*domain.Organization, error)
}
