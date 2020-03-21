package delivery

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/session-service/session"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type SessionServiceServer struct {
	repo       session.Repository
	secret     string
	expiration time.Duration
}

func NewServer(repo session.Repository, secret string, expiration time.Duration) *SessionServiceServer {
	return &SessionServiceServer{
		repo:       repo,
		secret:     secret,
		expiration: expiration,
	}
}

func (s *SessionServiceServer) Create(ctx context.Context, req *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(s.expiration).Unix(),
		Subject:   req.GetUserId(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(s.secret)
	if err != nil {
		zap.L().Error("failed to sign jwt token.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	return &pb.CreateSessionResponse{Token: ss}, nil
}

func (s *SessionServiceServer) Verify(ctx context.Context, req *pb.VerifySessionRequest) (*pb.VerifySessionResponse, error) {
	return nil, nil
}

func (s *SessionServiceServer) Delete(ctx context.Context, req *pb.DeleteSessionRequest) (*pb.DeleteSessionResponse, error) {
	return nil, nil
}
