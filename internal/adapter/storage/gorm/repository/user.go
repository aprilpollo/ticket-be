package repository

import (
	"task-management/internal/adapter/storage/gorm/models"
	"task-management/internal/core/domain"
	"task-management/internal/util"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx *fiber.Ctx, user *domain.User) error {
	userModel := models.User{
		Email:              user.Email,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		DisplayName:        user.DisplayName,
		Bio:                user.Bio,
		Avatar:             user.Avatar,
		DateOfBirth:        user.DateOfBirth,
		Gender:             user.Gender,
		PhoneNumber:        user.PhoneNumber,
		LanguagePreference: user.LanguagePreference,
		TimeZone:           user.TimeZone,
		IsEmailVerified:    user.IsEmailVerified,
		IsPhoneVerified:    user.IsPhoneVerified,
		LastLoginAt:        user.LastLoginAt,
	}

	if err := r.db.Create(&userModel).Error; err != nil {
		return err
	}

	user.ID = userModel.ID
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt

	return nil
}

func (r *UserRepository) GetUserByID(ctx *fiber.Ctx, id uint) (*domain.User, error) {
	user, err := util.FindOne[models.User](ctx, r.db, int64(id))
	if err != nil {
		return nil, err
	}
	return r.modelToDomain(user), nil
}

func (r *UserRepository) GetAllUsers(ctx *fiber.Ctx) (int64, int64, int64, []*domain.User, error) {

	total, page, limit, users, err := util.FindAll[models.User](ctx, r.db)

	if err != nil {
		return 0, 0, 0, nil, err
	}

	return total, page, limit, r.modelsToDomain(users), nil
}

func (r *UserRepository) modelToDomain(userModel *models.User) *domain.User {
	return &domain.User{
		ID:                 userModel.ID,
		Email:              userModel.Email,
		FirstName:          userModel.FirstName,
		LastName:           userModel.LastName,
		DisplayName:        userModel.DisplayName,
		Bio:                userModel.Bio,
		Avatar:             userModel.Avatar,
		DateOfBirth:        userModel.DateOfBirth,
		Gender:             userModel.Gender,
		PhoneNumber:        userModel.PhoneNumber,
		LanguagePreference: userModel.LanguagePreference,
		TimeZone:           userModel.TimeZone,
		IsEmailVerified:    userModel.IsEmailVerified,
		IsPhoneVerified:    userModel.IsPhoneVerified,
		LastLoginAt:        userModel.LastLoginAt,
		CreatedAt:          userModel.CreatedAt,
		UpdatedAt:          userModel.UpdatedAt,
	}
}

func (r *UserRepository) modelsToDomain(userModels []models.User) []*domain.User {
	domainUsers := make([]*domain.User, len(userModels))
	for i, userModel := range userModels {
		domainUsers[i] = r.modelToDomain(&userModel)
	}
	return domainUsers
}
