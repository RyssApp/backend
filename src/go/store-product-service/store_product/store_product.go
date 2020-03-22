package store_product

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/types"
)

type Repository interface {
	GetStoreProduct(ctx context.Context, req *types.GetStoreProductRequest) (*types.GetStoreProductResponse, error)
	GetStoreProducts(ctx context.Context, req *types.GetStoreProductsRequest) (*types.GetStoreProductsResponse, error)
}

type Usecase interface {
	GetStoreProduct(ctx context.Context, req *types.GetStoreProductRequest) (*types.GetStoreProductResponse, error)
	GetStoreProducts(ctx context.Context, req *types.GetStoreProductsRequest) (*types.GetStoreProductsResponse, error)
}
