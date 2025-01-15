package main

import (
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/adasarpan404/gopostgrespoc/config"
	"github.com/adasarpan404/gopostgrespoc/graph"
	"github.com/adasarpan404/gopostgrespoc/graph/generated"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := config.InitDB()

	r := gin.Default()

	resolver := &graph.Resolver{DB: db}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	r.POST("/query", gin.WrapH(srv))
	r.GET("/", gin.WrapH(playground.Handler("GraphQL Playground", "/query")))

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(r.Run(":8080"))
}
