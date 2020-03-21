package store

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type Store struct {
	Id          string
	DisplayName string
	Logo        string
	Location    *pb.Location
	CreatedAt   *timestamp.Timestamp
}

func (s *Store) ToProto() *pb.Store {
	return &pb.Store{
		Id:          s.Id,
		DisplayName: s.DisplayName,
		Logo:        s.Logo,
		Location:    s.Location,
		CreatedAt:   s.CreatedAt,
	}
}

type Repository interface {
	GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error)
	GetStores(ctx context.Context, req *pb.GetStoresRequest) (*pb.GetStoresResponse, error)
}
