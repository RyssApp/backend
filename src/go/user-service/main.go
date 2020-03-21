package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ryssapp/backend/src/go/common/env"
	"github.com/ryssapp/backend/src/go/common/log"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/user-service/config"
	"github.com/ryssapp/backend/src/go/user-service/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"net"
)

type userServiceServer struct {
	hashCost int
}

func (s *userServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), s.hashCost)
	if err != nil {
		return nil, err
	}
	fmt.Println("hash generated: ", hash)

	// TODO: Check if name or email is already present

	uuid := uuid.New().String()
	user := &pb.User{Id: uuid, Email: req.GetEmail(), Username: req.GetUsername(), CreatedAt: ptypes.TimestampNow()}

	fmt.Println("saving user: ", user.String())

	// TODO: store in database

	return &pb.RegisterResponse{User: user}, nil
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

func main() {
	log.Init()
	env.Init()
	c := config.Load()

	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	rep := repository.NewPostgresRepository(db)
	zap.L().Info("Repository created", zap.Any("repository", rep))

	lis, err := net.Listen("tcp", c.Address)
	if err != nil {
		zap.L().Fatal("Failed to start tcp server.", zap.Error(err))
	}
	zap.L().Info("Serving grpc service.", zap.String("address", c.BindAddress))

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{c.Cost})
	grpcServer.Serve(lis)
}
