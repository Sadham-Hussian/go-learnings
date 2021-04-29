package product

import (
	"github.com/gofiber/fiber/v2"
)

// SetUpRoutes function to set up routes
func SetUpRoutes(app *fiber.App) {
	route := app.Group("/api/v1/products")

	route.Get("/", GetAllProductsHandler)
	route.Get("/:id", GetProductHandler)
	route.Post("/", CreateProductHandler)
	route.Put("/:id", UpdateProductHandler)
	route.Delete("/:id", DeleteProductHandler)
}
