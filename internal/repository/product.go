package repository

import (
	"sirawit/shop/internal/model"
)

const (
	limit = 5
)

func (p *productQuery) GetCarts(username string) ([]model.Cart, error) {
	var cart []model.Cart
	if err := p.db.Preload("Product").Where("username = ?", username).Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (p *productQuery) DeleteCart(id uint64) error {
	var cart model.Cart
	if err := p.db.Delete(&cart, id).Error; err != nil {
		return err
	}
	return nil
}
func (p *productQuery) ManageCart(input model.Cart) (*model.Cart, error) {
	var findCart []model.Cart
	if err := p.db.Preload("Product").Where("username = ?", input.Username).Find(&findCart).Error; err != nil {
		return nil, err
	}
	//new cart
	if len(findCart) == 0 {
		if err := p.db.Create(&input).Error; err != nil {
			return nil, err
		}
		return &input, nil
	}
	// manage cart
	for i := range findCart {
		if findCart[i].ProductID == input.ProductID {
			findCart[i].Qty += input.Qty
			if err := p.db.Save(&findCart).Error; err != nil {
				return nil, err
			}
			return &findCart[i], nil
		}
	}
	//add new item to cart
	findCart = append(findCart, input)
	if err := p.db.Save(&findCart).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (p *productQuery) GetProducts(id uint64) (*GetProductsRes, error) {
	var products []model.Product
	if err := p.db.Where("id > ?", id).Limit(limit + 1).Find(&products).Error; err != nil {
		return nil, err
	}
	return &GetProductsRes{
		Products: products[:limit],
		HasMore:  len(products) == limit+1,
	}, nil
}

func (p *productQuery) GetProduct(id uint64) (*model.Product, error) {
	var product model.Product
	if err := p.db.Where("id = ?", id).First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil

}

func (p *productQuery) CreateProduct(input model.Product) (*model.Product, error) {
	if err := p.db.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}
