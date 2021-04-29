package product

import (
	"context"
	"io"
	"log"

	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/client/utils"
	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/proto"
	"github.com/gofiber/fiber/v2"
)

// GetAllProductsHandler handler to get all products
func GetAllProductsHandler(c *fiber.Ctx) error {
	stream, err := utils.Client.GetProducts(context.Background(), &proto.GetProductsRequest{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}
	var productItems []ProductItem
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "no products available"})
		}

		productItem := ProductItem{
			ID:    result.Product.Id,
			Name:  result.Product.Name,
			Type:  result.Product.Type,
			Prize: result.Product.Prize,
		}

		productItems = append(productItems, productItem)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"products": productItems})
}

// GetProductHandler handler to get a product by id
func GetProductHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	result, err := utils.Client.GetProduct(context.Background(), &proto.GetProductRequest{Id: int32(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	productItem := &ProductItem{
		ID:    result.Product.Id,
		Name:  result.Product.Name,
		Type:  result.Product.Type,
		Prize: result.Product.Prize,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": productItem})
}

// CreateProductHandler handler to create a product
func CreateProductHandler(c *fiber.Ctx) error {
	productItem := new(ProductItem)

	if err := c.BodyParser(productItem); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	product := &proto.Product{
		Name:  productItem.Name,
		Type:  productItem.Type,
		Prize: productItem.Prize,
	}

	result, err := utils.Client.CreateProduct(context.Background(), &proto.CreateProductRequest{Product: product})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid argument"})
	}

	productItem.ID = result.Product.Id

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": productItem})
}

// UpdateProductHandler handler to update a product
func UpdateProductHandler(c *fiber.Ctx) error {
	productItem := new(ProductItem)

	if err := c.BodyParser(productItem); err != nil {
		log.Println(err)
		return c.SendStatus(500)
	}

	product := &proto.Product{
		Id:    productItem.ID,
		Name:  productItem.Name,
		Type:  productItem.Type,
		Prize: productItem.Prize,
	}

	result, err := utils.Client.UpdateProduct(context.Background(), &proto.UpdateProductRequest{Product: product})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Invalid product id"})
	}

	productItem.Name = result.Product.Name
	productItem.Type = result.Product.Type
	productItem.Prize = result.Product.Prize

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"product": productItem})
}

// DeleteProductHandler handler to delete a product by id
func DeleteProductHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(400)
	}

	result, err := utils.Client.DeleteProduct(context.Background(), &proto.DeleteProductRequest{Id: int32(id)})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
	}

	if result.Success == true {
		return c.SendStatus(fiber.StatusOK)
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid id"})
}
