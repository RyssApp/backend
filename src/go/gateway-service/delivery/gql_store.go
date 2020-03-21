package delivery

import (
	"github.com/graphql-go/graphql"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var storeType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Store",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"display_name": &graphql.Field{
				Type: graphql.String,
			},
			"logo": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: locationType,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (s *httpServer) storeQuery() *graphql.Field {
	return &graphql.Field{
		Type: storeType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			if p.Args["id"] != nil {
				res, err := s.StoreService.GetStore(p.Context, &pb.GetStoreRequest{Id: p.Args["id"].(string)})
				if err != nil {
					return nil, err
				}
				if res.GetStore() == nil {
					return nil, nil
				}
				return types.StoreFromProto(res.GetStore()), nil
			}
			return nil, nil
		},
	}
}

func (s *httpServer) storesQuery() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(storeType),
		Args: graphql.FieldConfigArgument{
			"display_name": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"city": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"zip_code": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"address": &graphql.ArgumentConfig{
				Type: graphql.String,
			},
			"longitude": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"latitude": &graphql.ArgumentConfig{
				Type: graphql.Float,
			},
			"limit": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
			"offset": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &pb.GetStoresRequest{}
			if p.Args["display_name"] != nil {
				req.DisplayName = p.Args["display_name"].(string)
			}
			if p.Args["city"] != nil || p.Args["zip_code"] != nil || p.Args["address"] != nil || p.Args["longitude"] != nil || p.Args["latitude"] != nil {
				l := types.Location{}
				if p.Args["city"] != nil {
					l.City = p.Args["city"].(string)
				}
				if p.Args["zip_code"] != nil {
					l.ZipCode = p.Args["zip_code"].(string)
				}
				if p.Args["address"] != nil {
					l.Address = p.Args["address"].(string)
				}
				if p.Args["longitude"] != nil {
					lo, ok := p.Args["longitude"].(float32)
					if !ok {
						return nil, status.Error(codes.InvalidArgument, "Longitude must be a float.")
					}
					l.Longitude = lo
				}
				if p.Args["latitude"] != nil {
					la, ok := p.Args["latitude"].(float32)
					if !ok {
						return nil, status.Error(codes.InvalidArgument, "Latitude must be an float.")
					}
					l.Longitude = la
				}
				req.Location = l.ToProto()
			}
			if p.Args["limit"] != nil || p.Args["offset"] != nil {
				pa := types.Pagination{}
				if p.Args["limit"] != nil {
					l, ok := p.Args["limit"].(int32)
					if !ok {
						return nil, status.Error(codes.InvalidArgument, "Limit must be an integer.")
					}
					pa.Limit = l
				}
				if p.Args["offset"] != nil {
					o, ok := p.Args["offset"].(int32)
					if !ok {
						return nil, status.Error(codes.InvalidArgument, "Offset must be an integer.")
					}
					pa.Limit = o
				}
				req.Pagination = pa.ToProto()
			}
			res, err := s.StoreService.GetStores(p.Context, req)
			if err != nil {
				return nil, err
			}
			stores := types.GetStoresResponseFromProto(res).Stores
			return stores, nil
		},
	}
}
