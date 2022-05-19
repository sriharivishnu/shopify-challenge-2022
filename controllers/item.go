package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	models "github.com/sriharivishnu/shopify-challenge/models/api"
	"github.com/sriharivishnu/shopify-challenge/services"
	"github.com/sriharivishnu/shopify-challenge/utils"
)

// TYPES
type ItemController struct {
	ItemService    services.ItemLayer
	WeatherService services.WeatherLayer
}

// FUNCTIONS
func (i *ItemController) CreateItem(c *gin.Context) {
	createItemPayload := models.CreateItemPayload{}

	if err := c.BindJSON(&createItemPayload); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	if err := createItemPayload.Validate(); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	item, err := i.ItemService.CreateItem(createItemPayload)

	if err != nil {
		utils.RespondSQLError(c, err)
		return
	}

	c.JSON(200, gin.H{"message": "Item created successfully", "item": item})
}

func (i *ItemController) UpdateItem(c *gin.Context) {
	itemIdRaw := c.Param("item_id")
	var err error
	var itemId int

	if itemId, err = strconv.Atoi(itemIdRaw); err != nil || itemId <= 0 {
		if err != nil {
			utils.RespondError(c, err, http.StatusBadRequest)
		} else {
			utils.RespondError(c, errors.New("Invalid Item ID"), http.StatusBadRequest)
		}
		return
	}

	updateItemPayload := models.UpdateItemPayload{}
	if err := c.BindJSON(&updateItemPayload); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	item, err := i.ItemService.UpdateItem(uint(itemId), updateItemPayload)

	if err != nil {
		utils.RespondSQLError(c, err)
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
		utils.RespondSQLError(c, err)
		return
	}
	c.JSON(200, gin.H{"item": item})
}

func (i *ItemController) GetItems(c *gin.Context) {
	items, err := i.ItemService.GetItems()
	if err != nil {
		utils.RespondSQLError(c, err)
		return
	}
	for idx := range items {
		weather, err := i.WeatherService.FetchWeather(items[idx].City)
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
		utils.RespondSQLError(c, err)
		return
	}
	c.JSON(200, gin.H{"message": "Item deleted successfully"})
}
