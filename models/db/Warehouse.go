package models

import (
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	Name        string  `json:"name" gorm:"not null;unique"`
	Description string  `json:"description" gorm:"column:description;not null;default:''"`
	Longitude   float32 `json:"longitude" gorm:"column:lon;not null;default:0"`
	Latitude    float32 `json:"latitude" gorm:"column:lat;not null;default:0"`
	Items       []Item  `json:"items,omitempty"`
}

func (w Warehouse) Validate() bool {
	return w.Name != "" && w.Longitude != 0 && w.Latitude != 0
}
