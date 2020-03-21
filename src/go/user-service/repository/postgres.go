package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/user-service/user"
)

type postgresUserRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) user.Repository {
	return &postgresUserRepository{
		db: db,
	}
}

func (p postgresUserRepository) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u := &pb.User{}
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
	return &pb.GetUserResponse{User: u}, nil
}

func (p postgresUserRepository) StoreUser(ctx context.Context, user *user.User) error {
	return p.db.Insert(user)
}
