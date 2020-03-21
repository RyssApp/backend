package delivery

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/product-service/product"
)

type productServiceServer struct {
	use product.Usecase
}

func NewServer(u product.Usecase) *productServiceServer {
	return &productServiceServer{
		use: u,
	}
}

func (s *productServiceServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	res, err := s.use.GetProduct(ctx, types.GetProductRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}

func (s *productServiceServer) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	res, err := s.use.GetProducts(ctx, types.GetProductsRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}
