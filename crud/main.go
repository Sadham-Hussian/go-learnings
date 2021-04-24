package main

import (
	"github.com/Sadham-Hussian/go-learnings/crud/pkg/configs"
	"github.com/Sadham-Hussian/go-learnings/crud/pkg/routes"
	"github.com/Sadham-Hussian/go-learnings/crud/pkg/utils"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config
	app := fiber.New(config)

	routes.SetUpRoutes(app)

	utils.StartServer(app)
}
