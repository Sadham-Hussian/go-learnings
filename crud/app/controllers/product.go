package controllers

import (
	"log"

	"github.com/Sadham-Hussian/go-learnings/crud/app/models"
	"github.com/Sadham-Hussian/go-learnings/crud/app/services"
	"github.com/gofiber/fiber/v2"
)

// Home controller for api /api/v1/
func Home(c *fiber.Ctx) error {
	return c.SendString("Welcome home")
}

// GetProductByID controller for api /api/v1/product/:id
func GetProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	response, status := services.GetProductByID(id)

	c.JSON(response)
	return c.SendStatus(status)
}

// GetAllProducts controller for api /api/v1/product/all
func GetAllProducts(c *fiber.Ctx) error {
	response, status := services.GetAllProducts()

	c.JSON(response)
	return c.SendStatus(status)
}

// CreateProduct controller for api /api/v1/product/new
func CreateProduct(c *fiber.Ctx) error {
	product := new(models.Product)

	if err := c.BodyParser(product); err != nil {
		log.Println(err)
		c.JSON(fiber.Map{
			"error":  "error parsing request body",
			"status": "fail",
		})
		return c.SendStatus(500)
	}

	response, status := services.CreateProduct(product)
	c.JSON(response)
	return c.SendStatus(status)
}

// DeleteProductByID controller for api /api/v1/product/delete/:id
func DeleteProductByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	response, status := services.DeleteProductByID(id)

	c.JSON(response)
	return c.SendStatus(status)
}
