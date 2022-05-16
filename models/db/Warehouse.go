package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null"`
	Description string  `json:"description" gorm:"column:description;not null;default:''"`
	Longitude   float32 `json:"lon" gorm:"column:lon;not null;default:0"`
	Latitude    float32 `json:"lat" gorm:"column:lat;not null;default:0"`
	Items       []Item  `json:"items"`
}
