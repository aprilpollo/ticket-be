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


func (s *OrganizationService) GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error) {
	return s.oRepo.GetUserRoleInOrganization(ctx, userID)
}