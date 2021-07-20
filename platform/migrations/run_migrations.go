package migrations

import (
	"dbt_go/app/models"
	"dbt_go/platform/database"
)

func Migrate() {
	database.DBCon.AutoMigrate(models.Project{})
}
