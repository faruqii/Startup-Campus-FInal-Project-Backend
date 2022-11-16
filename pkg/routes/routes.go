package routes

import (
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/controllers"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/signin", controllers.SignIn)
	user.Post("/logout", controllers.SignOut)

	// Balance
	balance := user.Group("/balance").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	balance.Post("/add", controllers.TopupBalance)
	balance.Get("/get", controllers.GetBalance)

	// Seller
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
	product.Get("/products/:id", controllers.GetProductDetails)
	product.Get("/products", controllers.GetProductList)
	product.Put("/products/:id", controllers.UpdateProduct)
	product.Delete("/products/:id", controllers.DeleteProduct)

	// category endpoints can be accessed by seller
	category := seller.Group("/category").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	category.Post("/create", controllers.CreateCategory)
	category.Get("/categories", controllers.GetCategories)

}
