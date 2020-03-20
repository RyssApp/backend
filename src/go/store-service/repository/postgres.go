package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/store-service/store"
)

type postgresStoreRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) store.Repository {
	return &postgresStoreRepository{
		db: db,
	}
}

func (p postgresStoreRepository) GetStore(ctx context.Context, req *pb.GetStoreRequest) (*pb.GetStoreResponse, error) {
	s := &pb.Store{Id: req.GetId()}
	err := p.db.Select(s)
	if err != nil {
		return nil, err
	}
	return &pb.GetStoreResponse{Store: s}, nil
}

func (p postgresStoreRepository) GetStores(ctx context.Context, req *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	//TODO: do this
	return nil, nil
}
