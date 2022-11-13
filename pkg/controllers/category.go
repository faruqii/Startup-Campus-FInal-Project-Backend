package controllers

import (
	"net/http"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(c *fiber.Ctx) error {
	req := models.CategoryRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	category := models.Category{
		Name: req.Name,
	}

	err := database.DB.Create(&category).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Category created",
	})
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category

	err := database.DB.Find(&categories).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Categories fetched",
		"data":    categories,
	})
}

