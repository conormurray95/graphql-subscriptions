package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/friendsofgo/graphiql"
	graphql "github.com/graph-gophers/graphql-go"

	"github.com/graph-gophers/graphql-go/relay"
)

// Data Structure
type Vegetable struct {
	name  string
	price int
	image *string
}

var vegetables map[string]Vegetable

// Utils
func strPtr(str string) *string {
	return &str
}

// Resolver
func (q *query) Vegetable(ctx context.Context, args struct{ Name string }) *VegetableResolver {
	v, ok := vegetables[strings.ToLower(args.Name)]
	if ok {
		return &VegetableResolver{v: &v}
	}
	return nil
}

type VegetableResolver struct {
	v *Vegetable
}

func (r *VegetableResolver) Name() string   { return r.v.name }
func (r *VegetableResolver) Price() int32   { return int32(r.v.price) }
func (r *VegetableResolver) Image() *string { return r.v.image }

type query struct{}

const Schema = `
type Vegetable {
    name: String!
    price: Int!
    image: String
}
type Query {
    vegetable(name: String!): Vegetable
}
schema {
    query: Query
}
`

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// allow cross domain AJAX requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Allow-Wildcard", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	schema := graphql.MustParseSchema(Schema, &query{})
	http.Handle("/graphql", CorsMiddleware(&relay.Handler{Schema: schema}))
	// init model
	vegetables = map[string]Vegetable{
		"tomato": Vegetable{name: "Tomato", price: 100, image: strPtr("https://picsum.photos/id/152/100/100")},
		"potato": Vegetable{name: "Potato", price: 50, image: strPtr("https://picsum.photos/id/159/100/100")},
		"corn":   Vegetable{name: "Corn", price: 200},
	}
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/graphql")
	if err != nil {
		panic(err)
	}
	http.Handle("/", CorsMiddleware(graphiqlHandler))
	// Run
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
