package product

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/types"
)

type Repository interface {
	GetProduct(ctx context.Context, req *types.GetProductRequest) (*types.GetProductResponse, error)
	GetProducts(ctx context.Context, req *types.GetProductsRequest) (*types.GetProductsResponse, error)
}

type Usecase interface {
	GetProduct(ctx context.Context, req *types.GetProductRequest) (*types.GetProductResponse, error)
	GetProducts(ctx context.Context, req *types.GetProductsRequest) (*types.GetProductsResponse, error)
}
