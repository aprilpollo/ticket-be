package service

import (
	"task-management/internal/core/domain"
	"task-management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	uRepo port.UserRepository
}

func NewUserService(uRepo port.UserRepository) *UserService {
	return &UserService{uRepo: uRepo}
}

func (s *UserService) GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error) {
	return s.uRepo.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error) {
	return s.uRepo.GetUserByID(ctx, id)
}

func (s *UserService) UpdateUser(ctx *fiber.Ctx, id uint, user *domain.UpdateUserRequest) (*domain.User, error) {
	return s.uRepo.UpdateUser(ctx, id, user)
}

func (s *UserService) DeleteUser(ctx *fiber.Ctx, id uint) error {
	return s.uRepo.DeleteUser(ctx, id)
}