package routes

import (
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/cmd/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/signin", controllers.SignIn)

}