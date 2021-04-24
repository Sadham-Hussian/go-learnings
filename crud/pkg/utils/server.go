package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

// StartServer function to start the Fiber server
func StartServer(app *fiber.App) {
	// Run server.

	if err := app.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running! Reason: %v", err)
	} else {
		fmt.Println("Server running")
	}
}
