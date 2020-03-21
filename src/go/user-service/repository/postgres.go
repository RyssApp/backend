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

func (p postgresUserRepository) GetUser(ctx context.Context, req *pb.GetUserRequest) (*types.User, error) {
	u := &types.User{}
	if req.GetId() != "" {
		u.Id = req.GetId()
	}
	if req.GetEmail() != "" {
		u.Email = req.GetEmail()
	}
	if req.GetUsername() != "" {
		u.Username = req.GetUsername()
	}

	err := p.db.Select(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (p postgresUserRepository) StoreUser(ctx context.Context, user *types.User) error {
	return p.db.Insert(user)
}
