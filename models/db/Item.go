package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name               string  `json:"name" gorm:"not null;uniqueIndex:idx_name_city"`
	Description        string  `json:"description" gorm:"column:description;not null;default:''"`
	Price              float32 `json:"price" gorm:"column:price;not null;default:0"`
	Quantity           int     `json:"quantity" gorm:"column:quantity;not null;default:0"`
	City               string  `json:"city" gorm:"column:city;not null;default:'';uniqueIndex:idx_name_city"`
	WarehouseID        *uint   `json:"warehouse_id" gorm:"column:warehouse_id;default:null"`
	WeatherDescription string  `json:"weather_description" gorm:"-"` // populated by API
}
