package types

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type StoreProduct struct {
	Id        string               `json:"id"`
	StoreId   string               `json:"store_id"`
	ProductId string               `json:"product_id"`
	Quantity  int32                `json:"quantity"`
	CreatedAt *timestamp.Timestamp `json:"created_at"`
}

func (p *StoreProduct) ToProto() *pb.StoreProduct {
	return &pb.StoreProduct{
		Id:        p.Id,
		StoreId:   p.StoreId,
		ProductId: p.ProductId,
		Quantity:  pb.StoreProduct_Quantity(p.Quantity),
		CreatedAt: p.CreatedAt,
	}
}

func StoreProductFromProto(p *pb.StoreProduct) *StoreProduct {
	return &StoreProduct{
		Id:        p.GetId(),
		StoreId:   p.GetStoreId(),
		ProductId: p.GetProductId(),
		Quantity:  int32(p.Quantity),
		CreatedAt: p.GetCreatedAt(),
	}
}

type GetStoreProductRequest struct {
	Id string `json:"id"`
}

func GetStoreProductRequestFromProto(r *pb.GetStoreProductRequest) *GetStoreProductRequest {
	return &GetStoreProductRequest{
		Id: r.GetId(),
	}
}

type GetStoreProductResponse struct {
	StoreProduct *StoreProduct `json:"store_product"`
}

func (r *GetStoreProductResponse) ToProto() *pb.GetStoreProductResponse {
	return &pb.GetStoreProductResponse{
		StoreProduct: r.StoreProduct.ToProto(),
	}
}

func GetStoreProductResponseFromProto(r *pb.GetStoreProductResponse) *GetStoreProductResponse {
	return &GetStoreProductResponse{
		StoreProduct: StoreProductFromProto(r.GetStoreProduct()),
	}
}

type GetStoreProductsRequest struct {
	StoreId    string      `json:"store_id"`
	ProductId  string      `json:"product_id"`
	Quantity   int32       `json:"quantity"`
	Pagination *Pagination `json:"pagination"`
}

func (r *GetStoreProductsRequest) ToProto() *pb.GetStoreProductsRequest {
	return &pb.GetStoreProductsRequest{
		StoreId:    r.StoreId,
		ProductId:  r.ProductId,
		Quantity:   pb.StoreProduct_Quantity(r.Quantity),
		Pagination: r.Pagination.ToProto(),
	}
}

func GetStoreProductsRequestFromProto(r *pb.GetStoreProductsRequest) *GetStoreProductsRequest {
	return &GetStoreProductsRequest{
		StoreId:    r.StoreId,
		ProductId:  r.ProductId,
		Quantity:   int32(r.Quantity),
		Pagination: PaginationFromProto(r.Pagination),
	}
}

type GetStoreProductsResponse struct {
	StoreProducts []*StoreProduct `json:"store_products"`
}

func (r *GetStoreProductsResponse) ToProto() *pb.GetStoreProductsResponse {
	pbStoreProducts := make([]*pb.StoreProduct, len(r.StoreProducts))
	for i, sp := range r.StoreProducts {
		pbStoreProducts[i] = sp.ToProto()
	}
	return &pb.GetStoreProductsResponse{
		StoreProducts: pbStoreProducts,
	}
}

func GetStoreProductsResponseFromProto(r *pb.GetStoreProductsResponse) *GetStoreProductsResponse {
	storeProducts := make([]*StoreProduct, len(r.StoreProducts))
	for i, sp := range r.StoreProducts {
		storeProducts[i] = StoreProductFromProto(sp)
	}
	return &GetStoreProductsResponse{
		StoreProducts: storeProducts,
	}
}
