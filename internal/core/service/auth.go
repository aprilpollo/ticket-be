package service

import (
	"errors"
	"task-management/internal/core/domain"
	"task-management/internal/core/port"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthService struct {
	authRepo port.AuthRepository
}

func NewAuthService(authRepo port.AuthRepository) *AuthService {
	return &AuthService{authRepo: authRepo}
}

func (s *AuthService) SignIn(ctx *fiber.Ctx, req *domain.SignInRequest) (*domain.AuthResponse, error) {
	// Get user by email
	user, err := s.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Validate password
	if err := s.authRepo.ValidatePassword(ctx, user.ID, req.Password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Update last login
	if err := s.authRepo.UpdateLastLogin(ctx, user.ID); err != nil {
		// Log error but don't fail the signin
		// TODO: Add proper logging
	}

	// Generate JWT tokens
	accessToken, refreshToken, expiresIn, err := s.authRepo.GenerateJWTToken(ctx, user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	return &domain.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

func (s *AuthService) SignUp(ctx *fiber.Ctx, req *domain.SignUpRequest) (*domain.AuthResponse, error) {
	// Check if user already exists
	existingUser, err := s.authRepo.GetUserByEmail(ctx, req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	// Create user
	user := &domain.User{
		Email:              req.Email,
		FirstName:          req.FirstName,
		LastName:           req.LastName,
		DisplayName:        req.DisplayName,
		LanguagePreference: "en",
		TimeZone:           "UTC",
		IsEmailVerified:    false,
		IsPhoneVerified:    false,
	}

	if err := s.authRepo.CreateUser(ctx, user, req.Password); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Generate unique slug for organization
	slug, err := s.authRepo.GenerateUniqueSlug(ctx)
	if err != nil {
		return nil, errors.New("failed to generate organization slug")
	}

	// Create organization
	organization := &domain.Organization{
		Name:        "MyOrganization",
		Slug:        slug,
		Description: "",
		PlanType:    "free",
		StatusID:    1, // Assuming 1 is active status
		Settings:    "{}",
	}

	if err := s.authRepo.CreateOrganization(ctx, organization); err != nil {
		return nil, errors.New("failed to create organization")
	}

	// Get default role
	defaultRole, err := s.authRepo.GetDefaultRole(ctx)
	if err != nil {
		return nil, errors.New("failed to get default role")
	}

	// Create organization member
	now := time.Now()
	organizationMember := &domain.OrganizationMember{
		OrganizationID: organization.ID,
		UserID:         user.ID,
		RoleID:         defaultRole.ID,
		StatusID:       1, // Assuming 1 is active status
		JoinedAt:       &now,
	}

	if err := s.authRepo.CreateOrganizationMember(ctx, organizationMember); err != nil {
		return nil, errors.New("failed to create organization member")
	}

	// Generate JWT tokens
	accessToken, refreshToken, expiresIn, err := s.authRepo.GenerateJWTToken(ctx, user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate tokens")
	}

	return &domain.AuthResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    expiresIn,
	}, nil
}

func (s *AuthService) ValidateToken(ctx *fiber.Ctx, tokenString string) (*domain.JWTClaims, error) {
	return s.authRepo.ValidateJWTToken(ctx, tokenString)
}
