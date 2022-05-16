package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Name               string  `json:"name" gorm:"not null"`
	Description        string  `json:"description" gorm:"column:description;not null;default:''"`
	Price              float32 `json:"price" gorm:"column:price;not null;default:0"`
	Count              int     `json:"count" gorm:"column:count;not null;default:0"`
	City               string  `json:"city" gorm:"column:city;not null;default:''"`
	WarehouseID        *uint   `json:"warehouse_id" gorm:"column:warehouse_id;default:null"`
	WeatherDescription string  `json:"weather_description" gorm:"-"` // populated by API
}
