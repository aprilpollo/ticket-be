package routes

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
	"task-management/internal/core/domain"
	"task-management/internal/core/port"
)

type OrganizationHandler struct {
	organizationService port.OrganizationService
	validate            *validator.Validate
}

func NewOrganizationHandler(organizationService port.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		organizationService: organizationService,
		validate:            validator.New(),
	}
}

func (h *OrganizationHandler) GetOrganization(ctx *fiber.Ctx) error {
	total, page, limit, organizations, err := h.organizationService.GetOrganization(ctx)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}
	return ResData(ctx, fiber.StatusOK, "SUCCESS", "", organizations, int(total), int(page), int(limit))
}

func (h *OrganizationHandler) GetUserRoleInOrganization(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	total, page, limit, role, err := h.organizationService.GetUserRoleInOrganization(ctx, userID)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}

	return ResData(ctx, fiber.StatusOK, "SUCCESS", "", role, int(total), int(page), int(limit))

}

func (h *OrganizationHandler) UpdateOrganization(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "id must be a valid number", nil)
	}

	var req domain.UpdateOrganizationRequest

	if err = ctx.BodyParser(&req); err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "Invalid request body", nil)
	}

	if err = h.validate.Struct(&req); err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", "Validation failed: "+err.Error(), nil)
	}

	organization, err := h.organizationService.UpdateOrganization(ctx, uint(id), &req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ResData(ctx, fiber.StatusNotFound, "NOT FOUND", "organization not found", nil)
		}
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}

	return ResData(ctx, fiber.StatusOK, "SUCCESS", "", organization)
}
