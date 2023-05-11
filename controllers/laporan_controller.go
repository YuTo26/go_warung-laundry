package controllers

import (
	"net/http"
	"time"

	"go_warung-laundry/config"
	"go_warung-laundry/models"

	"github.com/labstack/echo/v4"
)

// Get all laporan
func GetAllLaporan(c echo.Context) error {
	var laporan []models.Laporan

	if err := config.DB.Find(&laporan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all laporan",
		"laporan": laporan,
	})
}

// Get laporan by ID
func GetLaporanByID(c echo.Context) error {
	id := c.Param("id")

	var laporan models.Laporan

	if err := config.DB.First(&laporan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get laporan by ID",
		"laporan": laporan,
	})
}

// Create new laporan
func CreateLaporan(c echo.Context) error {
	laporan := models.Laporan{
		TglLaporan: time.Now(),
	}

	if err := c.Bind(&laporan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&laporan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new laporan",
		"laporan": laporan,
	})
}

// Update laporan by ID
func UpdateLaporanByID(c echo.Context) error {
	id := c.Param("id")

	laporan := models.Laporan{}

	if err := config.DB.First(&laporan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&laporan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&laporan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update laporan by ID",
		"laporan": laporan,
	})
}

// Delete laporan by ID
func DeleteLaporanByID(c echo.Context) error {
	id := c.Param("id")

	laporan := models.Laporan{}

	if err := config.DB.First(&laporan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&laporan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete laporan by ID",
	})
}
