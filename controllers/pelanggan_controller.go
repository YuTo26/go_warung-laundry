package controllers

import (
	"go_warung-laundry/config"
	"go_warung-laundry/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get all pelanggan
func GetAllPelanggan(c echo.Context) error {
	var pelanggan []models.Pelanggan

	if err := config.DB.Find(&pelanggan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"pelanggan": pelanggan,
		"message":   "success get all pelanggan",
	})
}

// Get pelanggan by ID
func GetPelangganByID(c echo.Context) error {
	id := c.Param("id")

	var pelanggan models.Pelanggan

	if err := config.DB.First(&pelanggan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"pelanggan": pelanggan,
		"message":   "success get pelanggan by ID",
	})
}

// Create new pelanggan
func CreatePelanggan(c echo.Context) error {
	pelanggan := models.Pelanggan{}

	if err := c.Bind(&pelanggan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&pelanggan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"pelanggan": pelanggan,
		"message":   "success create new pelanggan",
	})
}

// Update pelanggan by ID
func UpdatePelangganByID(c echo.Context) error {
	id := c.Param("id")

	pelanggan := models.Pelanggan{}

	if err := config.DB.First(&pelanggan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&pelanggan); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&pelanggan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"pelanggan": pelanggan,
		"message":   "success update pelanggan by ID",
	})
}

// Delete pelanggan by ID
func DeletePelangganByID(c echo.Context) error {
	id := c.Param("id")

	pelanggan := models.Pelanggan{}

	if err := config.DB.First(&pelanggan, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&pelanggan).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pelanggan by ID",
	})
}

//controller pelanggan untuk menghandle pembayaran

// Get all pembayaran
func PelangganGetAllPembayaran(c echo.Context) error {
	var pembayaran []models.Pembayaran

	if err := config.DB.Find(&pembayaran).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get all pembayaran",
		"pembayaran": pembayaran,
	})
}

// Get pembayaran by ID
func PelangganGetPembayaranByID(c echo.Context) error {
	id := c.Param("id")

	var pembayaran models.Pembayaran

	if err := config.DB.First(&pembayaran, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success get pembayaran by ID",
		"pembayaran": pembayaran,
	})
}

// Create new pembayaran
func PelangganCreatePembayaran(c echo.Context) error {
	pembayaran := models.Pembayaran{}

	if err := c.Bind(&pembayaran); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&pembayaran).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":    "success create new pembayaran",
		"pembayaran": pembayaran,
	})
}

// Update pembayaran by ID
func PelangganUpdatePembayaranByID(c echo.Context) error {
	id := c.Param("id")

	pembayaran := models.Pembayaran{}

	if err := config.DB.First(&pembayaran, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&pembayaran); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&pembayaran).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success update pembayaran by ID",
		"pembayaran": pembayaran,
	})
}

// Delete pembayaran by ID
func PelangganDeletePembayaranByID(c echo.Context) error {
	id := c.Param("id")

	pembayaran := models.Pembayaran{}

	if err := config.DB.First(&pembayaran, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&pembayaran).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pembayaran by ID",
	})
}

// Method untuk mengelola transaksi ke dalam database
func PelangganGetAllTransaksi(c echo.Context) error {
	transaksi := []models.Transaction{}

	if err := config.DB.Find(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get all transaksi",
		"transaksi": transaksi,
	})
}

func PelangganGetTransaksiByID(c echo.Context) error {
	var transaksi models.Transaction
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success get transaksi by ID",
		"transaksi": transaksi,
	})
}

func PelangganCreateTransaksi(c echo.Context) error {
	transaksi := models.Transaction{}
	if err := c.Bind(&transaksi); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Create(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":   "success create new transaksi",
		"transaksi": transaksi,
	})
}

func PelangganUpdateTransaksiByID(c echo.Context) error {
	var transaksi models.Transaction
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := c.Bind(&transaksi); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := config.DB.Save(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":   "success update transaksi by ID",
		"transaksi": transaksi,
	})
}

func PelangganDeleteTransaksiByID(c echo.Context) error {
	var transaksi models.Transaction
	id := c.Param("id")

	if err := config.DB.Where("id = ?", id).First(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	if err := config.DB.Delete(&transaksi).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete transaksi by ID",
	})
}
