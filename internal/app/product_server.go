package app

import (
	"sirawit/shop/internal/config"
	"sirawit/shop/internal/service"
	"sirawit/shop/pkg/pb"
)

type productServer struct {
	productService service.ProductService
	pb.UnimplementedProductServiceServer
	config       config.ProductConfig
	tokenManager service.TokenManager
}

func NewProductServer(p service.ProductService, config config.ProductConfig) *productServer {
	tokenManager := service.NewTokenManager(config.Sign)
	return &productServer{
		productService: p,
		config:         config,
		tokenManager:   tokenManager,
	}
}
