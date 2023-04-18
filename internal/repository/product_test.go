package repository

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/random"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var testProductDB *gorm.DB
var testPruductQuery ProductQuery

func TestConnectToProductDB(t *testing.T) {
	var err error
	config, err := config.LoadProductConfig("../../cmd/product")
	assert.NoError(t, err)
	testProductDB, err = ConnectToProductDB(config.DSN)
	assert.NoError(t, err)
	testPruductQuery = NewProductQuery(testProductDB)

}

func TestGetCarts(t *testing.T) {
	product, err := testPruductQuery.CreateProduct(model.Product{
		Name:     random.RandomUsername(),
		Price:    float64(random.RandomUInt64(0, 1000)),
		Details:  random.RandomString(20),
		ImageUrl: random.RandomString(20),
	})
	assert.NoError(t, err)
	item := model.Cart{
		Username:  random.RandomUsername(),
		Qty:       int(random.RandomUInt64(0, 100)),
		ProductID: product.ID,
	}
	_, err = testPruductQuery.ManageCart(item)
	assert.NoError(t, err)
	result, err := testPruductQuery.GetCarts(item.Username)
	assert.NoError(t, err)
	assert.Equal(t, item.ProductID, result[0].ProductID)
	assert.Equal(t, item.Qty, result[0].Qty)
}

func TestManageCart(t *testing.T) {
	t.Run("delete cart", func(t *testing.T) {
		product, err := testPruductQuery.CreateProduct(model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		})
		assert.NoError(t, err)
		item := model.Cart{
			Username:  random.RandomUsername(),
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product.ID,
		}
		_, err = testPruductQuery.ManageCart(item)
		assert.NoError(t, err)
		err = testPruductQuery.DeleteCart(item.ID)
		assert.Nil(t, err)

	})
	t.Run("new item", func(t *testing.T) {
		product1, err := testPruductQuery.CreateProduct(model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		})
		assert.NoError(t, err)

		product2, err := testPruductQuery.CreateProduct(model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		})
		assert.NoError(t, err)
		item1 := model.Cart{
			Username:  random.RandomUsername(),
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product1.ID,
		}

		item2 := model.Cart{
			Username:  item1.Username,
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product2.ID,
		}
		result, err := testPruductQuery.ManageCart(item1)
		assert.NoError(t, err)
		assert.Equal(t, item1.Qty, result.Qty)
		assert.Equal(t, item1.ProductID, result.ProductID)

		result, err = testPruductQuery.ManageCart(item2)
		assert.NoError(t, err)
		assert.Equal(t, item2.Qty, result.Qty)
		assert.Equal(t, item2.ProductID, result.ProductID)
	})
	t.Run("new cart", func(t *testing.T) {
		product, err := testPruductQuery.CreateProduct(model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		})
		assert.NoError(t, err)
		item := model.Cart{
			Username:  random.RandomUsername(),
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product.ID,
		}
		result, err := testPruductQuery.ManageCart(item)
		assert.NoError(t, err)
		assert.Equal(t, item.Qty, result.Qty)
		assert.Equal(t, item.ProductID, result.ProductID)
	})

	t.Run("add to cart", func(t *testing.T) {
		product, err := testPruductQuery.CreateProduct(model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		})
		assert.NoError(t, err)
		item1 := model.Cart{
			Username:  random.RandomUsername(),
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product.ID,
		}

		item2 := model.Cart{
			Username:  item1.Username,
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product.ID,
		}
		_, err = testPruductQuery.ManageCart(item1)
		assert.NoError(t, err)
		result, err := testPruductQuery.ManageCart(item2)
		assert.NoError(t, err)
		assert.Equal(t, item1.Qty+item2.Qty, result.Qty)
		assert.Equal(t, item1.ProductID, result.ProductID)
		assert.Equal(t, item2.ProductID, result.ProductID)
	})
	t.Run("remove item from cart", func(t *testing.T) {
		product, err := testPruductQuery.CreateProduct(model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		})
		assert.NoError(t, err)
		item1 := model.Cart{
			Username:  random.RandomUsername(),
			Qty:       int(random.RandomUInt64(0, 100)),
			ProductID: product.ID,
		}

		item1_ := model.Cart{
			Username:  item1.Username,
			Qty:       -item1.Qty + 1,
			ProductID: product.ID,
		}
		_, err = testPruductQuery.ManageCart(item1)
		assert.NoError(t, err)
		result, err := testPruductQuery.ManageCart(item1_)
		assert.NoError(t, err)
		assert.Equal(t, 1, result.Qty)
		assert.Equal(t, item1.ProductID, result.ProductID)
		assert.Equal(t, item1_.ProductID, result.ProductID)
	})
}

func TestCreateProduct(t *testing.T) {
	qty := 3
	for i := 0; i < qty; i++ {
		product := model.Product{
			Name:     random.RandomUsername(),
			Price:    float64(random.RandomUInt64(0, 1000)),
			Details:  random.RandomString(20),
			ImageUrl: random.RandomString(20),
		}
		result, err := testPruductQuery.CreateProduct(product)
		assert.NoError(t, err)
		assert.Equal(t, result.Details, product.Details)
		assert.Equal(t, result.Price, product.Price)
		assert.Equal(t, result.ImageUrl, product.ImageUrl)
		assert.Equal(t, result.Name, product.Name)
		assert.NotNil(t, result.ID)
	}
}
