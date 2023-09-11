package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model // Embed gorm.Model to include common fields
	AuthorName string `json:"author_name"`
	Rating     int    `json:"rating"`
	Text       string `json:"text"`
	RestaurantID uint // Foreign key
}

type Restaurant struct {
	gorm.Model // Embed gorm.Model to include common fields
	RestaurantName string  `json:"restaurantName"`
	Address       string  `json:"address"`
	Photo         string  `json:"photo"`
	Lat           float64 `json:"lat"`
	Long          float64 `json:"long"`
	Rating        int     `json:"rating"`
	Ratings       []Rating `gorm:"foreignKey:RestaurantID"`
}


