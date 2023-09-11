package initializers

import "github.com/jenniekibiri/go-backend/models"

func SyncDb() {
	DB.AutoMigrate(&models.Restaurant{}, &models.Rating{})
}
