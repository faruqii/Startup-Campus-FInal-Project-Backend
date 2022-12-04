package controllers

import (
	"net/http"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Create Chart
func CreateChart(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	req := models.UserCartRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	// check if product exist
	var product models.Product

	err := database.DB.Where("id = ?", req.ProductID).First(&product).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	// add to cart
	cart := models.UserCart{
		UserID:    user.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Sizes:     req.Sizes,
	}

	err = database.DB.Create(&cart).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Product added to cart",
	})

}

// Get Cart
func GetCart(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var carts []models.UserCart

	err := database.DB.Where("user_id = ?", user.UserID).Find(&carts).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	// preloading product and category in product models
	err = database.DB.Preload("Product").Preload("Product.Category").First(&carts).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"cart": carts,
	})
}

// Delete Cart
func DeleteCart(c *fiber.Ctx) error {
	id := c.Params("id")

	var cart models.UserCart

	err := database.DB.Where("id = ?", id).First(&cart).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	err = database.DB.Delete(&cart).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Cart deleted",
	})
}

// get shipping cost
/*
Only accessible by logged in users
Can only be accessed by users who already have a cart
Calculated based on the current user cart
There will be 2 types of shipping methods:
Regular:
If the total item price < 200: Shipping price is 15% of the total item price purchased
If the total item price >= 200: Shipping price is 20% of the total item price purchased
Next Day:
If the total item price < 300: Shipping price is 20% of the total item price purchased
If the total item price >= 300: Shipping price is 25% of the total item price purchased
*/
func GetShippingCost(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var carts []models.UserCart

	err := database.DB.Where("user_id = ?", user.UserID).Find(&carts).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	// preloading product and category in product models
	err = database.DB.Preload("Product").Preload("Product.Category").First(&carts).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	var product models.Product

	// get product price
	err = database.DB.Where("id = ?", carts[0].ProductID).First(&product).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	// product price convert to int
	productPrice := int(product.Price)

	// get total price
	var totalPrice int
	for _, cart := range carts {
		totalPrice += cart.Quantity * productPrice
	}

	// get shipping cost
	var shippingCostReguler int64
	if totalPrice < 200 {
		shippingCostReguler = int64(float64(totalPrice) * 0.15)
	} else {
		shippingCostReguler = int64(float64(totalPrice) * 0.2)
	}

	var shippingCostNextDay int64
	if totalPrice < 300 {
		shippingCostNextDay = int64(float64(totalPrice) * 0.2)
	} else {
		shippingCostNextDay = int64(float64(totalPrice) * 0.25)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"shipping_cost": []models.ShippingCost{
			{
				Name:  "regular",
				Price: shippingCostReguler,
			},
			{
				Name:  "next day",
				Price: shippingCostNextDay,
			},
		},
	})
}

// Create Order
/*
Only accessible by logged in users
Will reduce user balance based on total price
If the balance is less than the total price, return error
Will delete all carts in that user
*/
func Order(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var carts []models.UserCart

	err := database.DB.Where("user_id = ?", user.UserID).Find(&carts).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Cart not found",
		})
	}

	// preloading product
	err = database.DB.Preload("Product").First(&carts).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Cart is empty",
		})
	}

	var product models.Product

	// get product price
	err = database.DB.Where("id = ?", carts[0].ProductID).First(&product).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	// product price convert to int
	productPrice := int(product.Price)

	// get total price
	var totalPrice int
	for _, cart := range carts {
		totalPrice += cart.Quantity * productPrice
	}

	// get user balance
	var userBalance models.UserBalance

	err = database.DB.Where("user_id = ?", user.UserID).First(&userBalance).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// check user balance
	if userBalance.Balance < float64(totalPrice) {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Your balance is not enough",
		})
	}

	req := models.OrderRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	// get shipping cost
	var shippingCostReguler int64
	if totalPrice < 200 {
		shippingCostReguler = int64(float64(totalPrice) * 0.15)
	} else {
		shippingCostReguler = int64(float64(totalPrice) * 0.2)
	}

	var shippingCostNextDay int64
	if totalPrice < 300 {
		shippingCostNextDay = int64(float64(totalPrice) * 0.2)
	} else {
		shippingCostNextDay = int64(float64(totalPrice) * 0.25)
	}

	if req.ShippingMethod == "regular" {
		req.ShippingPrice = shippingCostReguler
	} else if req.ShippingMethod == "next day" {
		req.ShippingPrice = shippingCostNextDay
	} else {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Shipping method not found",
		})
	}

	// find user shipping address
	var userShippingAddress models.ShippingAddress

	err = database.DB.Where("user_id = ?", user.UserID).First(&userShippingAddress).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Shipping address not found",
		})
	}

	// create order
	order := models.Order{
		UserID:            user.UserID,
		ShippingMethod:    req.ShippingMethod,
		ShippingAddressID: userShippingAddress.ID,
		ShippingPrice:     req.ShippingPrice,
		TotalPrice:        float64(totalPrice),
	}

	err = database.DB.Create(&order).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// update user balance
	err = database.DB.Model(&userBalance).Update("balance", userBalance.Balance-float64(totalPrice)).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// delete cart that belongs to user
	err = database.DB.Where("user_id = ?", user.UserID).Delete(&carts).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Order success",
	})
}
