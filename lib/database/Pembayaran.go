package database

import (
	"go_warung-laundry/models"

	"gorm.io/gorm"
)

func CreatePembayaran(db *gorm.DB, pembayaran *models.Pembayaran) error {
	result := db.Create(pembayaran)
	return result.Error
}

func GetAllPembayaran(db *gorm.DB) (*[]models.Pembayaran, error) {
	var pembayaran []models.Pembayaran
	if err := db.Find(&pembayaran).Error; err != nil {
		return nil, err
	}
	return &pembayaran, nil
}

func GetPembayaranById(db *gorm.DB, id int) (*models.Pembayaran, error) {
	var pembayaran models.Pembayaran
	if err := db.Where("id_pembayaran = ?", id).First(&pembayaran).Error; err != nil {
		return nil, err
	}
	return &pembayaran, nil
}

func UpdatePembayaran(db *gorm.DB, pembayaran *models.Pembayaran) error {
	result := db.Save(pembayaran)
	return result.Error
}

func DeletePembayaran(db *gorm.DB, pembayaran *models.Pembayaran) error {
	result := db.Delete(pembayaran)
	return result.Error
}
