package repository

import (
	"context"
	"github.com/go-pg/pg/v9"
	"github.com/ryssapp/backend/src/go/common/types"
	"github.com/ryssapp/backend/src/go/store-product-service/store_product"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type postgresStoreProductRepository struct {
	db *pg.DB
}

func NewPostgresRepository(db *pg.DB) store_product.Repository {
	return &postgresStoreProductRepository{
		db: db,
	}
}

func (p postgresStoreProductRepository) GetStoreProduct(ctx context.Context, req *types.GetStoreProductRequest) (*types.GetStoreProductResponse, error) {
	pr := &types.StoreProduct{Id: req.Id}
	err := p.db.Select(pr)
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "StoreProduct with id `%s` does not exist.", req.Id)
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetStoreProductResponse{StoreProduct: pr}, nil
}

func (p postgresStoreProductRepository) GetStoreProducts(ctx context.Context, req *types.GetStoreProductsRequest) (*types.GetStoreProductsResponse, error) {
	var storeProducts []*types.StoreProduct
	q := p.db.Model(&storeProducts)
	if req.StoreId != "" {
		q.Where("store_id = ?", req.StoreId)
	}
	if req.ProductId != "" {
		q.Where("product_id = ?", req.ProductId)
	}
	if req.Quantity != 0 {
		q.Where("quantity = ?", req.Quantity)
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
			return nil, status.Errorf(codes.NotFound, "No store products with that requirements found.")
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.GetStoreProductsResponse{
		StoreProducts: storeProducts,
	}, nil
}
