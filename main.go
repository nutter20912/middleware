package main

import (
	"fmt"
	"middleware/graph/loaders"
	"middleware/graph/resolver"
	"middleware/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8887"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{}}))

	return gin.WrapH(h)
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return gin.WrapH(h)
}

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(config))
	r.Use(middleware.GinContextToContextMiddleware())
	r.Use(loaders.Middleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(fmt.Sprintf(":%v", defaultPort))
}
