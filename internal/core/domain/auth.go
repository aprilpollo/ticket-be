package domain

import (

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




type ValidateUserRequest struct {
	Uuid         string `json:"uuid" bson:"uuid"`
	BusinessCode string `json:"business_code" bson:"business_code"`
	ProfileImage string `json:"profile_image" bson:"profile_image"`
	Role         string `json:"role" bson:"role"`
	Email        string `json:"email" bson:"email"`
	Active       bool   `json:"active" bson:"active"`
}