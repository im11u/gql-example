package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/im11u/gql-example/go/graph"
	"github.com/im11u/gql-example/go/graph/generated"
)

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
