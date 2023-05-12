package database

import (
	"go_warung-laundry/models"

	"gorm.io/gorm"
)

func GetAllTransaction(db *gorm.DB) (*[]models.Transaction, error) {
	var transactions []models.Transaction
	if err := db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return &transactions, nil
}

func GetTransactionById(db *gorm.DB, id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := db.Where("id_transaksi = ?", id).First(&transaction).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}

func CreateTransaction(db *gorm.DB, transaction *models.Transaction) error {
	if err := db.Create(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func UpdateTransaction(db *gorm.DB, transaction *models.Transaction) error {
	if err := db.Save(&transaction).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTransaction(db *gorm.DB, transaction *models.Transaction) error {
	if err := db.Delete(&transaction).Error; err != nil {
		return err
	}
	return nil
}
