package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/repath-io/auth"
	"log"
	"net/http"
	"os"
	"repath.io/config"
	"repath.io/data"
	"repath.io/graph"
	"repath.io/graph/generated"

	// "repath.io/data"
)

const defaultPort = "8081"

func main() {
	configuration := config.LoadConfiguration()
	fmt.Printf("starting %s service\n", configuration.ServiceName)
	data.Initialize(configuration.Neo4j.URI, configuration.Neo4j.Username, configuration.Neo4j.Password);

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	if configuration.DevelopmentMode != "True" {
		router.Use(auth.AuthorizationProvider())
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", playground.Handler("GraphQL playground", "query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
