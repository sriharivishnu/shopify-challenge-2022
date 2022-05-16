package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sriharivishnu/shopify-challenge/external"
	"github.com/sriharivishnu/shopify-challenge/services"
	"github.com/sriharivishnu/shopify-challenge/utils"
)

// TYPES
type ItemController struct {
	ItemService    services.ItemLayer
	WeatherService services.WeatherLayer
	StorageService external.Storage
}

type CreateItemPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Count       int     `json:"count"`
	City        string  `json:"city"`
}

type UpdateItemPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Count       int     `json:"count"`
	City        string  `json:"city"`
}

// FUNCTIONS
func (i *ItemController) CreateItem(c *gin.Context) {
	createItemPayload := CreateItemPayload{}

	if err := c.BindJSON(&createItemPayload); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	item, err := i.ItemService.CreateItem(
		createItemPayload.Name,
		createItemPayload.City,
		createItemPayload.Description,
		createItemPayload.Count,
		createItemPayload.Price,
	)

	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{"message": "Item created successfully", "item": item})
}

func (i *ItemController) UpdateItem(c *gin.Context) {
	itemIdRaw := c.Param("item_id")
	var err error
	var itemId int

	if itemId, err = strconv.Atoi(itemIdRaw); err != nil || itemId <= 0 {
		utils.RespondError(c, nil, http.StatusBadRequest)
		return
	}

	updateItemPayload := UpdateItemPayload{}
	if err := c.BindJSON(&updateItemPayload); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	item, err := i.ItemService.UpdateItem(
		uint(itemId),
		updateItemPayload.Name,
		updateItemPayload.City,
		updateItemPayload.Description,
		updateItemPayload.Count,
		updateItemPayload.Price,
	)

	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(200, gin.H{"message": "Item updated successfully", "item": item})
}

func (i *ItemController) GetItem(c *gin.Context) {
	itemIdRaw := c.Param("item_id")
	var err error
	var itemId int

	if itemId, err = strconv.Atoi(itemIdRaw); err != nil || itemId <= 0 {
		utils.RespondError(c, nil, http.StatusBadRequest)
		return
	}
	item, err := i.ItemService.GetItemById(uint(itemId))
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(200, gin.H{"item": item})
}

func (i *ItemController) GetItems(c *gin.Context) {
	items, err := i.ItemService.GetItems()
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	for idx := range items {
		weather, err := i.WeatherService.FetchWeather(items[idx].City, 80.5204, 43.4643)
		if err == nil {
			items[idx].WeatherDescription = weather.Weather[0].Description
		} else {
			items[idx].WeatherDescription = "No weather data available"
		}
	}
	c.JSON(200, gin.H{"items": items})
}

func (i *ItemController) DeleteItem(c *gin.Context) {
	itemIdRaw := c.Param("item_id")
	var err error
	var itemId int

	if itemId, err = strconv.Atoi(itemIdRaw); err != nil || itemId <= 0 {
		utils.RespondError(c, nil, http.StatusBadRequest)
		return
	}

	err = i.ItemService.DeleteItem(uint(itemId))
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}
	c.JSON(200, gin.H{"message": "Item deleted successfully"})
}
