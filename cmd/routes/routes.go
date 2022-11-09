package routes

import (
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/cmd/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	user := api.Group("/user")
	user.Post("/register", controllers.SignUp)
	user.Post("/login", controllers.SignIn)

}