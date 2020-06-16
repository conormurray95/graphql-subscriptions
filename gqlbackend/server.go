package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	gogql "github.com/99designs/gqlgen/handler"

	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/graph"
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/graph/generated"
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/livenotifier"
	"github.com/conormurraypuppet/graphql-subscriptions/gqlbackend/notifier"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

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
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	done := make(chan bool)
	defer close(done)

	notifier := notifier.New(done)

	liveDone := make(chan bool)
	defer close(liveDone)

	liveNotifier := livenotifier.New(liveDone)

	srv := gogql.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Notifier: notifier, LiveNotifier: liveNotifier}}), handler.WebsocketUpgrader(websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}))

	http.Handle("/", CorsMiddleware(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", CorsMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
