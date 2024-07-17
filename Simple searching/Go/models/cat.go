package models

import (
	"gorm.io/gorm"
)

type Cat struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func GetCats(db *gorm.DB) ([]Cat, error) {
	var cats []Cat
	result := db.Find(&cats)
	return cats, result.Error
}
