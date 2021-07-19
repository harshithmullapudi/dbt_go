package routes

import (
	"dbt_go/app/controllers"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method:

	route.Get("/token/new", controllers.GetNewAccessToken) // create a new access tokens
}
