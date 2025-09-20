package repository

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"task-management/internal/adapter/storage/gorm/models"
	"task-management/internal/core/domain"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	config "task-management/internal/adapter/config"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(ctx *fiber.Ctx, user *domain.User, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

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

	authMethod := models.UserAuthMethod{
		UserID:       userModel.ID,
		AuthType:     "password",
		IsPrimary:    true,
		PasswordHash: string(hashedPassword),
	}

	if err := r.db.Create(&authMethod).Error; err != nil {
		return err
	}

	user.ID = userModel.ID
	user.CreatedAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt

	return nil
}

func (r *AuthRepository) GetUserByEmail(ctx *fiber.Ctx, email string) (*domain.User, error) {
	var userModel models.User
	if err := r.db.Where("email = ?", email).First(&userModel).Error; err != nil {
		return nil, err
	}
	return r.userModelToDomain(&userModel), nil
}

func (r *AuthRepository) ValidatePassword(ctx *fiber.Ctx, userID uint, password string) error {
	var authMethod models.UserAuthMethod
	if err := r.db.Where("user_id = ? AND auth_type = ?", userID, "password").First(&authMethod).Error; err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword([]byte(authMethod.PasswordHash), []byte(password))
}

func (r *AuthRepository) UpdateLastLogin(ctx *fiber.Ctx, userID uint) error {
	now := time.Now()
	return r.db.Model(&models.User{}).Where("id = ?", userID).Update("last_login_at", now).Error
}

func (r *AuthRepository) CreateOrganization(ctx *fiber.Ctx, org *domain.Organization) error {
	orgModel := models.Organization{
		Name:        org.Name,
		Slug:        org.Slug,
		Description: org.Description,
		LogoURL:     org.LogoURL,
		PlanType:    org.PlanType,
		StatusID:    org.StatusID,
		Settings:    org.Settings,
	}

	if err := r.db.Create(&orgModel).Error; err != nil {
		return err
	}

	org.ID = orgModel.ID
	org.CreatedAt = orgModel.CreatedAt
	org.UpdatedAt = orgModel.UpdatedAt

	return nil
}

func (r *AuthRepository) CreateOrganizationMember(ctx *fiber.Ctx, member *domain.OrganizationMember) error {
	memberModel := models.OrganizationMember{
		OrganizationID: member.OrganizationID,
		UserID:         member.UserID,
		RoleID:         member.RoleID,
		StatusID:       member.StatusID,
		InvitedAt:      member.InvitedAt,
		JoinedAt:       member.JoinedAt,
		InvitedBy:      member.InvitedBy,
	}

	if err := r.db.Create(&memberModel).Error; err != nil {
		return err
	}

	member.ID = memberModel.ID
	member.CreatedAt = memberModel.CreatedAt
	member.UpdatedAt = memberModel.UpdatedAt

	return nil
}

func (r *AuthRepository) GetDefaultRole(ctx *fiber.Ctx) (*domain.OrganizationMemberRole, error) {
	var roleModel models.OrganizationMemberRole
	if err := r.db.Where("is_default = ?", true).First(&roleModel).Error; err != nil {
		return nil, err
	}
	return r.roleModelToDomain(&roleModel), nil
}

func (r *AuthRepository) GenerateUniqueSlug(ctx *fiber.Ctx) (string, error) {
	// Generate 10-digit random number
	var slug string
	for {
		// Generate random 10-digit number
		randomNum, err := rand.Int(rand.Reader, big.NewInt(9999999999))
		if err != nil {
			return "", err
		}

		// Format as 10-digit string with leading zeros
		slug = fmt.Sprintf("%010d", randomNum.Int64())

		// Check if slug already exists
		var count int64
		if err := r.db.Model(&models.Organization{}).Where("slug = ?", slug).Count(&count).Error; err != nil {
			return "", err
		}

		if count == 0 {
			break
		}
	}

	return slug, nil
}

func (r *AuthRepository) GenerateJWTToken(ctx *fiber.Ctx, userID uint, email string) (string, string, int64, error) {
	secretKey := []byte(config.Env.JWT.SecretKey)

	accessExpirationTime := time.Now().Add(time.Duration(config.Env.JWT.JwtExpireDaysCount) * 24 * time.Hour)
	accessClaims := &domain.JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", "", 0, err
	}

	refreshExpirationTime := time.Now().Add(30 * 24 * time.Hour)
	refreshClaims := &domain.JWTClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", "", 0, err
	}

	expiresIn := int64(time.Until(accessExpirationTime).Seconds())

	return accessTokenString, refreshTokenString, expiresIn, nil
}

func (r *AuthRepository) ValidateJWTToken(ctx *fiber.Ctx, tokenString string) (*domain.JWTClaims, error) {
	secretKey := []byte(config.Env.JWT.SecretKey) 

	token, err := jwt.ParseWithClaims(tokenString, &domain.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*domain.JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}


func (r *AuthRepository) userModelToDomain(userModel *models.User) *domain.User {
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

func (r *AuthRepository) roleModelToDomain(roleModel *models.OrganizationMemberRole) *domain.OrganizationMemberRole {
	return &domain.OrganizationMemberRole{
		ID:                    roleModel.ID,
		Name:                  roleModel.Name,
		Description:           roleModel.Description,
		IsDefault:             roleModel.IsDefault,
		IsPreview:             roleModel.IsPreview,
		CanManageOrganization: roleModel.CanManageOrganization,
		CanManageMembers:      roleModel.CanManageMembers,
		CanManageProjects:     roleModel.CanManageProjects,
		CanCreateProjects:     roleModel.CanCreateProjects,
		CanViewAllProjects:    roleModel.CanViewAllProjects,
		CanManageTasks:        roleModel.CanManageTasks,
		CanViewReports:        roleModel.CanViewReports,
	}
}
