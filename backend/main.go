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

func main() {
	schema := graphql.MustParseSchema(Schema, &query{})
	http.Handle("/query", &relay.Handler{Schema: schema})
	// init model
	vegetables = map[string]Vegetable{
		"tomato": Vegetable{name: "Tomato", price: 100, image: strPtr("https://picsum.photos/id/152/100/100")},
		"potato": Vegetable{name: "Potato", price: 50, image: strPtr("https://picsum.photos/id/159/100/100")},
		"corn":   Vegetable{name: "Corn", price: 200},
	}
	graphiqlHandler, err := graphiql.NewGraphiqlHandler("/query")
	if err != nil {
		panic(err)
	}
	http.Handle("/", graphiqlHandler)
	// Run
	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
