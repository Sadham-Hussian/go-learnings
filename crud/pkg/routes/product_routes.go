package routes

import (
	"github.com/Sadham-Hussian/go-learnings/crud/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetProductRoutes function to set product related routes
func SetProductRoutes(route fiber.Router) {
	product := route.Group("/products")

	product.Get("/", controllers.GetAllProducts)
	product.Get("/:id", controllers.GetProductByID)
	product.Post("/", controllers.CreateProduct)
	product.Delete("/:id", controllers.DeleteProductByID)
}
