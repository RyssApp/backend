package usecase

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/store-product-service/store_product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type storeProductServiceUsecase struct {
	repo store_product.Repository
}

func New(repo store_product.Repository) *storeProductServiceUsecase {
	return &storeProductServiceUsecase{
		repo: repo,
	}
}

func (s *storeProductServiceUsecase) GetProduct(ctx context.Context, req *types.GetStoreProductRequest) (*types.GetStoreProductResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID cannot be empty.")
	}
	return s.repo.GetStoreProduct(ctx, req)
}

func (s *storeProductServiceUsecase) GetStoreProducts(ctx context.Context, req *types.GetStoreProductsRequest) (*types.GetStoreProductsResponse, error) {
	return s.repo.GetStoreProducts(ctx, req)
}
