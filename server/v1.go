package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sriharivishnu/shopify-challenge/controllers"
	"github.com/sriharivishnu/shopify-challenge/external"
	"github.com/sriharivishnu/shopify-challenge/services"
)

func SetUpV1(router *gin.Engine) {
	// set up controllers and inject dependencies
	weatherCache := &external.InMemoryCache{}
	weatherCache.Init()
	itemController := controllers.ItemController{
		ItemService: &services.ItemService{},
		WeatherService: &services.WeatherService{
			Cache: weatherCache,
			HttpClient: &http.Client{
				Timeout: time.Second * 2, // Timeout after 2 seconds
			},
		},
	}

	warehouseController := controllers.WarehouseController{
		WarehouseService: &services.WarehouseService{},
	}

	// Set up routes
	v1 := router.Group("v1")

	// items
	items := v1.Group("items")
	items.GET("/", itemController.GetItems)
	items.GET("/:item_id", itemController.GetItem)
	items.POST("/", itemController.CreateItem)
	items.PUT("/:item_id", itemController.UpdateItem)
	items.DELETE("/:item_id", itemController.DeleteItem)

	// warehouses
	warehouses := v1.Group("warehouses")
	warehouses.GET("/", warehouseController.GetAll)
	warehouses.GET("/:warehouse_id", warehouseController.Get)
	warehouses.POST("/", warehouseController.Create)
	warehouses.POST("/:warehouse_id/items", warehouseController.AddItemToWarehouse)
}
