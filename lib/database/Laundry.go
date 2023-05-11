package database

import (
	"go_warung-laundry/models"

	"gorm.io/gorm"
)

func GetAllLaundry(db *gorm.DB) (*[]models.Laundry, error) {
	var laundries []models.Laundry
	if err := db.Find(&laundries).Error; err != nil {
		return nil, err
	}
	return &laundries, nil
}

func GetLaundryById(db *gorm.DB, id int) (*models.Laundry, error) {
	var laundry models.Laundry
	if err := db.Where("id_laundry = ?", id).First(&laundry).Error; err != nil {
		return nil, err
	}
	return &laundry, nil
}

func CreateLaundry(db *gorm.DB, laundry *models.Laundry) error {
	if err := db.Create(&laundry).Error; err != nil {
		return err
	}
	return nil
}

func UpdateLaundry(db *gorm.DB, laundry *models.Laundry) error {
	if err := db.Save(&laundry).Error; err != nil {
		return err
	}
	return nil
}

func DeleteLaundry(db *gorm.DB, laundry *models.Laundry) error {
	if err := db.Delete(&laundry).Error; err != nil {
		return err
	}
	return nil
}
