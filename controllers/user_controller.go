package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"go_warung-laundry/config"
	"go_warung-laundry/models"
)

// Create new user
func CreateUser(c echo.Context) error {
	user := models.User{}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Hash user password before saving to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// Get all users
func GetAllUsers(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// Get user by ID
func GetUserByID(c echo.Context) error {
	id := c.Param("id")

	var user models.User

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user by ID",
		"user":    user,
	})
}

// Update user by ID
func UpdateUserByID(c echo.Context) error {
	id := c.Param("id")

	user := models.User{}

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// Hash user password before saving to database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user by ID",
		"user":    user,
	})
}

// Delete user by ID
func DeleteUserByID(c echo.Context) error {
	id := c.Param("id")

	user := models.User{}

	if err := config.DB.First(&user, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user by ID",
	})
}
