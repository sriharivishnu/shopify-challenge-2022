package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sriharivishnu/shopify-challenge/services"
	"github.com/sriharivishnu/shopify-challenge/utils"
)

type WarehouseController struct {
	WarehouseService services.WarehouseLayer
}

type CreateWarehousePayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Lon         float32 `json:"lon"`
	Lat         float32 `json:"lat"`
}

type AddItemPayload struct {
	ItemID uint `json:"item_id"`
}

func (w *WarehouseController) Create(c *gin.Context) {
	createWarehousePayload := CreateWarehousePayload{}

	if err := c.BindJSON(&createWarehousePayload); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	warehouse, err := w.WarehouseService.CreateWarehouse(
		createWarehousePayload.Name,
		createWarehousePayload.Description,
		createWarehousePayload.Lon,
		createWarehousePayload.Lat,
	)
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warehouse created successfully", "warehouse": warehouse})
}

func (w *WarehouseController) GetAll(c *gin.Context) {
	warehouses, err := w.WarehouseService.GetAllWarehouses()
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warehouses fetched successfully", "warehouses": warehouses})
}

func (w *WarehouseController) Get(c *gin.Context) {
	warehouseID, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	warehouse, err := w.WarehouseService.GetWarehouseById(uint(warehouseID))
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Warehouse fetched successfully", "warehouse": warehouse})
}

func (w *WarehouseController) AddItemToWarehouse(c *gin.Context) {
	warehouseID, err := strconv.Atoi(c.Param("warehouse_id"))
	if err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	addItemPayload := AddItemPayload{}
	if err := c.BindJSON(&addItemPayload); err != nil {
		utils.RespondError(c, err, http.StatusBadRequest)
		return
	}

	item, err := w.WarehouseService.AddItemToWarehouse(uint(warehouseID), uint(addItemPayload.ItemID))
	if err != nil {
		utils.RespondError(c, err, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to warehouse successfully", "item": item})
}
