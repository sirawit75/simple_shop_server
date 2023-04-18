package repository

import (
	"sirawit/shop/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GetProductsRes struct {
	Products []model.Product
	HasMore  bool
}

type ProductQuery interface {
	GetProducts(id uint64) (*GetProductsRes, error)
	GetProduct(id uint64) (*model.Product, error)
	CreateProduct(input model.Product) (*model.Product, error)
	ManageCart(input model.Cart) (*model.Cart, error)
	DeleteCart(id uint64) error
	GetCarts(username string) ([]model.Cart, error)
}

type productQuery struct {
	db *gorm.DB
}

func NewProductQuery(db *gorm.DB) ProductQuery {
	db.AutoMigrate(&model.Product{}, &model.Cart{})
	return &productQuery{db}
}

func ConnectToProductDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
