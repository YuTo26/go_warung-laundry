package database

import (
	"go_warung-laundry/models"

	"gorm.io/gorm"
)

func GetAllJenisLaundry(db *gorm.DB) (*[]models.JenisLaundry, error) {
	var jenisLaundry []models.JenisLaundry
	if err := db.Find(&jenisLaundry).Error; err != nil {
		return nil, err
	}
	return &jenisLaundry, nil
}

func GetJenisLaundryById(db *gorm.DB, id int) (*models.JenisLaundry, error) {
	var jenisLaundry models.JenisLaundry
	if err := db.Where("id_jenis = ?", id).First(&jenisLaundry).Error; err != nil {
		return nil, err
	}
	return &jenisLaundry, nil
}

func CreateJenisLaundry(db *gorm.DB, jenisLaundry *models.JenisLaundry) error {
	if err := db.Create(&jenisLaundry).Error; err != nil {
		return err
	}
	return nil
}

func UpdateJenisLaundry(db *gorm.DB, jenisLaundry *models.JenisLaundry) error {
	if err := db.Save(&jenisLaundry).Error; err != nil {
		return err
	}
	return nil
}

func DeleteJenisLaundry(db *gorm.DB, jenisLaundry *models.JenisLaundry) error {
	if err := db.Delete(&jenisLaundry).Error; err != nil {
		return err
	}
	return nil
}
