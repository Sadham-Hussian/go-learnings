package routes

import (
	"github.com/Sadham-Hussian/go-learnings/crud/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetUpRoutes function to add routes
func SetUpRoutes(app *fiber.App) {
	route := app.Group("/api/v1")

	route.Get("/", controllers.Home)

	SetProductRoutes(route)
}
