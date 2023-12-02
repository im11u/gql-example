package server

import (
	"database/sql"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"

	"github.com/im11u/gql-example/go/infrastructure/database"
	"github.com/im11u/gql-example/go/infrastructure/graph"
	"github.com/im11u/gql-example/go/infrastructure/graph/gen"
)

// Ginを使用するサーバー
type GinServer struct {
	config *config
	db     *sql.DB
}

func NewGinServer(config *config) *GinServer {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	return &GinServer{
		config: config,
		db:     db,
	}
}

func (g *GinServer) Run() {
	defer g.db.Close()
	r := g.setupRouter()
	r.Run(g.config.addr)
}

func (g *GinServer) graphqlHandler() gin.HandlerFunc {
	r := graph.NewResolver(g.db)
	h := handler.NewDefaultServer(gen.NewExecutableSchema(gen.Config{Resolvers: r}))

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
