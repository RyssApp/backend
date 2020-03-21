package store

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
)

type Store struct {
	Id          string
	DisplayName string
	Logo        string
	CreatedAt   *timestamp.Timestamp
	*types.Location
}

func (s *Store) ToProto() *pb.Store {
	return &pb.Store{
		Id:          s.Id,
		DisplayName: s.DisplayName,
		Logo:        s.Logo,
		Location:    s.Location.ToProto(),
		CreatedAt:   s.CreatedAt,
	}
}

func StoreFromProto(s *pb.Store) *Store {
	return &Store{
		Id:          s.GetId(),
		DisplayName: s.GetDisplayName(),
		Logo:        s.GetLogo(),
		Location:    types.LocationFromProto(s.Location),
		CreatedAt:   s.GetCreatedAt(),
	}
}

type GetStoreRequest struct {
	Id string
}

func GetStoreRequestFromProto(r *pb.GetStoreRequest) *GetStoreRequest {
	return &GetStoreRequest{
		Id: r.GetId(),
	}
}

type GetStoreResponse struct {
	Store *Store
}

func (r *GetStoreResponse) ToProto() *pb.GetStoreResponse {
	return &pb.GetStoreResponse{
		Store: r.Store.ToProto(),
	}
}

func GetStoreResponseFromProto(r *pb.GetStoreResponse) *GetStoreResponse {
	return &GetStoreResponse{
		Store: StoreFromProto(r.Store),
	}
}

type GetStoresRequest struct {
	Location    *types.Location
	DisplayName string
	Pagination  *types.Pagination
}

func (r *GetStoresRequest) ToProto() *pb.GetStoresRequest {
	return &pb.GetStoresRequest{
		Location:    r.Location.ToProto(),
		DisplayName: r.DisplayName,
		Pagination:  r.Pagination.ToProto(),
	}
}

func GetStoresRequestFromProto(r *pb.GetStoresRequest) *GetStoresRequest {
	return &GetStoresRequest{
		Location:    types.LocationFromProto(r.Location),
		DisplayName: r.GetDisplayName(),
		Pagination:  types.PaginationFromProto(r.Pagination),
	}
}

type GetStoresResponse struct {
	Stores []*Store
}

func (r *GetStoresResponse) ToProto() *pb.GetStoresResponse {
	pbStores := make([]*pb.Store, len(r.Stores))
	for i, st := range r.Stores {
		pbStores[i] = st.ToProto()
	}
	return &pb.GetStoresResponse{
		Stores: pbStores,
	}
}

type Repository interface {
	GetStore(ctx context.Context, req *GetStoreRequest) (*GetStoreResponse, error)
	GetStores(ctx context.Context, req *GetStoresRequest) (*GetStoresResponse, error)
}

type Usecase interface {
	GetStore(ctx context.Context, req *GetStoreRequest) (*GetStoreResponse, error)
	GetStores(ctx context.Context, req *GetStoresRequest) (*GetStoresResponse, error)
}
