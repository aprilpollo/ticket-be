package domain

import "time"

type User struct {
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
	LastLoginAt        *time.Time `json:"last_login_at"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
}

type CreateUserRequest struct {
	Email              string     `json:"email" validate:"required,email"`
	FirstName          string     `json:"first_name" validate:"required"`
	LastName           string     `json:"last_name" validate:"required"`
	DisplayName        string     `json:"display_name"`
	Bio                string     `json:"bio"`
	Avatar             string     `json:"avatar"`
	DateOfBirth        *time.Time `json:"date_of_birth"`
	Gender             string     `json:"gender"`
	PhoneNumber        string     `json:"phone_number"`
	LanguagePreference string     `json:"language_preference"`
	TimeZone           string     `json:"time_zone"`
}

type UpdateUserRequest struct {
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	DisplayName        string    `json:"display_name"`
	Bio                string    `json:"bio"`
	Avatar             string    `json:"avatar"`
	DateOfBirth        time.Time `json:"date_of_birth"`
	Gender             string    `json:"gender"`
	PhoneNumber        string    `json:"phone_number"`
	LanguagePreference string    `json:"language_preference"`
	TimeZone           string    `json:"time_zone"`
}
