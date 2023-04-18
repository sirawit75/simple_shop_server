package app

import (
	"context"
	"sirawit/shop/internal/model"
	"sirawit/shop/pkg/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func convertToProductRes(input model.Product) *pb.Product {
	return &pb.Product{
		ID:       input.ID,
		Name:     input.Name,
		Price:    float32(input.Price),
		Details:  input.Details,
		ImageUrl: input.ImageUrl,
	}
}

func (p *productServer) DeleteCart(ctx context.Context, req *pb.DeleteCartReq) (*emptypb.Empty, error) {
	_, err := p.getUsernameFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	if err := p.productService.DeleteCart(req.GetId()); err != nil {
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}
func (p *productServer) GetCarts(ctx context.Context, empty *emptypb.Empty) (*pb.Carts, error) {
	username, err := p.getUsernameFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	result, err := p.productService.GetCarts(username)
	if err != nil {
		return nil, err
	}
	var cart []*pb.Cart
	for i := 0; i < len(result); i++ {
		cart = append(cart, &pb.Cart{
			ID:        result[i].ID,
			ProductId: result[i].ProductID,
			Qty:       int64(result[i].Qty),
			Product:   convertToProductRes(result[i].Product),
		})
	}
	return &pb.Carts{
		Carts: cart,
	}, nil
}

func (p *productServer) GetProduct(ctx context.Context, req *pb.GetProductsReq) (*pb.Product, error) {
	result, err := p.productService.GetProduct(req.GetId())
	if err != nil {
		return nil, err
	}
	return convertToProductRes(*result), nil
}

func (p *productServer) GetProducts(ctx context.Context, req *pb.GetProductsReq) (*pb.GetProductsRes, error) {
	result, err := p.productService.GetProducts(req.GetId())
	if err != nil {
		return nil, err
	}
	var products []*pb.Product
	for i := 0; i < len(result.Products); i++ {
		products = append(products, convertToProductRes(result.Products[i]))
	}
	return &pb.GetProductsRes{
		Products: products,
		HasMore:  result.HasMore,
	}, nil
}

func (p *productServer) CreateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	username, err := p.getUsernameFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	if username != p.config.Admin {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	result, err := p.productService.CreateProduct(model.Product{
		Name:     req.GetName(),
		Details:  req.GetDetails(),
		Price:    float64(req.GetPrice()),
		ImageUrl: req.GetImageUrl(),
	})
	if err != nil {
		return nil, err
	}
	return convertToProductRes(*result), nil

}

func (p *productServer) ManageCart(ctx context.Context, req *pb.Cart) (*pb.Cart, error) {
	username, err := p.getUsernameFromToken(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "user isn't authorized")
	}
	result, err := p.productService.ManageCart(model.Cart{
		Username:  username,
		Qty:       int(req.GetQty()),
		ProductID: req.GetProductId(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Cart{
		ID:        result.ID,
		Qty:       int64(result.Qty),
		ProductId: result.ProductID,
		Product:   convertToProductRes(result.Product),
	}, nil
}
