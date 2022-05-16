package services

import (
	db "github.com/sriharivishnu/shopify-challenge/external"
	models "github.com/sriharivishnu/shopify-challenge/models/db"
)

type WarehouseLayer interface {
	CreateWarehouse(name, description string, lon, lat float32) (models.Warehouse, error)
	GetWarehouseById(warehouseId uint) (models.Warehouse, error)
	GetAllWarehouses() ([]models.Warehouse, error)
	AddItemToWarehouse(warehouseId uint, itemId uint) (models.Item, error)
}

type WarehouseService struct{}

func (service *WarehouseService) CreateWarehouse(name, description string, lon, lat float32) (models.Warehouse, error) {
	warehouse := models.Warehouse{Name: name, Description: description, Longitude: lon, Latitude: lat}
	if res := db.DbConn.Create(&warehouse); res.Error != nil {
		return models.Warehouse{}, res.Error
	}
	return warehouse, nil
}

func (service *WarehouseService) GetWarehouseById(warehouseId uint) (models.Warehouse, error) {
	warehouse := models.Warehouse{}
	if res := db.DbConn.First(&warehouse, warehouseId); res.Error != nil {
		return models.Warehouse{}, res.Error
	}
	db.DbConn.Model(&warehouse).Association("Items").Find(&warehouse.Items)

	return warehouse, nil
}

func (service *WarehouseService) GetAllWarehouses() ([]models.Warehouse, error) {
	warehouses := []models.Warehouse{}
	if res := db.DbConn.Find(&warehouses); res.Error != nil {
		return []models.Warehouse{}, res.Error
	}
	return warehouses, nil
}

func (service *WarehouseService) AddItemToWarehouse(warehouseId uint, itemId uint) (models.Item, error) {
	warehouse := models.Warehouse{}
	if res := db.DbConn.First(&warehouse, warehouseId); res.Error != nil {
		return models.Item{}, res.Error
	}
	item := models.Item{}
	if res := db.DbConn.First(&item, itemId); res.Error != nil {
		return models.Item{}, res.Error
	}
	err := db.DbConn.Model(&warehouse).Association("Items").Append(&item)
	if err != nil {
		return models.Item{}, err
	}
	return item, nil
}
