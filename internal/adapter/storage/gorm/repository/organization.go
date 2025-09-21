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

func (r *OrganizationRepository) GetOrganization(ctx *fiber.Ctx) (int64, int64, int64, []*domain.Organization, error) {
	total, page, limit, organizations, err := util.FindAll[models.Organization](ctx, r.db)
	if err != nil {
		return 0, 0, 0, nil, err
	}
	return total, page, limit, r.modelsToDomain(organizations), nil
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

func (r *OrganizationRepository) GetUserRoleInOrganizationByID(ctx *fiber.Ctx, orgId, userId uint) (*domain.OrganizationMemberRole, error) {
	var role models.VWOrganizationMemberRole

	err := r.db.Where("organization_id = ? AND user_id = ?", orgId, userId).First(&role).Error
	if err != nil {
		return nil, err
	}

	roleModel := &domain.OrganizationMemberRole{
		ID:                    role.OrganizationID,
		Name:                  role.Name,
		Description:           role.Description,
		IsDefault:             role.IsDefault,
		IsPreview:             role.IsPreview,
		CanManageOrganization: role.CanManageOrganization,
		CanManageMembers:      role.CanManageMembers,
		CanManageProjects:     role.CanManageProjects,
		CanCreateProjects:     role.CanCreateProjects,
		CanViewAllProjects:    role.CanViewAllProjects,
		CanManageTasks:        role.CanManageTasks,
		CanViewReports:        role.CanViewReports,
	}
	return roleModel, nil
}

func (r *OrganizationRepository) UpdateOrganization(ctx *fiber.Ctx, id uint, organization *domain.UpdateOrganizationRequest) (*domain.Organization, error) {
	organizationModel := models.Organization{
		Name:        organization.Name,
		Description: organization.Description,
		LogoURL:     organization.LogoURL,
		PlanType:    organization.PlanType,
		StatusID:    organization.StatusID,
		Settings:    organization.Settings,
	}

	resault, err := util.UpdateOne[models.Organization](ctx, r.db, int64(id), organizationModel)

	if err != nil {
		return nil, err
	}

	return r.modelToDomain(resault), nil
}

func (r *OrganizationRepository) modelToDomain(model *models.Organization) *domain.Organization {
	return &domain.Organization{
		ID:          model.ID,
		Name:        model.Name,
		Slug:        model.Slug,
		Description: model.Description,
		LogoURL:     model.LogoURL,
		PlanType:    model.PlanType,
		StatusID:    model.StatusID,
		Settings:    model.Settings,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func (r *OrganizationRepository) modelsToDomain(models []models.Organization) []*domain.Organization {
	organizations := make([]*domain.Organization, len(models))
	for i, model := range models {
		organizations[i] = r.modelToDomain(&model)
	}
	return organizations
}
