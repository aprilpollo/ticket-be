package domain

import "time"

type Account struct {
	ID                 uint       `json:"id"`
	Email              string     `json:"email"`
	FirstName          string     `json:"first_name"`
	LastName           string     `json:"last_name"`
	DisplayName        string     `json:"display_name"`
	Bio                string     `json:"bio"`
	Avatar             string     `json:"avatar"`
	DateOfBirth        *time.Time `json:"date_of_birth"`
	Gender             string     `json:"gender"`
	PhoneNumber        string     `json:"phone_number"`
	LanguagePreference string     `json:"language_preference"`
	TimeZone           string     `json:"time_zone"`
	IsEmailVerified    bool       `json:"is_email_verified"`
	IsPhoneVerified    bool       `json:"is_phone_verified"`
}

type RoleInOrganization struct {
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