package database

import "user-personalize/models"

func Migrate() {
	DB.AutoMigrate(&models.Users{})
	DB.AutoMigrate(&models.Foto{})
	DB.AutoMigrate(&models.Login{})
}