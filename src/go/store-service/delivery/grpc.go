package delivery

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
)

type storeServiceServer struct {
	use types.Usecase
}

func NewServer(u types.Usecase) *storeServiceServer {
	return &storeServiceServer{
		use: u,
	}
}

func (s *storeServiceServer) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	res, err := s.use.GetStore(ctx, types.GetStoreRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}

func (s *storeServiceServer) GetStores(ctx context.Context, req *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	res, err := s.use.GetStores(ctx, types.GetStoresRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}
