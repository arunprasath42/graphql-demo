package main

import (
	"log"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/arunprasath42/graphql-live/graph"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
	r.GET("/playground", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
	r.POST("/query", gin.WrapH(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(r.Run(":" + port))
}
