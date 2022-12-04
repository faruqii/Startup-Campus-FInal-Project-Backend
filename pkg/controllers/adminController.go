package controllers

import (
	"net/http"
	"strconv"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Get User Order
/*
	Admin Get User Order
	Can sort by order price in query params
	query params:
		- sort_by : price a_z, price z_a
		- page : page number
		- page_size : page size

	Example:
		- /api/v1/admin/orders?sort_by=price_a_z&page=1&page_size=10

	Response:
		{
		"data": [
			{
				"id": "order_id(uuid)",
				"user_name": "nama user",
				"created_at": "Tue, 25 august 2022",
				"user_id": "uuid",
				"user_email": "user@gmail.com",
				"total": 1000
			}
    ]
}
*/
func GetOrders(c *fiber.Ctx) error {
	var orders []models.Order
	var total int64

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))

	sortBy := c.Query("sort_by", "price_a_z")

	query := database.DB.Model(&models.Order{}).Preload("OrderItems").Preload("OrderItems.Product").Preload("User").Preload("User.UserBalance").Preload("User.UserAddress").Preload("User.UserAddress.Province").Preload("User.UserAddress.City").Preload("User.UserAddress.Subdistrict").Preload("User.UserAddress.Village").Preload("User.UserAddress.PostalCode").Preload("User.UserAddress.PostalCode.Province").Preload("User.UserAddress.PostalCode.City").Preload("User.UserAddress.PostalCode.Subdistrict").Preload("User.UserAddress.PostalCode.Village").Preload("User.UserAddress.PostalCode.PostalCode")

	if sortBy == "price_a_z" {
		query = query.Order("total_price asc")
	} else if sortBy == "price_z_a" {
		query = query.Order("total_price desc")
	}

	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&orders)
	database.DB.Model(&models.Order{}).Count(&total)

	for _, order := range orders {
		ordersResponse := models.OrderResponse{
			ID:        order.ID,
			UserName:  order.User.Name,
			CreatedAt: order.CreatedAt,
			UserID:    order.User.ID,
			UserEmail: order.User.Email,
			Total:     order.TotalPrice,
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Get Orders Success",
			"data":    ordersResponse,
			"total":   total,
		})
	}

	ordersResponse := models.OrderResponse{
		ID:        orders[0].ID,
		UserName:  orders[0].User.Name,
		CreatedAt: orders[0].CreatedAt,
		UserID:    orders[0].User.ID,
		UserEmail: orders[0].User.Email,
		Total:     orders[0].TotalPrice,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Get Orders Success",
		"data":    ordersResponse,
		"total":   total,
	})
}
