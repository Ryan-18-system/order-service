package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Ryan-18-system/order-service/internal/order/usecase"
	"github.com/graphql-go/graphql"
)

func GraphQLHandler(uc usecase.OrderUseCase) http.HandlerFunc {
	fields := graphql.Fields{
		"listOrders": &graphql.Field{
			Type: graphql.NewList(orderType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return uc.ListOrders(p.Context)
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ := graphql.NewSchema(schemaConfig)

	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query string `json:"query"`
		}
		_ = json.NewDecoder(r.Body).Decode(&params)

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})

		json.NewEncoder(w).Encode(result)
	}
}

var orderType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Order",
	Fields: graphql.Fields{
		"id":            &graphql.Field{Type: graphql.Int},
		"customer_name": &graphql.Field{Type: graphql.String},
		"total":         &graphql.Field{Type: graphql.Float},
		"created_at":    &graphql.Field{Type: graphql.String},
	},
})
