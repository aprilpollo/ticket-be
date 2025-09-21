package service

import (
	"task-management/internal/core/domain"
	"task-management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type OrganizationService struct {
	oRepo port.OrganizationRepository
}

func NewOrganizationService(oRepo port.OrganizationRepository) *OrganizationService {
	return &OrganizationService{oRepo: oRepo}
}


func (s *OrganizationService) GetOrganization(ctx *fiber.Ctx) (int64, int64, int64, []*domain.Organization, error) {
	return s.oRepo.GetOrganization(ctx)
}

func (s *OrganizationService) GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error) {
	return s.oRepo.GetUserRoleInOrganization(ctx, userID)
}

func (s *OrganizationService) GetUserRoleInOrganizationByID(ctx *fiber.Ctx, orgId uint, userId uint) (*domain.OrganizationMemberRole, error) {
	return s.oRepo.GetUserRoleInOrganizationByID(ctx, orgId, userId)
}

func (s *OrganizationService) UpdateOrganization(ctx *fiber.Ctx, id uint, organization *domain.UpdateOrganizationRequest) (*domain.Organization, error) {
	return s.oRepo.UpdateOrganization(ctx, id, organization)
}