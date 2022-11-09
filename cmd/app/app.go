package main

import (
	"log"
	"os"
	"regexp"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/cmd/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/cmd/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file from root directory
	projectDirName := "Startup-Campus-Final-Project-Backend"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currDir, _ := os.Getwd()
	rootDir := projectName.FindString(currDir)

	err := godotenv.Load(string(rootDir) + "/.env")

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
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err)
	}

}
