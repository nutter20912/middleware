package main

import (
	"context"
	"fmt"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
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

func main() {
	service := micro.NewService()
	re := consul.NewRegistry(registry.Addrs(":8500"))

	service.Init(
		// micro.Server(getHttpServer()),
		micro.Registry(re),
		micro.WrapHandler(logWrapper),
	)

	service.Run()
}

// func getHttpServer() server.Server {
// 	srv := apiServer.NewServer(
// 		server.Name("srv.gql"),
// 		server.Address(fmt.Sprintf(":%v", defaultPort)),
// 		server.WrapHandler(logWrapper),
// 	)
//
// 	h := serviceHander.NewDefaultServer(handler.NewExecutableSchema(handler.Config{Resolvers: &handler.Resolver{}}))
//
// 	srv.HandleFunc("/", h)
//
// 	return srv
// }

// func initRouter(router *gin.Engine) {
// 	srv := serviceHander.NewDefaultServer(handler.NewExecutableSchema(handler.Config{Resolvers: &handler.Resolver{}}))
// 	router.GET("/", gin.WrapH(playground.Handler("GraphQL playground", "/query")))
// 	router.POST("/query", gin.WrapH(srv))
// }
