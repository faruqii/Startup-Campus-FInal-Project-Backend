package routes

import (
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/controllers"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	// ================== User ==================
	// User
	user := api.Group("/user")
	user.Post("/signup", controllers.SignUp)
	user.Post("/signin", controllers.SignIn)
	user.Post("/logout", controllers.SignOut)

	// User Address
	userAddress := user.Group("/shipping_address").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	userAddress.Get("/", controllers.GetShippingAddress)
	userAddress.Post("/", controllers.ChangeShippingAddress)

	// shipping address
	shippingAddress := api.Group("/shipping-address").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	shippingAddress.Post("", controllers.CreateShippingAddress)

	// Balance
	balance := user.Group("/balance").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	balance.Post("", controllers.TopupBalance)
	balance.Get("", controllers.GetBalance)

	// order
	order := api.Group("/order").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	order.Post("", controllers.Order)

	// Cart
	cart := api.Group("/cart").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	cart.Post("", controllers.CreateChart)
	cart.Get("", controllers.GetCart)
	cart.Delete("/:id", controllers.DeleteCart)

	shipping_price := api.Group("/shipping_price").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	shipping_price.Get("", controllers.GetShippingCost)
	// ================== End User ==================

	// =================== ADMIN ===================
	// product endpoints can be accessed by seller
	product := api.Group("/product").Use(middleware.New(middleware.Config{
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
	category := api.Group("/category").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	category.Post("/create", controllers.CreateCategory)
	category.Get("/categories", controllers.GetCategories)

	// order endpoints can be accessed by seller
	orders := api.Group("/orders").Use(middleware.New(middleware.Config{
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(401).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		},
	}))
	orders.Get("", controllers.GetOrders)

}
