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
	val, _ := s.repo.GetToken(req.GetUserId())
	if val != "" {
		return &pb.CreateSessionResponse{Token: val}, nil
	}

	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(s.expiration).Unix(),
		Subject:   req.GetUserId(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(s.secret))
	if err != nil {
		zap.L().Error("failed to sign jwt token.", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	err = s.repo.SetToken(req.GetUserId(), ss)
	if err != nil {
		zap.L().Error("failed to set token", zap.Error(err))
		return nil, status.Error(codes.Internal, "Internal server error occured")
	}

	return &pb.CreateSessionResponse{Token: ss}, nil
}

func (s *SessionServiceServer) Verify(ctx context.Context, req *pb.VerifySessionRequest) (*pb.VerifySessionResponse, error) {
	val, err := validateToken(req, s.secret)
	return val, err.Err()
}

func (s *SessionServiceServer) Delete(ctx context.Context, req *pb.DeleteSessionRequest) (*pb.DeleteSessionResponse, error) {
	_, err := validateToken(&pb.VerifySessionRequest{Token: req.GetToken(), UserId: req.GetUserId()}, s.secret)
	if err != nil && err.Code() != codes.Unauthenticated {
		return nil, err.Err()
	}

	err2 := s.repo.DelToken(req.GetUserId())
	if err2 != nil {
		zap.L().Error("Failed to delete token", zap.Error(err2))
		return nil, status.Error(codes.Internal, "Internal server error. Token was not deleted.")
	}
	return &pb.DeleteSessionResponse{}, nil
}

func validateToken(req *pb.VerifySessionRequest, secret string) (*pb.VerifySessionResponse, *status.Status) {
	token, err := jwt.ParseWithClaims(req.GetToken(), &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Error(codes.InvalidArgument, "Expected HMAC signing method.")
		}
		return []byte(secret), nil
	})

	if token != nil && token.Valid {
		if token.Claims.(*jwt.StandardClaims).Subject == req.GetUserId() {
			return &pb.VerifySessionResponse{}, nil
		} else {
			return nil, status.New(codes.InvalidArgument, "User Ids doesn't match.")
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, status.New(codes.InvalidArgument, "Invalid token provided")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			err := s.repo.DelToken(req.GetUserId())
			if err != nil {
				zap.L().Error("Failed to delete token", zap.Error(err))
			}
			return nil, status.New(codes.Unauthenticated, "Token is expired.")
		} else {
			zap.L().Warn("JWT token couldn't be parsed.", zap.Error(err))
			return nil, status.New(codes.InvalidArgument, "JWT Token could not be parsed.")
		}
	} else {
		zap.L().Warn("JWT token couldn't be parsed.", zap.Error(err))
		return nil, status.New(codes.InvalidArgument, "JWT Token could not be parsed.")
	}
}
