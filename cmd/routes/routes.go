package routes

import (
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/cmd/controllers"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/cmd/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/signin", controllers.SignIn)

	seller := api.Group("/seller")
	seller.Post("/signup", controllers.SignUpBuyer)
	seller.Post("/signin", controllers.SignIn)

	// product endpoints can be accessed by seller
	product := seller.Group("/product").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	product.Post("/create", controllers.CreateProduct)
	product.Get("/get", controllers.GetProduct)


}