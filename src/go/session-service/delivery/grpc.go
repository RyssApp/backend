package delivery

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/session-service/session"
)

type SessionServiceServer struct {
	repo session.Repository
}

func NewServer(repo session.Repository) *SessionServiceServer {
	return &SessionServiceServer{
		repo: repo,
	}
}

func (s *SessionServiceServer) Create(ctx context.Context, req *pb.CreateSessionRequest) (*pb.CreateSessionResponse, error) {
	return nil, nil
}

func (s *SessionServiceServer) Verify(ctx context.Context, req *pb.VerifySessionRequest) (*pb.VerifySessionResponse, error) {
	return nil, nil
}

func (s *SessionServiceServer) Delete(ctx context.Context, req *pb.DeleteSessionRequest) (*pb.DeleteSessionResponse, error) {
	return nil, nil
}
