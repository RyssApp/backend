package store

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/types"
)

type Repository interface {
	GetStore(ctx context.Context, req *types.GetStoreRequest) (*types.GetStoreResponse, error)
	GetStores(ctx context.Context, req *types.GetStoresRequest) (*types.GetStoresResponse, error)
}

type Usecase interface {
	GetStore(ctx context.Context, req *types.GetStoreRequest) (*types.GetStoreResponse, error)
	GetStores(ctx context.Context, req *types.GetStoresRequest) (*types.GetStoresResponse, error)
}