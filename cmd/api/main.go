package main

import (
	"log"

	"github.com/Mr-Chegini/simple-golang-crud/database"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message"`
	Version string `json:"version,omitempty"`
}

func main() {
	// Initialize database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	// Run database migrations
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	app := fiber.New()

	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		log.Printf("%s %s", c.Method(), c.Path())
		return c.Next()
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	})

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Status:  "ok",
			Message: "Simple Go CRUD API is running",
		})
	})

	// API root
	app.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Message: "Welcome to Simple Go CRUD API",
			Version: "v1",
		})
	})

	app.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.JSON(Response{
			Message: "Welcome to Simple Go CRUD API",
			Version: "v1",
		})
	})

	log.Println("🚀 Server starting on port 8080...")
	log.Fatal(app.Listen(":8080"))
}
