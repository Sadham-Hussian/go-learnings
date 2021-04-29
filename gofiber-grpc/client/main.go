package main

import (
	"log"

	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/client/product"
	"github.com/Sadham-Hussian/go-learnings/gofiber-grpc/client/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	utils.ConnectServer()

	app := fiber.New()

	product.SetUpRoutes(app)

	log.Fatal(app.Listen(":5000"))
}
