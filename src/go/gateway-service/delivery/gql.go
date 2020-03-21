package delivery

import (
	"github.com/graphql-go/graphql"
)

func (s *httpServer) createQuery() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"store": s.storeQuery(),
				"stores": s.storesQuery(),
			},
		},
	)
}
