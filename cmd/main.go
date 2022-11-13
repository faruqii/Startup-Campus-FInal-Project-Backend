package main

import (
	"log"
	"os"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &database.Config{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMODE:  os.Getenv("DB_SSLMODE"),
	}

	database.Connect(config)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	routes.Setup(app)
	err = app.Listen(":3000")

	if err != nil {
		panic(err)
	}

}
