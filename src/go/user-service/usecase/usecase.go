package usecase

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/user-service/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServiceUsecase struct {
	repo user.Repository
}

func New(repo user.Repository) *userServiceUsecase {
	return &userServiceUsecase{
		repo: repo,
	}
}

func (s *userServiceUsecase) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if req.GetId() == "" && req.GetEmail() == "" && req.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, "You have to provide an item to search for.")
	}
	return s.repo.GetUser(ctx, req)
}

func (s *userServiceUsecase) StoreUser(ctx context.Context, user *user.User) error {
	return s.repo.StoreUser(ctx, user)
}
