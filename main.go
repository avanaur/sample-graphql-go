package main

import (
	"github.com/graphql-go/graphql"
	handler "github.com/graphql-go/graphql-go-handler"
	"math/rand"
	"net/http"
)

var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"latestPost": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "Hello World!", nil
			},
		},
		"randomNum": &graphql.Field{
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return rand.Intn(1000), nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})


func main() {

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := handler.New(&handler.Config{
		Schema: &Schema,
		Pretty: true,
	})

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)
}


// curl -XPOST http://localhost:8080/graphql \
//-H 'Content-Type: application/graphql' \
//-d 'query Root{ randomNum }'
//{
//        "data": {
//                "randomNum": 456
//        }
//}%                                                                                                                                                                          tyronevillaluna@Tyrones-Air sample-grapql-go % curl -XPOST http://localhost:8080/graphql \
// curl -XPOST http://localhost:8080/graphql \
//-H 'Content-Type: application/graphql' \
//-d 'query Root{ randomNum, latestPost }'
//{
//        "data": {
//                "latestPost": "Hello World!",
//                "randomNum": 511
//        }
//}%