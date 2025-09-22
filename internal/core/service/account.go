package service

import (
	"task-management/internal/core/domain"
	"task-management/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type AccountService struct {
	aRepo port.AccountRepository
}

func NewAccountService(aRepo port.AccountRepository) *AccountService {
	return &AccountService{aRepo: aRepo}
}

func (s *AccountService) GetAccount(ctx *fiber.Ctx, accID uint) (*domain.Account, error) {
	return s.aRepo.GetAccount(ctx, accID)
}

func (s *AccountService) GetOrganization(ctx *fiber.Ctx, accID uint) (int64, int64, int64, []*domain.Organization, error) {
	return s.aRepo.GetOrganization(ctx, accID)
}

func (s *AccountService) GetRoleInOrganization(ctx *fiber.Ctx, accID, orgID uint) (*domain.RoleInOrganization, error) {
	return s.aRepo.GetRoleInOrganization(ctx, accID, orgID)
}