package services

import (
	"errors"

	db "github.com/sriharivishnu/shopify-challenge/external"
	apiModels "github.com/sriharivishnu/shopify-challenge/models/api"
	dbModels "github.com/sriharivishnu/shopify-challenge/models/db"
)

type ItemLayer interface {
	CreateItem(payload apiModels.CreateItemPayload) (dbModels.Item, error)
	UpdateItem(itemId uint, payload apiModels.UpdateItemPayload) (dbModels.Item, error)
	GetItemById(itemId uint) (dbModels.Item, error)
	GetItems() ([]dbModels.Item, error)
	DeleteItem(itemId uint) error
}

type ItemService struct{}

// could add a cache to improve performance by avoiding database, but kept it simple for now
func (service *ItemService) CreateItem(payload apiModels.CreateItemPayload) (dbModels.Item, error) {
	item := dbModels.Item{
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
		City:        payload.City,
	}
	if res := db.DbConn.Create(&item); res.Error != nil {
		return dbModels.Item{}, res.Error
	}
	return item, nil
}

func (service *ItemService) UpdateItem(itemId uint, payload apiModels.UpdateItemPayload) (dbModels.Item, error) {
	item := dbModels.Item{}
	item.ID = itemId
	res := db.DbConn.Model(&item).Updates(
		dbModels.Item{
			Name:        payload.Name,
			Description: payload.Description,
			Price:       payload.Price,
			Quantity:    payload.Quantity,
			City:        payload.City,
		},
	)
	if res.Error != nil {
		return dbModels.Item{}, res.Error
	}
	if res.RowsAffected == 0 {
		return dbModels.Item{}, errors.New("No item found")
	}
	return item, nil
}

func (service *ItemService) GetItemById(itemId uint) (dbModels.Item, error) {
	item := dbModels.Item{}
	if res := db.DbConn.First(&item, itemId); res.Error != nil {
		return dbModels.Item{}, res.Error
	}
	return item, nil
}

func (service *ItemService) GetItems() ([]dbModels.Item, error) {
	items := []dbModels.Item{}
	if res := db.DbConn.Find(&items); res.Error != nil {
		return []dbModels.Item{}, res.Error
	}
	return items, nil
}

func (service *ItemService) DeleteItem(itemId uint) error {
	res := db.DbConn.Unscoped().Delete(&dbModels.Item{}, itemId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("No item found")
	}
	return nil
}
