package repository

import (
	"task-management/internal/adapter/storage/gorm/models"
	"task-management/internal/core/domain"

	"gorm.io/gorm"

	"task-management/internal/util"

	"github.com/gofiber/fiber/v2"
)

type OrganizationRepository struct {
	db *gorm.DB
}

func NewOrganizationRepository(db *gorm.DB) *OrganizationRepository {
	return &OrganizationRepository{db: db}
}

func (r *OrganizationRepository) GetUserRoleInOrganization(ctx *fiber.Ctx, userID uint) (int64, int64, int64, []*domain.OrganizationMemberRole, error) {
	total, page, limit, role, err := util.FindAllByCondition[models.VWOrganizationMemberRole](ctx, r.db, "user_id", userID)

	if err != nil {
		return 0, 0, 0, nil, err
	}

	roleModels := make([]*domain.OrganizationMemberRole, len(role))
	for i, roleModel := range role {
		roleModels[i] = &domain.OrganizationMemberRole{
			ID:                    roleModel.OrganizationID,
			Name:                  roleModel.Name,
			Description:           roleModel.Description,
			IsDefault:             roleModel.IsDefault,
			IsPreview:             roleModel.IsPreview,
			CanManageOrganization: roleModel.CanManageOrganization,
			CanManageMembers:      roleModel.CanManageMembers,
			CanManageProjects:     roleModel.CanManageProjects,
			CanCreateProjects:     roleModel.CanCreateProjects,
			CanViewAllProjects:    roleModel.CanViewAllProjects,
			CanManageTasks:        roleModel.CanManageTasks,
			CanViewReports:        roleModel.CanViewReports,
		}
	}

	return total, page, limit, roleModels, nil
}
