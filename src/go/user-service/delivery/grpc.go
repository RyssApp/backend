package delivery

import (
	"context"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/user-service/user"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServiceServer struct {
	hashCost int
	u        user.Usecase
}

func NewServer(u user.Usecase, hashCost int) *userServiceServer {
	return &userServiceServer{
		hashCost: hashCost,
		u:        u,
	}
}

func (s *userServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), s.hashCost)
	if err != nil {
		zap.L().Error("Failed to hash password.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	result, err := s.u.GetUser(ctx, &pb.GetUserRequest{Username: req.GetUsername()})
	if err != nil {
		zap.L().Error("Failed to retrieve user from database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	if result != nil {
		return nil, status.Error(codes.InvalidArgument, "A user with this username already exists.")
	}

	result, err = s.u.GetUser(ctx, &pb.GetUserRequest{Email: req.GetEmail()})
	if result != nil {
		return nil, status.Error(codes.InvalidArgument, "A user with this email already exists.")
	}

	if err != nil {
		zap.L().Error("Failed to retrieve user from database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	uuid := uuid.New().String()
	user := &user.User{Id: uuid, Email: req.GetEmail(), Username: req.GetUsername(), CreatedAt: ptypes.TimestampNow(), Password: string(hash)}

	err = s.u.StoreUser(ctx, user)
	if err != nil {
		zap.L().Error("Failed to store user in database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	return &pb.RegisterResponse{User: user.ToProto()}, nil
}

func (s *userServiceServer) Login(ctx context.Context, reg *pb.LoginRequest) (*pb.LoginResponse, error) {
	return nil, nil
}

func (s *userServiceServer) ResendEmail(ctx context.Context, reg *pb.EmailResendRequest) (*pb.EmailResendResponse, error) {
	return nil, nil
}

func (s *userServiceServer) ResetPassword(ctx context.Context, reg *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	return nil, nil
}

func (s *userServiceServer) GetUser(ctx context.Context, reg *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return nil, nil
}
