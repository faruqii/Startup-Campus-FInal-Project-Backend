package controllers

import (
	"net/http"
	"sort"
	"strconv"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
	req := models.ProductRequest{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	product := models.Product{
		Name:         req.Name,
		Sizes:        req.Sizes,
		Details:      req.Details,
		Price:        req.Price,
		ImageURL:     req.ImageURL,
		Condition:    req.Condition,
		CategoryID:   req.CategoryID,
	}

	err := database.DB.Create(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// preloading category
	err = database.DB.Preload("Category").First(&product).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Product created",
		"data":    product,
	})
}

func GetProductDetails(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	err := database.DB.Preload("Category").First(&product, id).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product details",
		"data":    product,
	})
}

func GetProductList(c *fiber.Ctx) error {
	// Use query params to filter products
	// Example: /api/product/products?category=1&size=XL
	page := c.Query("page")
	// if page is not provided, default to 1
	if page == "" {
		page = "1"
	}

	// convert page to int
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page",
		})
	}
	pageSize := c.Query("page_size")
	// if page_size is not provided, default to 10
	if pageSize == "" {
		pageSize = "10"
	}

	// convert page_size to int
	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page_size",
		})
	}

	sortBy := c.Query("sort_by")
	category := c.Query("category")
	price := c.Query("price")
	condition := c.Query("condition")
	productName := c.Query("product_name")

	var products []models.Product

	err = database.DB.Preload("Category").Find(&products).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// Filter products
	if category != "" {
		err := database.DB.Preload("Category").Where("category_id = ?", category).Find(&products).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	if price != "" {
		err := database.DB.Preload("Category").Where("price = ?", price).Find(&products).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	if condition != "" {
		err := database.DB.Preload("Category").Where("condition = ?", condition).Find(&products).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	if productName != "" {
		err := database.DB.Preload("Category").Where("name LIKE ?", "%"+productName+"%").Find(&products).Error
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	// Sort products
	if sortBy != "" {
		switch sortBy {
		case "price":
			sort.Slice(products, func(i, j int) bool {
				return products[i].Price < products[j].Price
			})
		case "name":
			sort.Slice(products, func(i, j int) bool {
				return products[i].Name < products[j].Name
			})
		}
	}

	// Pagination
	total := len(products)
	start := (pageInt - 1) * pageSizeInt

	if start > total {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid page",
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Product list",
		"data":    products,
	})

}
