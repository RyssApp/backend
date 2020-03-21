package delivery

import "github.com/graphql-go/graphql"

var locationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"zip_code": &graphql.Field{
				Type: graphql.String,
			},

			"latitude": &graphql.Field{
				Type: graphql.Float,
			},
			"longitude": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
)
