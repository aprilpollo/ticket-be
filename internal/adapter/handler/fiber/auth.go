package httpfiber

import (
	"task-management/internal/core/domain"
	"task-management/internal/core/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService port.AuthService
	validate    *validator.Validate
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		validate:    validator.New(),
	}
}

// SignIn handles user sign in
func (h *AuthHandler) SignIn(c *fiber.Ctx) error {
	var req domain.SignInRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return ResData(c, fiber.StatusBadRequest, "BAD REQUEST", "Invalid request body", nil)
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		return ResData(c, fiber.StatusBadRequest, "BAD REQUEST", "Validation failed: "+err.Error(), nil)
	}

	// Call service
	response, err := h.authService.SignIn(c, &req)
	if err != nil {
		return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", err.Error(), nil)
	}

	return ResData(c, fiber.StatusOK, "SUCCESS", "", response)
}

// SignUp handles user sign up
func (h *AuthHandler) SignUp(c *fiber.Ctx) error {
	var req domain.SignUpRequest

	// Parse request body
	if err := c.BodyParser(&req); err != nil {
		return ResData(c, fiber.StatusBadRequest, "BAD REQUEST", "Invalid request body", nil)
	}

	// Validate request
	if err := h.validate.Struct(&req); err != nil {
		return ResData(c, fiber.StatusBadRequest, "BAD REQUEST", "Validation failed: "+err.Error(), nil)
	}

	// Call service
	response, err := h.authService.SignUp(c, &req)
	if err != nil {
		return ResData(c, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}

	return ResData(c, fiber.StatusCreated, "SUCCESS", "", response)
}

// ValidateToken handles token validation
func (h *AuthHandler) ValidateToken(c *fiber.Ctx) error {
	// Get token from header
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Authorization header is required", nil)
	}

	// Extract token from "Bearer <token>"
	tokenString := ""
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		tokenString = authHeader[7:]
	} else {
		return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Invalid authorization header format", nil)
	}

	// Validate token
	claims, err := h.authService.ValidateToken(c, tokenString)
	if err != nil {
		return ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "Invalid token", nil)
	}

	return ResData(c, fiber.StatusOK, "SUCCESS", "", fiber.Map{
		"user_id": claims.UserID,
		"email":   claims.Email,
	})
}

func (h *AuthHandler) ValidateUser(c *fiber.Ctx) error {
	response := domain.ValidateUserRequest{
		Uuid: "#",
		Email: c.Locals("email").(string),
		Active: true,
		Role: "owner",
		ProfileImage: "#",
		BusinessCode: "#",
	}

	return ResData(c, fiber.StatusOK, "SUCCESS", "", response)

}