package controllers

import (
	"net/http"

	"go_warung-laundry/config"
	"go_warung-laundry/models"

	"github.com/labstack/echo/v4"
)

// Get all jenis laundry
func GetAllJenisLaundry(c echo.Context) error {
	var jenisLaundry []models.JenisLaundry

	if err := config.DB.Find(&jenisLaundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get all jenis laundry",
		"jenisLaundry": jenisLaundry,
	})
}

// Get jenis laundry by ID
func GetJenisLaundryByID(c echo.Context) error {
	id := c.Param("id")

	var jenisLaundry models.JenisLaundry

	if err := config.DB.First(&jenisLaundry, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success get jenis laundry by ID",
		"jenisLaundry": jenisLaundry,
	})
}

// Create new jenis laundry
func CreateJenisLaundry(c echo.Context) error {
	jenisLaundry := models.JenisLaundry{}

	if err := c.Bind(&jenisLaundry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&jenisLaundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":      "success create new jenis laundry",
		"jenisLaundry": jenisLaundry,
	})
}

// Update jenis laundry by ID
func UpdateJenisLaundryByID(c echo.Context) error {
	id := c.Param("id")

	jenisLaundry := models.JenisLaundry{}

	if err := config.DB.First(&jenisLaundry, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&jenisLaundry); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&jenisLaundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "success update jenis laundry by ID",
		"jenisLaundry": jenisLaundry,
	})
}

// Delete jenis laundry by ID
func DeleteJenisLaundryByID(c echo.Context) error {
	id := c.Param("id")

	jenisLaundry := models.JenisLaundry{}

	if err := config.DB.First(&jenisLaundry, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&jenisLaundry).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete jenis laundry by ID",
	})
}
