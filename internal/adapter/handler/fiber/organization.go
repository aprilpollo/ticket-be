package httpfiber

import (
	"task-management/internal/core/port"
	"github.com/gofiber/fiber/v2"

)

type OrganizationHandler struct {
	organizationService port.OrganizationService
}

func NewOrganizationHandler(organizationService port.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		organizationService: organizationService,
	}
}

func (h *OrganizationHandler) GetUserRoleInOrganization(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(uint)

	total, page, limit, role, err := h.organizationService.GetUserRoleInOrganization(ctx, userID)
	if err != nil {
		return ResData(ctx, fiber.StatusBadRequest, "BAD REQUEST", err.Error(), nil)
	}

	return ResData(ctx, 200, "SUCCESS", "", role, int(total), int(page), int(limit))

}