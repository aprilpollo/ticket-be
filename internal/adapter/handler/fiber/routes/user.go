package routes

import (
	"strconv"
	"task-management/internal/core/port"
	"task-management/internal/core/domain"
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	userService port.UserService
	validate    *validator.Validate
}

func NewUserHandler(userService port.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
		validate:    validator.New(),
	}
}


func (h *UserHandler) GetAllUsers(ctx *fiber.Ctx) error {
	total, page, limit, users, err := h.userService.GetAllUsers(ctx)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}
	return ResData(ctx, 200, "SUCCESS", "", users, int(total), int(page), int(limit))
}

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

func (h *UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "id must be a valid number", nil)
	}
	
	var req domain.UpdateUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "Invalid request body", nil)
	}
	
	if err := h.validate.Struct(&req); err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "Validation failed: "+err.Error(), nil)
	}
	
	user, err := h.userService.UpdateUser(ctx, uint(id), &req)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}
	return ResData(ctx, fiber.StatusOK, "SUCCESS", "", user)
}

func (h *UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	return nil
}
