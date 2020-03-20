package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ryssapp/backend/src/go/common/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"log"
	"net"
)

// TODO: Increase this to 14 if run in production.
const HashCost int = 4

type userServiceServer struct {
}

func (s *userServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), HashCost)
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
	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{})
	grpcServer.Serve(lis)
}
