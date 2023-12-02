package server

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/im11u/gql-example/go/graph"
	"github.com/im11u/gql-example/go/graph/generated"
)

type GinServer struct {
	config *config
}

func (g *GinServer) Run() {
	r := g.setupRouter()
	r.Run(g.config.addr)
}

func (g *GinServer) graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func (g *GinServer) playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func (g *GinServer) setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/query", g.graphqlHandler())
	r.GET("/", g.playgroundHandler())

	return r
}
