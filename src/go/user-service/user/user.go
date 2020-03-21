package user

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
)

type User struct {
	Id string
	Email string
	Verified bool
	Username string
	DisplayName string
	Password string
	CreatedAt *timestamp.Timestamp
}

func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id: u.Id,
		Email: u.Email,
		Verified: u.Verified,
		Username: u.Username,
		DisplayName: u.DisplayName,
		CreatedAt: u.CreatedAt,
	}
}

func UserFromProto(u *pb.User) *User {
	return &User{
		Id: u.Id,
		Email: u.Email,
		Verified: u.Verified,
		Username: u.Username,
		DisplayName: u.DisplayName,
		CreatedAt: u.CreatedAt,
	}
}

type Repository interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	StoreUser(ctx context.Context, user *User) error
}

type Usecase interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
	StoreUser(ctx context.Context, user *User) error
}
