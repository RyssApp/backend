package delivery

import (
	"context"
	"regexp"

	"github.com/go-pg/pg/v9"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userServiceServer struct {
	hashCost int
	u        types.UserUsecase
	s        pb.SessionServiceClient
}

func NewServer(u types.UserUsecase, hashCost int, s pb.SessionServiceClient) *userServiceServer {
	return &userServiceServer{
		hashCost: hashCost,
		u:        u,
		s:        s,
	}
}

func (s *userServiceServer) Sanitize(ctx, req *pb.RegisterRequest) error {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(req.GetEmail()) {
		return status.Error(codes.InvalidArgument, "Invalid E-Mail")
	}

	if len(req.GetPassword()) < 8 {
		return status.Error(codes.InvalidArgument, "Password too short")
	}
	if len(req.GetPassword()) > 256 {
		return status.Error(codes.InvalidArgument, "Passwort too long")
	}

	if len(req.GetUsername()) < 3 {
		return status.Error(codes.InvalidArgument, "Username too short")
	}
	if len(req.GetUsername()) > 20 {
		return status.Error(codes.InvalidArgument, "Username too long")
	}

	return nil
}

func (s *userServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	if err := user.Sanitize(); err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.GetPassword()), s.hashCost)
	if err != nil {
		zap.L().Error("Failed to hash password.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	result, err := s.u.GetUser(ctx, &pb.GetUserRequest{Username: req.GetUsername()})
	if err != nil && err != pg.ErrNoRows {
		zap.L().Error("Failed to retrieve user from database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	if result != nil {
		return nil, status.Error(codes.InvalidArgument, "A user with this username already exists.")
	}

	result, err = s.u.GetUser(ctx, &pb.GetUserRequest{Email: req.GetEmail()})
	if err != nil && err != pg.ErrNoRows {
		zap.L().Error("Failed to retrieve user from database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	if result != nil {
		return nil, status.Error(codes.InvalidArgument, "A user with this email already exists.")
	}

	uuid := uuid.New().String()
	user := &types.User{Id: uuid, Email: req.GetEmail(), Username: req.GetUsername(), CreatedAt: ptypes.TimestampNow(), Password: string(hash)}

	err2 := s.u.StoreUser(ctx, user)
	if err2 != nil {
		zap.L().Error("Failed to store user in database.", zap.Error(err2))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	return &pb.RegisterResponse{User: user.UserToProto()}, nil
}

func (s *userServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	ur := &pb.GetUserRequest{Username: req.GetLogin()}
	user, err := s.u.GetUser(ctx, ur)
	if err != nil {
		zap.L().Error("Failed to retrieve user from database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	valid := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.GetPassword()))
	if valid != nil {
		return nil, status.Error(codes.Unauthenticated, "Wrong password")
	}

	session, err := s.s.Create(ctx, &pb.CreateSessionRequest{UserId: user.Id})
	if err != nil {
		return nil, err
	}
	return &pb.LoginResponse{User: user.UserToProto(), Token: session.GetToken()}, nil
}

func (s *userServiceServer) ResendEmail(ctx context.Context, reg *pb.EmailResendRequest) (*pb.EmailResendResponse, error) {
	return nil, nil
}

func (s *userServiceServer) ResetPassword(ctx context.Context, reg *pb.PasswordResetRequest) (*pb.PasswordResetResponse, error) {
	return nil, nil
}

func (s *userServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.u.GetUser(ctx, req)
	if err != nil {
		zap.L().Error("Failed to retrieve user from database.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}
	return &pb.GetUserResponse{User: user.UserToProto()}, nil
}
