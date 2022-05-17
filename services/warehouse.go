package services

import (
	db "github.com/sriharivishnu/shopify-challenge/external"
	apiModels "github.com/sriharivishnu/shopify-challenge/models/api"
	dbModels "github.com/sriharivishnu/shopify-challenge/models/db"
)

type WarehouseLayer interface {
	CreateWarehouse(payload apiModels.CreateWarehousePayload) (dbModels.Warehouse, error)
	GetWarehouseById(warehouseId uint) (dbModels.Warehouse, error)
	GetAllWarehouses() ([]dbModels.Warehouse, error)
	AddItemToWarehouse(warehouseId uint, itemId uint) (dbModels.Item, error)
}

type WarehouseService struct{}

func (service *WarehouseService) CreateWarehouse(payload apiModels.CreateWarehousePayload) (dbModels.Warehouse, error) {
	warehouse := dbModels.Warehouse{
		Name:        payload.Name,
		Description: payload.Description,
		Longitude:   payload.Lon,
		Latitude:    payload.Lat,
	}
	if res := db.DbConn.Create(&warehouse); res.Error != nil {
		return dbModels.Warehouse{}, res.Error
	}
	return warehouse, nil
}

func (service *WarehouseService) GetWarehouseById(warehouseId uint) (dbModels.Warehouse, error) {
	warehouse := dbModels.Warehouse{}
	if res := db.DbConn.First(&warehouse, warehouseId); res.Error != nil {
		return dbModels.Warehouse{}, res.Error
	}
	db.DbConn.Model(&warehouse).Association("Items").Find(&warehouse.Items)

	return warehouse, nil
}

func (service *WarehouseService) GetAllWarehouses() ([]dbModels.Warehouse, error) {
	warehouses := []dbModels.Warehouse{}
	if res := db.DbConn.Find(&warehouses); res.Error != nil {
		return []dbModels.Warehouse{}, res.Error
	}
	return warehouses, nil
}

func (service *WarehouseService) AddItemToWarehouse(warehouseId uint, itemId uint) (dbModels.Item, error) {
	warehouse := dbModels.Warehouse{}
	if res := db.DbConn.First(&warehouse, warehouseId); res.Error != nil {
		return dbModels.Item{}, res.Error
	}
	item := dbModels.Item{}
	if res := db.DbConn.First(&item, itemId); res.Error != nil {
		return dbModels.Item{}, res.Error
	}
	err := db.DbConn.Model(&warehouse).Association("Items").Append(&item)
	if err != nil {
		return dbModels.Item{}, err
	}
	return item, nil
}
