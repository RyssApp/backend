package usecase

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/store-service/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type storeServiceUsecase struct {
	repo store.Repository
}

func New(repo store.Repository) *storeServiceUsecase {
	return &storeServiceUsecase{
		repo: repo,
	}
}

func (s *storeServiceUsecase) GetStore(ctx context.Context, req *types.GetStoreRequest) (*types.GetStoreResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "ID cannot be empty.")
	}
	return s.repo.GetStore(ctx, req)
}

func (s *storeServiceUsecase) GetStores(ctx context.Context, req *types.GetStoresRequest) (*types.GetStoresResponse, error) {
	return s.repo.GetStores(ctx, req)
}
