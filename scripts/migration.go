package main

import (
	config "task-management/internal/adapter/config"

	"fmt"
	"log"
	"strings"
	"task-management/internal/adapter/storage/gorm"
	"task-management/internal/adapter/storage/gorm/models"

	"github.com/fatih/color"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	blue  = color.New(color.FgBlue).SprintFunc()
	cyan  = color.New(color.FgCyan).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

func printSuccess(msg string) {
	fmt.Printf("%s %s\n", green("[v]"), msg)
}

func printInfo(msg string) {
	fmt.Printf("%s %s\n", blue("[INFO]"), msg)
}

func printHeader(msg string) {
	fmt.Printf("\n%s %s\n", cyan("[INFO]"), msg)
	fmt.Printf("%s\n", cyan(strings.Repeat("=", 50)))
}

func main() {
	printHeader("Starting Database Migration")

	printInfo("Loading configuration...")
	config.LoadConfig()

	printInfo("Connecting to database...")
	if err := gormOrm.Init(config.Env.Postgre.URI, 1, 1, 0); err != nil {
		log.Fatalf("%s failed to connect DB: %v", red("[x]"), err)
	}
	printSuccess("Database connected successfully")

	printInfo("Running auto migration...")
	if err := gormOrm.Trx.AutoMigrate(models.All()...); err != nil {
		log.Fatalf("%s migration failed: %v", red("[FAILED]"), err)
	}

	// printInfo("Dropping existing views...")
	// for name := range views.Views {
	// 	dropSQL := fmt.Sprintf("DROP VIEW IF EXISTS %s CASCADE;", name)
	// 	if err := gormOrm.Trx.Exec(dropSQL).Error; err != nil {
	// 		log.Fatalf("%s failed to drop view %s: %v", red("[x]"), name, err)
	// 	}
	// 	printSuccess(fmt.Sprintf("Dropped view: %s", name))
	// }

	// printInfo("Creating new views...")
	// for name, query := range views.Views {
	// 	if err := gormOrm.Trx.Exec(query).Error; err != nil {
	// 		log.Fatalf("%s failed to create view %s: %v", red("[x]"), name, err)
	// 	}
	// 	printSuccess(fmt.Sprintf("Created view: %s", name))
	// }

	upsertDefaultOrganization()
	fmt.Println()
	upsertDefaultProject()
	fmt.Println()
	upsertDefaultPriority()

	printHeader("Migration Completed Successfully!")
}

func upsertDefaultOrganization() {
	// Seed data for OrganizationStatus
	organizationStatuses := []*models.OrganizationStatus{
		{
			ID:          1,
			Name:        "Active",
			Description: "Organization is active and fully operational",
		},
		{
			ID:          2,
			Name:        "Suspended",
			Description: "Organization is temporarily suspended",
		},
		{
			ID:          3,
			Name:        "Inactive",
			Description: "Organization is inactive but can be reactivated",
		},
		{
			ID:          4,
			Name:        "Deleted",
			Description: "Organization is marked for deletion",
		},
		{
			ID:          5,
			Name:        "Pending",
			Description: "Organization setup is pending completion",
		},
	}

	// Seed data for MemberStatus
	memberStatuses := []*models.MemberStatus{
		{
			ID:          1,
			Name:        "Active",
			Description: "Member is active and has full access",
		},
		{
			ID:          2,
			Name:        "Invited",
			Description: "Member has been invited but hasn't joined yet",
		},
		{
			ID:          3,
			Name:        "Pending",
			Description: "Member invitation is pending approval",
		},
		{
			ID:          4,
			Name:        "Suspended",
			Description: "Member is temporarily suspended",
		},
		{
			ID:          5,
			Name:        "Inactive",
			Description: "Member is inactive but can be reactivated",
		},
		{
			ID:          6,
			Name:        "Left",
			Description: "Member has left the organization",
		},
		{
			ID:          7,
			Name:        "Removed",
			Description: "Member was removed from the organization",
		},
	}

	// Seed data for OrganizationMemberRole
	organizationMemberRoles := []*models.OrganizationMemberRole{
		{
			ID:          1,
			Name:        "Owner",
			Description: "Organization owner with full access to all features",
			IsDefault:   false,
			IsPreview:   false,

			// Full permissions
			CanManageOrganization: true,
			CanManageMembers:      true,
			CanManageProjects:     true,
			CanCreateProjects:     true,
			CanViewAllProjects:    true,
			CanManageTasks:        true,
			CanViewReports:        true,
		},
		{
			ID:          2,
			Name:        "Admin",
			Description: "Administrator with management permissions",
			IsDefault:   false,
			IsPreview:   false,

			// Most permissions except organization management
			CanManageOrganization: false,
			CanManageMembers:      true,
			CanManageProjects:     true,
			CanCreateProjects:     true,
			CanViewAllProjects:    true,
			CanManageTasks:        true,
			CanViewReports:        true,
		},
		{
			ID:          3,
			Name:        "Project Manager",
			Description: "Can manage projects and tasks",
			IsDefault:   false,
			IsPreview:   false,

			// Project and task management
			CanManageOrganization: false,
			CanManageMembers:      false,
			CanManageProjects:     true,
			CanCreateProjects:     true,
			CanViewAllProjects:    true,
			CanManageTasks:        true,
			CanViewReports:        true,
		},
		{
			ID:          4,
			Name:        "Member",
			Description: "Standard member with basic access",
			IsDefault:   true,
			IsPreview:   false,

			// Basic permissions
			CanManageOrganization: false,
			CanManageMembers:      false,
			CanManageProjects:     false,
			CanCreateProjects:     false,
			CanViewAllProjects:    false,
			CanManageTasks:        true,
			CanViewReports:        false,
		},
		{
			ID:          5,
			Name:        "Viewer",
			Description: "Read-only access to assigned projects",
			IsDefault:   false,
			IsPreview:   true,

			// View only permissions
			CanManageOrganization: false,
			CanManageMembers:      false,
			CanManageProjects:     false,
			CanCreateProjects:     false,
			CanViewAllProjects:    false,
			CanManageTasks:        false,
			CanViewReports:        false,
		},
		{
			ID:          6,
			Name:        "Guest",
			Description: "Limited access for external collaborators",
			IsDefault:   false,
			IsPreview:   true,

			// Very limited permissions
			CanManageOrganization: false,
			CanManageMembers:      false,
			CanManageProjects:     false,
			CanCreateProjects:     false,
			CanViewAllProjects:    false,
			CanManageTasks:        false,
			CanViewReports:        false,
		},
	}

	printInfo("Seeding organization statuses...")
	if err := gormOrm.Trx.Where("id IN ?", []uint{1, 2, 3, 4, 5}).Delete(&models.OrganizationStatus{}).Error; err != nil {
		log.Fatalf("%s failed to delete existing organization statuses: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Where("id IN ?", []uint{1, 2, 3, 4, 5, 6, 7}).Delete(&models.MemberStatus{}).Error; err != nil {
		log.Fatalf("%s failed to delete existing member statuses: %v", red("[x]"), err)
	}

	if err := gormOrm.Trx.Where("id IN ?", []uint{1, 2, 3, 4, 5, 6}).Delete(&models.OrganizationMemberRole{}).Error; err != nil {
		log.Fatalf("%s failed to delete existing organization member roles: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Create(organizationStatuses).Error; err != nil {
		log.Fatalf("%s failed to create default organization statuses: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Create(memberStatuses).Error; err != nil {
		log.Fatalf("%s failed to create default member statuses: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Create(organizationMemberRoles).Error; err != nil {
		log.Fatalf("%s failed to create default organization member roles: %v", red("[x]"), err)
	}

	printSuccess(fmt.Sprintf("Created %d organization statuses", len(organizationStatuses)))
	printSuccess(fmt.Sprintf("Created %d member statuses", len(memberStatuses)))
	printSuccess(fmt.Sprintf("Created %d organization member roles", len(organizationMemberRoles)))
}

func upsertDefaultProject() {
	// Seed data for ProjectStatus
	projectStatuses := []*models.ProjectStatus{
		{
			ID:          1,
			Name:        "Active",
			Description: "Project is active and in development",
		},
		{
			ID:          2,
			Name:        "On Hold",
			Description: "Project is temporarily paused",
		},
		{
			ID:          3,
			Name:        "Completed",
			Description: "Project has been completed successfully",
		},
		{
			ID:          4,
			Name:        "Cancelled",
			Description: "Project has been cancelled",
		},
		{
			ID:          5,
			Name:        "Planning",
			Description: "Project is in planning phase",
		},
		{
			ID:          6,
			Name:        "Archived",
			Description: "Project is archived for reference",
		},
	}

	// Seed data for ProjectMemberRole
	projectMemberRoles := []*models.ProjectMemberRole{
		{
			ID:          1,
			Name:        "Project Manager",
			Description: "Full project management access with all permissions",
			IsDefault:   false,

			// Full project permissions
			CanManageProject:    true,
			CanManageMembers:    true,
			CanCreateTasks:      true,
			CanManageTasks:      true,
			CanDeleteTasks:      true,
			CanViewAllTasks:     true,
			CanManageComponents: true,
			CanManageVersions:   true,
		},
		{
			ID:          2,
			Name:        "Lead Developer",
			Description: "Senior developer with management permissions",
			IsDefault:   false,

			// Most permissions except project management
			CanManageProject:    false,
			CanManageMembers:    true,
			CanCreateTasks:      true,
			CanManageTasks:      true,
			CanDeleteTasks:      true,
			CanViewAllTasks:     true,
			CanManageComponents: true,
			CanManageVersions:   true,
		},
		{
			ID:          3,
			Name:        "Developer",
			Description: "Standard developer with task management access",
			IsDefault:   true,

			// Task management permissions
			CanManageProject:    false,
			CanManageMembers:    false,
			CanCreateTasks:      true,
			CanManageTasks:      true,
			CanDeleteTasks:      false,
			CanViewAllTasks:     true,
			CanManageComponents: false,
			CanManageVersions:   false,
		},
		{
			ID:          4,
			Name:        "Tester",
			Description: "Quality assurance with testing permissions",
			IsDefault:   false,

			// Testing focused permissions
			CanManageProject:    false,
			CanManageMembers:    false,
			CanCreateTasks:      true,
			CanManageTasks:      true,
			CanDeleteTasks:      false,
			CanViewAllTasks:     true,
			CanManageComponents: false,
			CanManageVersions:   false,
		},
		{
			ID:          5,
			Name:        "Reporter",
			Description: "Can create and view tasks but limited management",
			IsDefault:   false,

			// Limited permissions
			CanManageProject:    false,
			CanManageMembers:    false,
			CanCreateTasks:      true,
			CanManageTasks:      false,
			CanDeleteTasks:      false,
			CanViewAllTasks:     true,
			CanManageComponents: false,
			CanManageVersions:   false,
		},
		{
			ID:          6,
			Name:        "Viewer",
			Description: "Read-only access to project tasks",
			IsDefault:   false,

			// View only permissions
			CanManageProject:    false,
			CanManageMembers:    false,
			CanCreateTasks:      false,
			CanManageTasks:      false,
			CanDeleteTasks:      false,
			CanViewAllTasks:     true,
			CanManageComponents: false,
			CanManageVersions:   false,
		},
	}

	printInfo("Seeding project statuses...")
	if err := gormOrm.Trx.Where("id IN ?", []uint{1, 2, 3, 4, 5, 6}).Delete(&models.ProjectStatus{}).Error; err != nil {
		log.Fatalf("%s failed to delete existing project statuses: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Create(projectStatuses).Error; err != nil {
		log.Fatalf("%s failed to create default project statuses: %v", red("[x]"), err)
	}

	if err := gormOrm.Trx.Where("id IN ?", []uint{1, 2, 3, 4, 5, 6}).Delete(&models.ProjectMemberRole{}).Error; err != nil {
		log.Fatalf("%s failed to delete existing project member roles: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Create(projectMemberRoles).Error; err != nil {
		log.Fatalf("%s failed to create default project member roles: %v", red("[x]"), err)
	}

	printSuccess(fmt.Sprintf("Created %d project statuses", len(projectStatuses)))
	printSuccess(fmt.Sprintf("Created %d project member roles", len(projectMemberRoles)))
}

func upsertDefaultPriority() {
	// Seed data for Priority
	priorities := []*models.Priority{
		{
			ID:          1,
			Name:        "Lowest",
			Description: "Lowest priority - can be addressed when time permits",
			Color:       "#57D9A3",
			Level:       1,
			IsActive:    true,
		},
		{
			ID:          2,
			Name:        "Low",
			Description: "Low priority - minor issues that don't block progress",
			Color:       "#79F2C0",
			Level:       2,
			IsActive:    true,
		},
		{
			ID:          3,
			Name:        "Medium",
			Description: "Medium priority - standard priority for most issues",
			Color:       "#2684FF",
			Level:       3,
			IsActive:    true,
		},
		{
			ID:          4,
			Name:        "High",
			Description: "High priority - important issues that should be addressed soon",
			Color:       "#FF8B00",
			Level:       4,
			IsActive:    true,
		},
		{
			ID:          5,
			Name:        "Highest",
			Description: "Highest priority - critical issues that need immediate attention",
			Color:       "#FF5630",
			Level:       5,
			IsActive:    true,
		},
		{
			ID:          6,
			Name:        "Blocker",
			Description: "Blocker - issues that completely block progress",
			Color:       "#DE350B",
			Level:       6,
			IsActive:    true,
		},
	}

	printInfo("Seeding priorities...")
	if err := gormOrm.Trx.Where("id IN ?", []uint{1, 2, 3, 4, 5, 6}).Delete(&models.Priority{}).Error; err != nil {
		log.Fatalf("%s failed to delete existing priorities: %v", red("[x]"), err)
	}
	if err := gormOrm.Trx.Create(priorities).Error; err != nil {
		log.Fatalf("%s failed to create default priorities: %v", red("[x]"), err)
	}
	printSuccess(fmt.Sprintf("Created %d priorities", len(priorities)))
}
