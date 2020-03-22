package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/product-service/product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postgresStoreProductRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) product.Repository {
	return &postgresProductRepository{
		db: db,
	}
}

func (p postgresProductRepository) GetProduct(ctx context.Context, req *types.GetProductRequest) (*types.GetProductResponse, error) {
	if req.Id != "" {
		pr := &types.Product{Id: req.Id}
		err := p.db.Select(pr)
		if err != nil {
			if err == pg.ErrNoRows {
				return nil, status.Errorf(codes.NotFound, "Product with id `%s` does not exist.", req.Id)
			}
			return nil, status.Error(codes.Internal, err.Error())
		}
		return &types.GetProductResponse{Product: pr}, nil
	}
	pr := new(types.Product)
	err := p.db.Model(pr).Where("barcode = ?", req.Barcode).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "Product with barcode `%s` does not exist.", req.Barcode)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetProductResponse{Product: pr}, nil
}

func (p postgresProductRepository) GetProducts(ctx context.Context, req *types.GetProductsRequest) (*types.GetProductsResponse, error) {
	var products []*types.Product
	q := p.db.Model(&products)
	if req.DisplayName != "" {
		q.Where("display_name = ?", req.DisplayName)
	}
	if req.Pagination != nil {
		if req.Pagination.Limit != 0 {
			q.Limit(int(req.Pagination.Limit))
		}
		if req.Pagination.Offset != 0 {
			q.Offset(int(req.Pagination.Offset))
		}
	}
	err := q.Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "No products with that requirements found.")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetProductsResponse{
		Products: products,
	}, nil
}
