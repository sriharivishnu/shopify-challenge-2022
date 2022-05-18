package controllers

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sriharivishnu/shopify-challenge/mocks"
	utils "github.com/sriharivishnu/shopify-challenge/mocks/helpers"
	models "github.com/sriharivishnu/shopify-challenge/models/api"
	dbModels "github.com/sriharivishnu/shopify-challenge/models/db"
	"github.com/stretchr/testify/assert"
)

func TestWarehouseCreate(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	dummyWarehouse := dbModels.Warehouse{
		Name:        "warehouseName",
		Description: "sample description",
		Longitude:   1.1321,
		Latitude:    1.4567,
	}
	t.Run("WarehouseCreateSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		utils.MockJsonPost(ctx, gin.H{"name": "warehouseName", "description": "sample description", "longitude": 1.1321, "latitude": 1.4567})

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("CreateWarehouse", models.CreateWarehousePayload{
			Name:        "warehouseName",
			Description: "sample description",
			Longitude:   1.1321,
			Latitude:    1.4567,
		}).Return(dummyWarehouse, nil)

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.Create(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"message":"Warehouse created successfully","warehouse":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"warehouseName","description":"sample description","longitude":1.1321,"latitude":1.4567}}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("WarehouseCreateFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		utils.MockJsonPost(ctx, gin.H{"name": "warehouseName", "description": "sample description", "longitude": 1.1321, "latitude": 1.4567})

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("CreateWarehouse", models.CreateWarehousePayload{
			Name:        "warehouseName",
			Description: "sample description",
			Longitude:   1.1321,
			Latitude:    1.4567,
		}).Return(dbModels.Warehouse{}, errors.New("error"))

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.Create(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}

func TestWarehouseGet(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	gin.SetMode(gin.TestMode)
	params := []gin.Param{
		{
			Key:   "warehouse_id",
			Value: "1",
		},
	}

	dummyWarehouse := dbModels.Warehouse{
		Name:        "warehouseName",
		Description: "sample description",
		Longitude:   1.1321,
		Latitude:    1.4567,
	}
	t.Run("WarehouseGetSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("GetWarehouseById", uint(1)).Return(dummyWarehouse, nil)

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.Get(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"message":"Warehouse fetched successfully","warehouse":{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"warehouseName","description":"sample description","longitude":1.1321,"latitude":1.4567}}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("WarehouseGetFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("GetWarehouseById", uint(1)).Return(dbModels.Warehouse{}, errors.New("error"))

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.Get(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})

}

func TestGetAllWarehouses(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	dummyWarehouses := []dbModels.Warehouse{
		{
			Name:        "warehouseName",
			Description: "sample description",
			Longitude:   1.1321,
			Latitude:    1.4567,
		},
	}
	t.Run("GetAllWarehousesSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("GetAllWarehouses").Return(dummyWarehouses, nil)

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.GetAll(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"message":"Warehouses fetched successfully","warehouses":[{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"warehouseName","description":"sample description","longitude":1.1321,"latitude":1.4567}]}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("GetAllWarehousesFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("GetAllWarehouses").Return([]dbModels.Warehouse{}, errors.New("error"))

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.GetAll(ctx)
		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}

func TestAddItemToWarehouse(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	params := []gin.Param{
		{
			Key:   "warehouse_id",
			Value: "1",
		},
	}
	warehouseId := uint(1)
	dummyItem := dbModels.Item{
		WarehouseID: &warehouseId,
	}
	dummyItem.ID = 1

	t.Run("AddItemToWarehouseSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		utils.MockJsonPost(ctx, gin.H{"item_id": 1})

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("AddItemToWarehouse", uint(1), uint(1)).Return(dummyItem, nil)

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.AddItemToWarehouse(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"item":{"ID":1,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"","description":"","price":0,"quantity":0,"city":"","warehouse_id":1,"weather_description":""},"message":"Item added to warehouse successfully"}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("AddItemToWarehouseFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		utils.MockJsonPost(ctx, gin.H{"item_id": 1})

		mockWarehouseService := mocks.WarehouseLayer{}
		mockWarehouseService.On("AddItemToWarehouse", uint(1), uint(1)).Return(dbModels.Item{}, errors.New("error"))

		warehouseController := WarehouseController{
			WarehouseService: &mockWarehouseService,
		}

		warehouseController.AddItemToWarehouse(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}
