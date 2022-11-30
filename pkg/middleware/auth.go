package middleware

import (
	"log"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// Make sure the user token is valid and admin token is valid
// user token is valid if the user is logged in
// admin token is valid if the admin is logged in

// user token cannot be used to access admin endpoints

// New ...
func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		header := c.GetReqHeaders()

		if _, ok := header["Token"]; !ok {
			return config.Unauthorized(c)
		}

		// if exist, check if the token is valid
		userToken := models.UserToken{}
		err := database.DB.Where("token = ?", header["Token"]).First(&userToken).Error
		if err != nil {
			return config.Unauthorized(c)
		}

		// Separate the buyer token and the seller token
		// If the user is a buyer, the user can only access buyer endpoints
		// If the user is a seller, the user can only access seller endpoints
		if userToken.Type == "buyer" {
			if _, ok := header["Buyer"]; !ok {
				return config.Unauthorized(c)
			}
		} else if userToken.Type == "seller" {
			if _, ok := header["Seller"]; !ok {
				return config.Unauthorized(c)
			}
		}

		c.Locals("user", userToken)
		log.Println("User token is valid")
		return c.Next()
	}
}
