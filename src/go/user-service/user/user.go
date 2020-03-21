package user

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type Repository interface {
	GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error)
}
