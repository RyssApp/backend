package types

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/ryssapp/backend/src/go/common/pb"
)

type Product struct {
	Id          string               `json:"id"`
	DisplayName string               `json:"display_name"`
	Barcode     string               `json:"barcode"`
	CreatedAt   *timestamp.Timestamp `json:"created_at"`
}

func (p *Product) ToProto() *pb.Product {
	return &pb.Product{
		Id:          p.Id,
		DisplayName: p.DisplayName,
		Barcode:     p.Barcode,
		CreatedAt:   p.CreatedAt,
	}
}

func ProductFromProto(p *pb.Product) *Product {
	return &Product{
		Id:          p.GetId(),
		DisplayName: p.GetDisplayName(),
		Barcode:     p.Barcode,
		CreatedAt:   p.GetCreatedAt(),
	}
}

type GetProductRequest struct {
	Id      string `json:"id"`
	Barcode string `json:"barcode"`
}

func GetProductRequestFromProto(r *pb.GetProductRequest) *GetProductRequest {
	return &GetProductRequest{
		Id:      r.GetId(),
		Barcode: r.GetBarcode(),
	}
}

type GetProductResponse struct {
	Product *Product `json:"product"`
}

func (r *GetProductResponse) ToProto() *pb.GetProductResponse {
	return &pb.GetProductResponse{
		Product: r.Product.ToProto(),
	}
}

func GetProductResponseFromProto(r *pb.GetProductResponse) *GetProductResponse {
	return &GetProductResponse{
		Product: ProductFromProto(r.Product),
	}
}

type GetProductsRequest struct {
	DisplayName string      `json:"display_name"`
	Pagination  *Pagination `json:"pagination"`
}

func (r *GetProductsRequest) ToProto() *pb.GetProductsRequest {
	return &pb.GetProductsRequest{
		DisplayName: r.DisplayName,
		Pagination:  r.Pagination.ToProto(),
	}
}

func GetProductsRequestFromProto(r *pb.GetProductsRequest) *GetProductsRequest {
	return &GetProductsRequest{
		DisplayName: r.GetDisplayName(),
		Pagination:  PaginationFromProto(r.Pagination),
	}
}

type GetProductsResponse struct {
	Products []*Product `json:"products"`
}

func (r *GetProductsResponse) ToProto() *pb.GetProductsResponse {
	pbProducts := make([]*pb.Product, len(r.Products))
	for i, pr := range r.Products {
		pbProducts[i] = pr.ToProto()
	}
	return &pb.GetProductsResponse{
		Products: pbProducts,
	}
}

func GetProductsResponseFromProto(r *pb.GetProductsResponse) *GetProductsResponse {
	products := make([]*Product, len(r.Products))
	for i, pr := range r.Products {
		products[i] = ProductFromProto(pr)
	}
	return &GetProductsResponse{
		Products: products,
	}
}
