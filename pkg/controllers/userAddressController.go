package controllers

import (
	"net/http"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Create Shipping Address
func CreateShippingAddress(c *fiber.Ctx) error {
	req := models.ShippingAddressRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := c.Locals("user").(models.UserToken)

	shippingAddress := models.ShippingAddress{}

	// check if user already have shipping address
	err := database.DB.Where("user_id = ?", user.ID).First(&shippingAddress).Error
	if err == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "User already have shipping address",
		})
	}

	// create shipping address
	shippingAddress = models.ShippingAddress{
		UserID:      user.UserID,
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Address:     req.Address,
		City:        req.City,
	}

	err = database.DB.Create(&shippingAddress).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Success create shipping address",
	})

}

// Get Shipping Address
func GetShippingAddress(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var shippingAddress models.ShippingAddress

	err := database.DB.Where("user_id = ?", user.UserID).First(&shippingAddress).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Shipping address not found",
		})
	}

	shippingAddressResponse := models.ShippingAddressResponse{
		ID:          shippingAddress.ID,
		Name:        shippingAddress.Name,
		PhoneNumber: shippingAddress.PhoneNumber,
		Address:     shippingAddress.Address,
		City:        shippingAddress.City,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success get shipping address",
		"data":    shippingAddressResponse,
	})
}

// Change Shipping Address
func ChangeShippingAddress(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	req := models.ShippingAddressRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	var shippingAddress models.ShippingAddress

	err := database.DB.Where("user_id = ?", user.UserID).First(&shippingAddress).Error

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Shipping address not found",
		})
	}

	shippingAddress.Name = req.Name
	shippingAddress.PhoneNumber = req.PhoneNumber
	shippingAddress.Address = req.Address
	shippingAddress.City = req.City

	err = database.DB.Save(&shippingAddress).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success change shipping address",
	})

}
