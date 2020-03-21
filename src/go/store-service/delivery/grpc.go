package delivery

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/store-service/store"
)

type storeServiceServer struct {
	use store.Usecase
}

func NewServer(u store.Usecase) *storeServiceServer {
	return &storeServiceServer{
		use: u,
	}
}

func (s *storeServiceServer) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	res, err := s.use.GetStore(ctx, store.GetStoreRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}

func (s *storeServiceServer) GetStores(ctx context.Context, req *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	res, err := s.use.GetStores(ctx, store.GetStoresRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}
