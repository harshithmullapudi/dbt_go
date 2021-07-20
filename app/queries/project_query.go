package queries

import (
	"dbt_go/app/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// ProjectQueries struct for queries from Project model.
type ProjectQueries struct {
	*gorm.DB
}

// GetProjects method for getting all projects.
func (q *ProjectQueries) GetProjects() ([]models.Project, error) {
	// Define projects variable.
	projects := []models.Project{}

	// Send query to database.
	result := q.Find(&projects)

	if result.Error != nil {
		return projects, result.Error
	}

	return projects, nil
}

func (q *ProjectQueries) CreateProject(project *models.Project) error {

	project.UUID = uuid.New()

	// Send query to database.
	result := q.Create(&project)
	if result.Error != nil {
		// Return only error.
		return result.Error
	}

	// This query returns nothing.
	return nil
}
