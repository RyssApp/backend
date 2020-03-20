package main

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/store-service/store"
)

type storeServiceServer struct {
	repo store.Repository
}

func NewServer(r store.Repository) *storeServiceServer {
	return &storeServiceServer{
		repo: r,
	}
}

func (s storeServiceServer) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	return s.repo.GetStore(ctx, req)
}

func (s storeServiceServer) GetStores(context.Context, *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	return nil, nil
}
