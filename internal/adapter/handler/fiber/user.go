package httpfiber

import (
	"strconv"
	"task-management/internal/core/port"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}


// GetUserByID retrieves a user by ID
// @Summary Get user by ID
// @Description Get a specific user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]interface{} "User not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users/{id} [get]
func (h *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	idStr := ctx.Params("id")
	if idStr == "" {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "id parameter is missing", nil)
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "id must be a valid number", nil)
	}

	user, err := h.userService.GetUserByID(ctx, uint(id))
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}

	return ResData(ctx, fiber.StatusOK, "SUCCESS", "", user)
}

// GetAllUsers retrieves all users with pagination
// @Summary Get all users
// @Description Get a list of all users with pagination support
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Number of users to return (default: 10, max: 100)"
// @Param offset query int false "Number of users to skip (default: 0)"
// @Success 200 {object} map[string]interface{} "Users retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /users [get]
func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {
	total, page, limit, users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}
	return ResData(ctx, 200, "SUCCESS", "", users, int(total), int(page), int(limit))
}
