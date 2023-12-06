package main

import (
	"context"
	"errors"
	"fmt"
	_ "middleware/config"
	"middleware/graph/extensions"
	"middleware/graph/loaders"
	"middleware/graph/resolver"
	"middleware/otel"
	"middleware/wrapper"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	microErrors "go-micro.dev/v4/errors"
	"go-micro.dev/v4/metadata"
)

func graphqlHandler() gin.HandlerFunc {
	srv := handler.New(resolver.NewExecutableSchema(resolver.Config{Resolvers: &resolver.Resolver{}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: func(ctx context.Context, initPayload transport.InitPayload) (context.Context, error) {
			ctx = metadata.Set(ctx, "Authorization", initPayload.Authorization())

			return ctx, nil
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.SetQueryCache(lru.New(1000))
	srv.Use(extensions.NewTracer())
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{Cache: lru.New(100)})

	srv.AddTransport(transport.SSE{})

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
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

	return gin.WrapH(srv)
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return gin.WrapH(h)
}

func main() {
	_, err := otel.SetupGlobalOTelSDK(context.Background(), "srv.middleware", "0.1.0")
	if err != nil {
		fmt.Println(err)
	}

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(wrapper.TracerMiddleware())
	r.Use(cors.New(config))
	r.Use(wrapper.GinContextToContextMiddleware())
	r.Use(loaders.Middleware())

	r.Any("/query", graphqlHandler())
	r.GET("/", playgroundHandler())

	appPort := viper.GetInt("app.port")
	fmt.Printf("[running]: localhost:%v\n", appPort)

	r.Run(fmt.Sprintf(":%v", appPort))
}
