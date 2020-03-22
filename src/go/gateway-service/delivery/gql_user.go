package delivery

import (
	"github.com/graphql-go/graphql"
	"github.com/ryssapp/backend/src/go/common/pb"
	"github.com/ryssapp/backend/src/go/common/types"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"verified": &graphql.Field{
				Type: graphql.Boolean,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"display_name": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var loginResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "LoginResponse",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
			},
			"token": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

func (s *httpServer) registerMutation() *graphql.Field {
	return &graphql.Field{
		Type: userType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			res, err := s.UserService.Register(p.Context, &pb.RegisterRequest{
				Username: p.Args["username"].(string),
				Email:    p.Args["email"].(string),
				Password: p.Args["password"].(string),
			})
			if err != nil {
				return nil, err
			}
			if res.GetUser() != nil {
				return types.UserFromProto(res.GetUser()), nil
			}
			return nil, nil
		},
	}
}

func (s *httpServer) loginMutation() *graphql.Field {
	return &graphql.Field{
		Type: loginResponseType,
		Args: graphql.FieldConfigArgument{
			"login": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			res, err := s.UserService.Login(p.Context, &pb.LoginRequest{
				Login:    p.Args["login"].(string),
				Password: p.Args["password"].(string),
			})
			if err != nil {
				return nil, err
			}
			return types.LoginResponseFromProto(res), nil
		},
	}
}
