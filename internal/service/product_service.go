package service

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
)

type ProductService interface {
	GetProducts(id uint64) (*repository.GetProductsRes, error)
	CreateProduct(input model.Product) (*model.Product, error)
	GetProduct(id uint64) (*model.Product, error)
	DeleteCart(id uint64) error
	ManageCart(input model.Cart) (*model.Cart, error)
	GetCarts(username string) ([]model.Cart, error)
}

type productService struct {
	db     repository.ProductQuery
	config config.ProductConfig
}

func NewProductService(db repository.ProductQuery, config config.ProductConfig) ProductService {
	return &productService{
		db:     db,
		config: config,
	}
}
