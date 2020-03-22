package delivery

import (
	"context"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/store-product-service/store_product"
)

type storeProductServiceServer struct {
	use store_product.Usecase
}

func NewServer(u store_product.Usecase) *storeProductServiceServer {
	return &storeProductServiceServer{
		use: u,
	}
}

func (s *storeProductServiceServer) GetStoreProduct(ctx context.Context, req *pb.GetStoreProductRequest) (*pb.GetStoreProductResponse, error) {
	res, err := s.use.GetStoreProduct(ctx, types.GetStoreProductRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}

func (s *storeProductServiceServer) GetStoreProducts(ctx context.Context, req *pb.GetStoreProductsRequest) (*pb.GetStoreProductsResponse, error) {
	res, err := s.use.GetStoreProducts(ctx, types.GetStoreProductsRequestFromProto(req))
	if err != nil {
		return nil, err
	}
	return res.ToProto(), nil
}
