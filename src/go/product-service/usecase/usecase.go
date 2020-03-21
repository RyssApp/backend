package usecase

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/product-service/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type productServiceUsecase struct {
	repo product.Repository
}

func New(repo product.Repository) *productServiceUsecase {
	return &productServiceUsecase{
		repo: repo,
	}
}

func (s *productServiceUsecase) GetProduct(ctx context.Context, req *types.GetProductRequest) (*types.GetProductResponse, error) {
	if req.Id == "" && req.Barcode == "" {
		return nil, status.Error(codes.InvalidArgument, "ID and barcode cannot be empty.")
	}
	return s.repo.GetProduct(ctx, req)
}

func (s *productServiceUsecase) GetProducts(ctx context.Context, req *types.GetProductsRequest) (*types.GetProductsResponse, error) {
	return s.repo.GetProducts(ctx, req)
}
