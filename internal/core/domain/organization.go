package domain

import "time"

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

type UpdateOrganizationRequest struct {
	Name        string    `json:"name" validate:"omitempty,min=3,max=20"`
	Description string    `json:"description" validate:"omitempty,min=3,max=200"`
	LogoURL     string    `json:"logo_url" validate:"omitempty,url"`
	PlanType    string    `json:"plan_type" validate:"omitempty,oneof=free pro enterprise"`
	StatusID    uint      `json:"status_id" validate:"omitempty,min=1"`
	Settings    string    `json:"settings"`
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