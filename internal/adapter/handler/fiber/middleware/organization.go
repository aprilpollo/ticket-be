package middleware

import (
	"strconv"
	"task-management/internal/adapter/handler/fiber/routes"
	"task-management/internal/core/domain"
	"task-management/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type OrganizationMiddleware struct {
	organizationService port.OrganizationService
}

func NewOrganizationMiddleware(organizationService port.OrganizationService) *OrganizationMiddleware {
	return &OrganizationMiddleware{
		organizationService: organizationService,
	}
}

func (m *OrganizationMiddleware) Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, ok := c.Locals("user_id").(uint)
		if !ok {
			return routes.ResData(c, fiber.StatusUnauthorized, "UNAUTHORIZED", "User ID not found in context", nil)
		}

		orgIDHeader := c.Get("X-Organization-ID")
		if orgIDHeader == "" {
			return routes.ResData(c, fiber.StatusBadRequest, "BAD REQUEST", "X-Organization-ID header is required", nil)
		}

		orgID, err := strconv.ParseUint(orgIDHeader, 10, 32)
		if err != nil {
			return routes.ResData(c, fiber.StatusBadRequest, "BAD REQUEST", "Invalid organization ID in header", nil)
		}

		hasAccess, userRole, err := m.checkUserAccessAndGetRole(c, userID, uint(orgID))
		if err != nil {
			return routes.ResData(c, fiber.StatusInternalServerError, "INTERNAL ERROR", "Failed to check user access", nil)
		}

		if !hasAccess {
			return routes.ResData(c, fiber.StatusForbidden, "FORBIDDEN", "You don't have access to this organization", nil)
		}

		c.Locals("organization_id", uint(orgID))
		c.Locals("user_role", userRole)

		return c.Next()
	}

}

func (m *OrganizationMiddleware) checkUserAccessAndGetRole(ctx *fiber.Ctx, userID uint, organizationID uint) (bool, *domain.OrganizationMemberRole, error) {
	userRole, err := m.organizationService.GetUserRoleInOrganizationByID(ctx, organizationID, userID)
	if err != nil {
		return false, nil, err
	}

	return true, userRole, nil
}

func (m *OrganizationMiddleware) MiddlewareWithPermission(permission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRole, ok := c.Locals("user_role").(*domain.OrganizationMemberRole)
		if !ok {
			return routes.ResData(c, fiber.StatusInternalServerError, "INTERNAL ERROR", "User role not found in context", nil)
		}

		if !m.hasPermission(userRole, permission) {
			return routes.ResData(c, fiber.StatusForbidden, "FORBIDDEN", "You don't have permission to perform this action", nil)
		}

		return c.Next()
	}
}

func (m *OrganizationMiddleware) hasPermission(role *domain.OrganizationMemberRole, permission string) bool {

	switch permission {
		case "IsPreview":
			return role.IsPreview
		case "CanManageOrganization":
			return role.CanManageOrganization
		case "CanManageMembers":
			return role.CanManageMembers
		case "CanManageProjects":
			return role.CanManageProjects
		case "CanCreateProjects":
			return role.CanCreateProjects
		case "CanViewAllProjects":
			return role.CanViewAllProjects
		case "CanManageTasks":
			return role.CanManageTasks
		case "CanViewReports":
			return role.CanViewReports
		default:
			return false
	}
}
