package types

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type Store struct {
	Id          string               `json:"id"`
	DisplayName string               `json:"display_name"`
	Logo        string               `json:"logo"`
	CreatedAt   *timestamp.Timestamp `json:"created_at"`
	*Location
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
		Location:    LocationFromProto(s.Location),
		CreatedAt:   s.GetCreatedAt(),
	}
}

type GetStoreRequest struct {
	Id string `json:"id"`
}

func GetStoreRequestFromProto(r *pb.GetStoreRequest) *GetStoreRequest {
	return &GetStoreRequest{
		Id: r.GetId(),
	}
}

type GetStoreResponse struct {
	Store *Store `json:"store"`
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
	Location    *Location   `json:"location"`
	DisplayName string      `json:"display_name"`
	Pagination  *Pagination `json:"pagination"`
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
		Location:    LocationFromProto(r.Location),
		DisplayName: r.GetDisplayName(),
		Pagination:  PaginationFromProto(r.Pagination),
	}
}

type GetStoresResponse struct {
	Stores []*Store `json:"stores"`
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

func GetStoresResponseFromProto(r *pb.GetStoresResponse) *GetStoresResponse {
	stores := make([]*Store, len(r.Stores))
	for i, st := range r.Stores {
		stores[i] = StoreFromProto(st)
	}
	return &GetStoresResponse{
		Stores: stores,
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
