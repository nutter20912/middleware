package grpc

import (
	"context"
	"fmt"
	"middleware/middleware"
	boardV1 "middleware/proto/board/v1"
	orderV1 "middleware/proto/order/v1"
	userV1 "middleware/proto/user/v1"
	walletV1 "middleware/proto/wallet/v1"

	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
)

type authWrapper struct {
	client.Client
}

var (
	USER_SERVICE_NAME   = "srv.user"
	BOARD_SERVICE_NAME  = "srv.board"
	WALLET_SERVICE_NAME = "srv.wallet"
	ORDER_SERVICE_NAME  = "srv.order"
)

var (
	service micro.Service
)

func init() {
	re := consul.NewRegistry(registry.Addrs(":8500"))
	s := micro.NewService(
		micro.Name("User.Client"),
		micro.Registry(re),
		micro.WrapClient(func(c client.Client) client.Client {
			return &authWrapper{grpc.NewClient()}
		}),
	)
	s.Init()

	service = s
}

func (a *authWrapper) Call(
	ctx context.Context,
	req client.Request, rsp interface{},
	opts ...client.CallOption,
) error {
	if gc, err := middleware.GinContextFromContext(ctx); err == nil {
		ctx = metadata.Set(ctx, "Authorization", gc.Request.Header.Get("Authorization"))
	} else {
		fmt.Println(err.Error())
	}

	return a.Client.Call(ctx, req, rsp)
}

func NewUserServiceClient() userV1.UserService {
	return userV1.NewUserService(USER_SERVICE_NAME, service.Client())
}

func NewPostServiceClient() boardV1.PostService {
	return boardV1.NewPostService(BOARD_SERVICE_NAME, service.Client())
}

func NewCommentServiceClient() boardV1.CommentService {
	return boardV1.NewCommentService(BOARD_SERVICE_NAME, service.Client())
}

func NewWalletServiceClient() walletV1.WalletService {
	return walletV1.NewWalletService(WALLET_SERVICE_NAME, service.Client())
}

func NewOrderServiceClient() orderV1.OrderService {
	return orderV1.NewOrderService(ORDER_SERVICE_NAME, service.Client())
}
