package serializers

import "dbt_go/app/models"

// Serialize Project
func ProjectSerializer(project *models.Project) models.ProjectSerializer {
	return models.ProjectSerializer{
		Title: project.Title,
		UUID:  project.UUID,
	}
}

// Serialize array of projects
func ProjectSerializerMap(projects []models.Project) []models.ProjectSerializer {
	var serialisedProjects []models.ProjectSerializer

	for _, project := range projects {
		serialisedProjects = append(serialisedProjects, ProjectSerializer(&project))
	}

	return serialisedProjects
}
