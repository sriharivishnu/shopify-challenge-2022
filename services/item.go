package services

import (
	"errors"

	db "github.com/sriharivishnu/shopify-challenge/external"
	models "github.com/sriharivishnu/shopify-challenge/models/db"
)

type ItemLayer interface {
	CreateItem(name, city, description string, count int, price float32) (models.Item, error)
	UpdateItem(itemId uint, name, city, description string, count int, price float32) (models.Item, error)
	GetItemById(itemId uint) (models.Item, error)
	GetItems() ([]models.Item, error)
	DeleteItem(itemId uint) error
}

type ItemService struct{}

// could add a cache to improve performance by avoiding database, but kept it simple for now
func (service *ItemService) CreateItem(name, city, description string, count int, price float32) (models.Item, error) {
	item := models.Item{Name: name, City: city, Description: description, Count: count, Price: price}
	if res := db.DbConn.Create(&item); res.Error != nil {
		return models.Item{}, res.Error
	}
	return item, nil
}

func (service *ItemService) UpdateItem(itemId uint, city, name, description string, count int, price float32) (models.Item, error) {
	item := models.Item{}
	item.ID = itemId
	res := db.DbConn.Model(&item).Updates(
		models.Item{
			Name:        name,
			Description: description,
			Price:       price,
			Count:       count,
			City:        city,
		},
	)
	if res.Error != nil {
		return models.Item{}, res.Error
	}
	if res.RowsAffected == 0 {
		return models.Item{}, errors.New("No item found")
	}
	return item, nil
}

func (service *ItemService) GetItemById(itemId uint) (models.Item, error) {
	item := models.Item{}
	if res := db.DbConn.First(&item, itemId); res.Error != nil {
		return models.Item{}, res.Error
	}
	return item, nil
}

func (service *ItemService) GetItems() ([]models.Item, error) {
	items := []models.Item{}
	if res := db.DbConn.Find(&items); res.Error != nil {
		return []models.Item{}, res.Error
	}
	return items, nil
}

func (service *ItemService) DeleteItem(itemId uint) error {
	res := db.DbConn.Delete(&models.Item{}, itemId)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("No item found")
	}
	return nil
}
