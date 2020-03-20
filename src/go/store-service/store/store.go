package store

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type Repository interface {
	GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error)
	GetStores(ctx context.Context, req *pb.GetStoresRequest) (*pb.GetStoresResponse, error)
}
