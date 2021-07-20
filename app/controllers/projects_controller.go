package controllers

import (
	"dbt_go/app/models"
	"dbt_go/app/serializers"
	"dbt_go/platform/database"

	"github.com/gofiber/fiber/v2"
)

// GetProjects func gets all exists projects.
// @Description Get all exists projects.
// @Summary get all exists projects
// @Tags Projects
// @Accept json
// @Produce json
// @Success 200 {array} models.Project
// @Router /v1/projects [get]
func GetProjects(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	// Get all projects.
	projects, err := db.GetProjects()
	if err != nil {
		// Return, if projects not found.
		return c.Status(fiber.StatusNotFound).JSON(models.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"count":    len(projects),
		"projects": serializers.ProjectSerializerMap(projects),
	})
}

// CreateProjects func creates a new  project.
// @Description Create a new projects.
// @Summary create a new project
// @Tags Projects
// @Accept json
// @Produce json
// @Success 200 {object} models.Project
// @Router /v1/projects/create [post]
func CreateProject(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	// Create new Project struct
	project := &models.Project{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(project); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(models.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	if err := db.CreateProject(project); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.APIResponse{
			Error:   true,
			Message: err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(serializers.ProjectSerializer(project))
}
