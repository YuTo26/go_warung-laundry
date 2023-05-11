package controllers

import (
	"go_warung-laundry/models"
	"net/http"

	"go_warung-laundry/config"

	"github.com/labstack/echo/v4"
)

// Get all laundry data
func GetAllLaundryController(c echo.Context) error {
	var laundries []models.Laundry

	if err := config.DB.Find(&laundries).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get all laundry data",
		"laundries": laundries,
	})
}

// Get laundry data by ID
func GetLaundryByIDController(c echo.Context) error {
	id := c.Param("id")

	var laundry models.Laundry
	if err := config.DB.First(&laundry, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get laundry data by ID",
		"laundry": laundry,
	})
}

// Add new laundry data
func AddLaundryController(c echo.Context) error {
	var laundry models.Laundry
	if err := c.Bind(&laundry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&laundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add new laundry data",
		"laundry": laundry,
	})
}

// Update laundry data by ID
func UpdateLaundryController(c echo.Context) error {
	id := c.Param("id")

	var laundry models.Laundry
	if err := config.DB.First(&laundry, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&laundry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&laundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update laundry data",
		"laundry": laundry,
	})
}

// Delete laundry data by ID
func DeleteLaundryController(c echo.Context) error {
	id := c.Param("id")

	var laundry models.Laundry
	if err := config.DB.First(&laundry, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Delete(&laundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete laundry data",
		"laundry": laundry,
	})
}
