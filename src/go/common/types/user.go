package types

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type User struct {
	Id          string
	Email       string `pg:",unique"`
	Verified    bool
	Username    string `pg:",unique"`
	DisplayName string
	Password    string
	CreatedAt   *timestamp.Timestamp
}

func (u *User) UserToProto() *pb.User {
	return &pb.User{
		Id:          u.Id,
		Email:       u.Email,
		Verified:    u.Verified,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		CreatedAt:   u.CreatedAt,
	}
}

func UserFromProto(u *pb.User) *User {
	return &User{
		Id:          u.Id,
		Email:       u.Email,
		Verified:    u.Verified,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		CreatedAt:   u.CreatedAt,
	}
}

type UserRepository interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*User, error)
	StoreUser(ctx context.Context, user *User) error
}

type UserUsecase interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*User, error)
	StoreUser(ctx context.Context, user *User) error
}
