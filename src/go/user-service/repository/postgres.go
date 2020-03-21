package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
)

type postgresUserRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) types.UserRepository {
	return &postgresUserRepository{
		db: db,
	}
}

func returnErrorOrValue(val *types.User, err error) (*types.User, error) {
	if err != nil {
		return nil, err
	}	
	return val, nil
}

func (p postgresUserRepository) GetUser(ctx context.Context, req *pb.GetUserRequest) (*types.User, error) {
	if req.GetId() != "" {
		u := &types.User{Id: req.GetId()}
		err := p.db.Select(u)
		return returnErrorOrValue(u, err)
	} else if req.GetUsername() != "" {
		u := &types.User{}
		err := p.db.Model(u).Where("username = ?", req.GetUsername()).Select()
		return returnErrorOrValue(u, err)
	} else if req.GetEmail() != "" {
		u := &types.User{}
		err := p.db.Model(u).Where("email = ?", req.GetEmail()).Select()
		return returnErrorOrValue(u, err)
	} else {
		return nil, nil
	}
}

func (p postgresUserRepository) StoreUser(ctx context.Context, user *types.User) error {
	return p.db.Insert(user)
}
