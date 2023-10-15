package main

import (
	"context"
	"errors"
	"fmt"
	"middleware/graph/loaders"
	"middleware/graph/resolver"
	"middleware/wapper"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	microErrors "go-micro.dev/v4/errors"
)

const defaultPort = "8887"

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{}}))
	h.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		var merr *microErrors.Error
		if errors.As(e, &merr) {
			err.Message = merr.Detail
			err.Extensions = map[string]interface{}{
				"id":     merr.Id,
				"code":   merr.Code,
				"status": merr.Status,
			}
		}

		return err
	})

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
	r.Use(wapper.GinContextToContextMiddleware())
	r.Use(loaders.Middleware())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	r.Run(fmt.Sprintf(":%v", defaultPort))
}
