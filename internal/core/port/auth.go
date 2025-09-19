package port

import (
	"task-management/internal/core/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthRepository interface {
	// User operations
	CreateUser(ctx *fiber.Ctx, user *domain.User, password string) error
	GetUserByEmail(ctx *fiber.Ctx, email string) (*domain.User, error)
	ValidatePassword(ctx *fiber.Ctx, userID uint, password string) error
	UpdateLastLogin(ctx *fiber.Ctx, userID uint) error

	// Organization operations
	CreateOrganization(ctx *fiber.Ctx, org *domain.Organization) error
	CreateOrganizationMember(ctx *fiber.Ctx, member *domain.OrganizationMember) error
	GetDefaultRole(ctx *fiber.Ctx) (*domain.OrganizationMemberRole, error)
	GenerateUniqueSlug(ctx *fiber.Ctx) (string, error)

	// JWT operations
	GenerateJWTToken(ctx *fiber.Ctx, userID uint, email string) (string, string, int64, error)
	ValidateJWTToken(ctx *fiber.Ctx, tokenString string) (*domain.JWTClaims, error)
}

type AuthService interface {
	SignIn(ctx *fiber.Ctx, req *domain.SignInRequest) (*domain.AuthResponse, error)
	SignUp(ctx *fiber.Ctx, req *domain.SignUpRequest) (*domain.AuthResponse, error)
	ValidateToken(ctx *fiber.Ctx, tokenString string) (*domain.JWTClaims, error)
}
