package database

import (
	"go_warung-laundry/config"
	"go_warung-laundry/models"
)

func CreateUser(user *models.User) error {
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUsers() (users []models.User, err error) {
	if err = config.DB.Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUser(user *models.User) (err error) {
	if err = config.DB.First(&user).Error; err != nil {
		return
	}
	return
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func LoginUser(user *models.User) error {
	if err := config.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUserToken(user *models.User) error {
	if err := config.DB.Model(&user).Update("token", user.Token).Error; err != nil {
		return err
	}
	return nil
}
