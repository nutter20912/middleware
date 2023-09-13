package main

import (
	"context"
	"fmt"
	"middleware/graph/resolver"
	"middleware/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/server"
)

const defaultPort = "8887"

func logWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {

		fmt.Println("in wapper")
		md, _ := metadata.FromContext(ctx)
		fmt.Println(md)

		err := fn(ctx, req, rsp)

		return err
	}
}

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
	r.Use(middleware.GinContextToContextMiddleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(fmt.Sprintf(":%v", defaultPort))
}
