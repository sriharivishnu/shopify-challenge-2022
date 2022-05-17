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
	apiModels "github.com/sriharivishnu/shopify-challenge/models/api"
	dbModels "github.com/sriharivishnu/shopify-challenge/models/db"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestItemCreateSuccess(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	params := []gin.Param{
		{
			Key:   "item_id",
			Value: "item_id",
		},
	}
	dummyItem := dbModels.Item{
		Name:        "itemName",
		Description: "sample description",
		Price:       10.99,
		Quantity:    10,
		City:        "city",
	}
	dummyItem.ID = 1234

	t.Run("ItemCreateSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		utils.MockJsonPost(ctx, gin.H{"name": "itemName", "description": "sample description", "price": 10.99, "quantity": 10, "city": "city"})

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("CreateItem", apiModels.CreateItemPayload{
			Name:        "itemName",
			Description: "sample description",
			Price:       10.99,
			Quantity:    10,
			City:        "city",
		}).Return(dummyItem, nil)

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.CreateItem(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"item":{"ID":1234,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"itemName","description":"sample description","price":10.99,"quantity":10,"city":"city","warehouse_id":null,"weather_description":""},"message":"Item created successfully"}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("ItemCreateFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		utils.MockJsonPost(ctx, gin.H{"name": "itemName", "description": "sample description", "price": 10.99, "quantity": 10, "city": "city"})

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("CreateItem", apiModels.CreateItemPayload{
			Name:        "itemName",
			Description: "sample description",
			Price:       10.99,
			Quantity:    10,
			City:        "city",
		}).Return(dbModels.Item{}, errors.New("error"))

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.CreateItem(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}

func TestUpdateItem(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	params := []gin.Param{
		{
			Key:   "item_id",
			Value: "1234",
		},
	}
	dummyItem := dbModels.Item{
		Name:        "itemName",
		Description: "sample description",
		Price:       10.99,
		Quantity:    100,
		City:        "city",
	}
	dummyItem.ID = 1234
	t.Run("UpdateItemSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		utils.MockJsonPost(ctx, gin.H{"name": "itemName", "description": "sample description", "price": 10.99, "quantity": 100, "city": "city"})

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("UpdateItem", uint(1234),
			apiModels.UpdateItemPayload{
				Name:        "itemName",
				Description: "sample description",
				Price:       10.99,
				Quantity:    100,
				City:        "city",
			}).Return(dummyItem, nil)

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.UpdateItem(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"item":{"ID":1234,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"itemName","description":"sample description","price":10.99,"quantity":100,"city":"city","warehouse_id":null,"weather_description":""},"message":"Item updated successfully"}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("UpdateItemFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		utils.MockJsonPost(ctx, gin.H{"name": "itemName", "description": "sample description", "price": 10.99, "quantity": 100, "city": "city"})

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("UpdateItem", uint(1234),
			apiModels.UpdateItemPayload{
				Name:        "itemName",
				Description: "sample description",
				Price:       10.99,
				Quantity:    100,
				City:        "city",
			}).Return(dbModels.Item{}, errors.New("error"))

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.UpdateItem(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}

func TestDeleteItem(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	params := []gin.Param{
		{
			Key:   "item_id",
			Value: "1234",
		},
	}
	t.Run("DeleteItemSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("DeleteItem", uint(1234)).Return(nil)

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.DeleteItem(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"message":"Item deleted successfully"}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("DeleteItemFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("DeleteItem", uint(1234)).Return(errors.New("error"))

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.DeleteItem(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}

func TestGetItem(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	params := []gin.Param{
		{
			Key:   "item_id",
			Value: "1234",
		},
	}
	dummyItem := dbModels.Item{
		Name:        "itemName",
		Description: "sample description",
		Price:       10.99,
		Quantity:    100,
		City:        "city",
	}
	dummyItem.ID = 1234
	t.Run("GetItemSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("GetItemById", uint(1234)).Return(dummyItem, nil)

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.GetItem(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"item":{"ID":1234,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"name":"itemName","description":"sample description","price":10.99,"quantity":100,"city":"city","warehouse_id":null,"weather_description":""}}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("GetItemFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("GetItemById", uint(1234)).Return(dbModels.Item{}, errors.New("error"))

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.GetItem(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("GetItemNotFound", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Params = params

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("GetItemById", uint(1234)).Return(dbModels.Item{}, gorm.ErrRecordNotFound)

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.GetItem(ctx)

		assert.Equal(t, 404, w.Code)
		expected := `{"error":"Resource not found!"}`
		assert.Equal(t, expected, w.Body.String())
	})
}

func TestGetAllItems(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	gin.SetMode(gin.TestMode)
	t.Run("GetAllItemsSuccess", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("GetItems").Return([]dbModels.Item{}, nil)

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.GetItems(ctx)

		assert.Equal(t, 200, w.Code)
		expected := `{"items":[]}`
		assert.Equal(t, expected, w.Body.String())
	})

	t.Run("GetAllItemsFail", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		mockItemService := mocks.ItemLayer{}
		mockItemService.On("GetItems").Return([]dbModels.Item{}, errors.New("error"))

		itemController := ItemController{
			ItemService: &mockItemService,
		}

		itemController.GetItems(ctx)

		assert.Equal(t, 500, w.Code)
		expected := `{"error":"error"}`
		assert.Equal(t, expected, w.Body.String())
	})
}
