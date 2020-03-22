package types

import (
	"context"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
)

// User represents a ryss user.
type User struct {
	Id          string
	Email       string `pg:",unique"`
	Verified    bool
	Username    string `pg:",unique"`
	DisplayName string
	Password    string
	CreatedAt   *timestamp.Timestamp
}

// ToProto converts a protobuf user struct to an independent type definition.
func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:          u.Id,
		Email:       u.Email,
		Verified:    u.Verified,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		CreatedAt:   u.CreatedAt,
	}
}

// Converts a protobuf message, reprenting a
// ryss user, to a User struct.
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

type LoginResponse struct {
	User  *User  `json:"user"`
	Token string `json:"token"`
}

func (r *LoginResponse) LoginResponseToProto() *pb.LoginResponse {
	return &pb.LoginResponse{
		User:  r.User.ToProto(),
		Token: r.Token,
	}
}

func LoginResponseFromProto(r *pb.LoginResponse) *LoginResponse {
	return &LoginResponse{
		User:  UserFromProto(r.User),
		Token: r.GetToken(),
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
