package database

import (
	"Kaho_BaaS/internal/apps/account/models"

	"gorm.io/gorm"
)

var modelList = []interface{}{
	&models.User{},
}

func autoMigrateModels(DB *gorm.DB) {
	for _, model := range modelList {
		DB.AutoMigrate(model)
	}
}
