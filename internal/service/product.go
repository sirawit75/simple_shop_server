package service

import (
	"sirawit/shop/internal/model"
	"sirawit/shop/internal/repository"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (p *productService) DeleteCart(id uint64) error {
	if err := p.db.DeleteCart(id); err != nil {
		return status.Errorf(codes.Internal, "failed to delete cart %v", err)
	}
	return nil
}
func (p *productService) GetCarts(username string) ([]model.Cart, error) {
	result, err := p.db.GetCarts(username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get carts %v", err)
	}
	return result, nil
}

func (p *productService) ManageCart(input model.Cart) (*model.Cart, error) {
	product, err := p.db.ManageCart(input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add to cart %v", err)
	}
	return product, nil
}

func (p *productService) GetProducts(id uint64) (*repository.GetProductsRes, error) {
	products, err := p.db.GetProducts(id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get product %v", err)
	}
	return products, nil
}

func (p *productService) CreateProduct(input model.Product) (*model.Product, error) {
	result, err := p.db.CreateProduct(input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create product %v", err)
	}
	return result, nil
}

func (p *productService) GetProduct(id uint64) (*model.Product, error) {
	product, err := p.db.GetProduct(id)
	if err != nil {
		if strings.Contains(err.Error(), NotFoundStatus) {
			return nil, status.Error(codes.NotFound, ProductNotFound)
		}
		return nil, status.Errorf(codes.Internal, "failed to find product %v", err)

	}
	return product, nil
}
