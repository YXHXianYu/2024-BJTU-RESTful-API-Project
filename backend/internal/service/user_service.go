package service

import (
	"restfulapi/internal/model"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user model.User) (error) {
	result := db.Create(&user)
	return result.Error
}

func GetUser(db *gorm.DB, id string) (model.User, error) {
	var targetUser model.User
	result := db.First(&targetUser, id)
	return targetUser, result.Error
}

func GetAllUsers(db *gorm.DB) ([]model.User, error) {
	var users []model.User
	result := db.Find(&users)
	return users, result.Error
}