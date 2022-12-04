package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/database"
	"github.com/faruqii/Startup-Campus-Final-Project-Backend/pkg/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func UserToken() string {
	return os.Getenv("USER_TOKEN_SECRET")
}

func SignUp(c *fiber.Ctx) error {
	req := models.UserRegister{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	// Create User type default is buyer
	var buyerType models.Type
	err = database.DB.Where("name = ?", "buyer").First(&buyerType).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: string(pass),
		Type:     "buyer",
	}

	// Check if email already exist
	var checkUser models.User
	err = database.DB.Where("email = ?", req.Email).First(&checkUser).Error
	if err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{
			"message": "Email already Registered",
		})
	}

	err = database.DB.Create(&user).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// create user type
	userType := models.UserType{
		UserID: user.ID,
		TypeID: buyerType.ID,
	}

	err = database.DB.Create(&userType).Error
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created",
	})
}

func SignIn(c *fiber.Ctx) error {
	req := models.UserLogin{}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	user := models.User{}

	err := database.DB.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user.ID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(UserToken()))

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	userToken := models.UserToken{
		UserID: user.ID,
		Type:   user.Type,
		Token:  token,
	}

	err = database.DB.Create(&userToken).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	resp := models.LoginResponse{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone,
		Type:  user.Type,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success login",
		"user":    resp,
		"token":   token,
	})
}

func SignOut(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	err := database.DB.Where("user_id = ?", user.ID).Delete(&models.UserToken{}).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Success logout",
	})
}

// user detail
func UserDetail(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserToken)

	var userDetail models.User

	err := database.DB.Where("id = ?", user.UserID).First(&userDetail).Error

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}

	userResponse := models.UserDetail{
		ID:    userDetail.ID,
		Name:  userDetail.Name,
		Email: userDetail.Email,
		Phone: userDetail.Phone,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": userResponse,
	})
}
