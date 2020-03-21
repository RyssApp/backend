package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/store-service/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	s := &store.Store{Id: req.GetId()}
	err := p.db.Select(s)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Store with id `%s` does not exist.", req.GetId())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetStoreResponse{Store: s.ToProto()}, nil
}

func (p postgresStoreRepository) GetStores(ctx context.Context, req *pb.GetStoresRequest) (*pb.GetStoresResponse, error) {
	//TODO: do this
	return nil, nil
}
