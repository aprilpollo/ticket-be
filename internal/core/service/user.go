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


func (s *UserService) GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error) {
	return s.uRepo.GetUserByID(ctx, id)
}


func (s *UserService) GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error) {
	return s.uRepo.GetAllUsers(ctx)
}
