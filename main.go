package main

import (
	"log"

	"baledev.com/products-api/database"
	_ "baledev.com/products-api/docs"
	"baledev.com/products-api/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	app := fiber.New()
	app.Use(cors.New())

	database.ConnectDB()

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
