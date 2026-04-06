package main

import (
	"asset-management/config"
	"asset-management/migrations"
	"asset-management/routes"
	"asset-management/seeders"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if err := config.ConnectDB(); err != nil {
		fmt.Println("DB Error:", err)
		return
	}

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "migrate":
			if err := migrations.Migrate(config.DB); err != nil {
				fmt.Println("Migration error:", err)
			}
		case "seed":
			if err := seeders.Seed(config.DB); err != nil {
				fmt.Println("Seeder error:", err)
			}
		case "fresh":
			_ = migrations.DropAll(config.DB)
			_ = migrations.Migrate(config.DB)
			_ = seeders.Seed(config.DB)
		case "serve":
			startServer()
		default:
			fmt.Println("Command not found")
		}
		return
	}

	startServer()
}

func startServer() {
	app := fiber.New()
	app.Use(cors.New())

	routes.SetupRoutes(app, config.DB)
	app.Static("/qrcodes", "./qrcodes")

	if err := app.Listen(":8080"); err != nil {
		fmt.Println("Server Error:", err)
	}
}
