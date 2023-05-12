package database

import (
	"go_warung-laundry/models"

	"gorm.io/gorm"
)

func CreateLaporan(db *gorm.DB, laporan *models.Laporan) error {
	result := db.Create(laporan)
	return result.Error
}

func GetAllLaporan(db *gorm.DB) (*[]models.Laporan, error) {
	var laporan []models.Laporan
	if err := db.Find(&laporan).Error; err != nil {
		return nil, err
	}
	return &laporan, nil
}

func GetLaporanById(db *gorm.DB, id int) (*models.Laporan, error) {
	var laporan models.Laporan
	if err := db.Where("id_laporan = ?", id).First(&laporan).Error; err != nil {
		return nil, err
	}
	return &laporan, nil
}

func UpdateLaporan(db *gorm.DB, laporan *models.Laporan) error {
	result := db.Save(laporan)
	return result.Error
}

func DeleteLaporan(db *gorm.DB, laporan *models.Laporan) error {
	result := db.Delete(laporan)
	return result.Error
}
