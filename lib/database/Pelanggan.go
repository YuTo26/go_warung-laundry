package database

import (
	"go_warung-laundry/models"

	"gorm.io/gorm"
)

func CreatePelanggan(db *gorm.DB, pelanggan *models.Pelanggan) error {
	result := db.Create(pelanggan)
	return result.Error
}

func GetAllPelanggan(db *gorm.DB) (*[]models.Pelanggan, error) {
	var pelanggan []models.Pelanggan
	if err := db.Find(&pelanggan).Error; err != nil {
		return nil, err
	}
	return &pelanggan, nil
}

func GetPelangganById(db *gorm.DB, id int) (*models.Pelanggan, error) {
	var pelanggan models.Pelanggan
	if err := db.Where("id_pelanggan = ?", id).First(&pelanggan).Error; err != nil {
		return nil, err
	}
	return &pelanggan, nil
}

func UpdatePelanggan(db *gorm.DB, pelanggan *models.Pelanggan) error {
	result := db.Save(pelanggan)
	return result.Error
}

func DeletePelanggan(db *gorm.DB, pelanggan *models.Pelanggan) error {
	result := db.Delete(&pelanggan)
	return result.Error
}
