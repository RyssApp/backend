package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/types"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postgresStoreRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) types.Repository {
	return &postgresStoreRepository{
		db: db,
	}
}

func (p postgresStoreRepository) GetStore(ctx context.Context, req *types.GetStoreRequest) (*types.GetStoreResponse, error) {
	s := &types.Store{Id: req.Id}
	err := p.db.Select(s)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Store with id `%s` does not exist.", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetStoreResponse{Store: s}, nil
}

func (p postgresStoreRepository) GetStores(ctx context.Context, req *types.GetStoresRequest) (*types.GetStoresResponse, error) {
	zap.L().Info("dasds", zap.Any("accc", req))
	var stores []*types.Store
	q := p.db.Model(&stores)
	if req.Location != nil {
		if req.Location.Latitude != 0 {
			q.Where("latitude = ?", req.Location.Latitude)
		}
		if req.Location.Longitude != 0 {
			q.Where("longitude = ?", req.Location.Longitude)
		}
		if req.Location.Address != "" {
			q.Where("address = ?", req.Location.Address)
		}
		if req.Location.ZipCode != "" {
			q.Where("zip_code = ?", req.Location.ZipCode)
		}
		if req.Location.City != "" {
			q.Where("city = ?", req.Location.City)
		}
	}
	if req.DisplayName != "" {
		q.Where("display_name = ?", req.DisplayName)
	}
	if req.Pagination != nil {
		if req.Pagination.Limit != 0 {
			q.Limit(int(req.Pagination.Limit))
		}
		if req.Pagination.Offset != 0 {
			q.Offset(int(req.Pagination.Offset))
		}
	}
	err := q.Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "No stores with that requirements found.")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetStoresResponse{
		Stores: stores,
	}, nil
}
