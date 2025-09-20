package domain

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignUpRequest struct {
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	DisplayName string `json:"display_name" validate:"required"`
	Password    string `json:"password" validate:"required,min=6"`
}

type AuthResponse struct {
	User         *User  `json:"user"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type JWTClaims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type Organization struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"`
	Description string    `json:"description"`
	LogoURL     string    `json:"logo_url"`
	PlanType    string    `json:"plan_type"`
	StatusID    uint      `json:"status_id"`
	Settings    string    `json:"settings"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OrganizationMember struct {
	ID             uint       `json:"id"`
	OrganizationID uint       `json:"organization_id"`
	UserID         uint       `json:"user_id"`
	RoleID         uint       `json:"role_id"`
	StatusID       uint       `json:"status_id"`
	InvitedAt      *time.Time `json:"invited_at"`
	JoinedAt       *time.Time `json:"joined_at"`
	InvitedBy      *uint      `json:"invited_by"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
}

type OrganizationMemberRole struct {
	ID                    uint      `json:"id"`
	Name                  string    `json:"name"`
	Description           string    `json:"description"`
	IsDefault             bool      `json:"is_default"`
	IsPreview             bool      `json:"is_preview"`
	CanManageOrganization bool      `json:"can_manage_organization"`
	CanManageMembers      bool      `json:"can_manage_members"`
	CanManageProjects     bool      `json:"can_manage_projects"`
	CanCreateProjects     bool      `json:"can_create_projects"`
	CanViewAllProjects    bool      `json:"can_view_all_projects"`
	CanManageTasks        bool      `json:"can_manage_tasks"`
	CanViewReports        bool      `json:"can_view_reports"`
}

type OrganizationWithMember struct {
	Organization *Organization           `json:"organization"`
	Member       *OrganizationMember     `json:"member"`
	Role         *OrganizationMemberRole `json:"role"`
}


type ValidateUserRequest struct {
	Uuid         string `json:"uuid" bson:"uuid"`
	BusinessCode string `json:"business_code" bson:"business_code"`
	ProfileImage string `json:"profile_image" bson:"profile_image"`
	Role         string `json:"role" bson:"role"`
	Email        string `json:"email" bson:"email"`
	Active       bool   `json:"active" bson:"active"`
}