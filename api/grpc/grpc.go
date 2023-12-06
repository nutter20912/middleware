package grpc

import (
	"context"
	"fmt"
	_ "middleware/config"
	"middleware/logging"
	boardV1 "middleware/proto/board/v1"
	marketV1 "middleware/proto/market/v1"
	notifyV1 "middleware/proto/notify/v1"
	orderV1 "middleware/proto/order/v1"
	userV1 "middleware/proto/user/v1"
	walletV1 "middleware/proto/wallet/v1"
	"middleware/wrapper"

	"github.com/go-micro/plugins/v4/client/grpc"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/go-micro/plugins/v4/wrapper/trace/opentelemetry"
	"github.com/spf13/viper"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/metadata"
	"go-micro.dev/v4/registry"
)

type authWrapper struct {
	client.Client
}

var (
	USER_SERVICE_NAME   = viper.GetString("grpc.service.user.name")
	BOARD_SERVICE_NAME  = viper.GetString("grpc.service.board.name")
	WALLET_SERVICE_NAME = viper.GetString("grpc.service.wallet.name")
	ORDER_SERVICE_NAME  = viper.GetString("grpc.service.order.name")
	MARKET_SERVICE_NAME = viper.GetString("grpc.service.market.name")
	NOTIFY_SERVICE_NAME = viper.GetString("grpc.service.notify.name")
	service             micro.Service
)

func init() {
	re := consul.NewRegistry(registry.Addrs(":8500"))
	s := micro.NewService(
		micro.Name("User.Client"),
		micro.Registry(re),
		micro.WrapClient(
			opentelemetry.NewClientWrapper(),
			newAuthWrapper(),
		))

	s.Init()

	service = s
}

func newAuthWrapper() client.Wrapper {
	return func(c client.Client) client.Client {
		return &authWrapper{grpc.NewClient()}
	}
}

func (a *authWrapper) Call(
	ctx context.Context,
	req client.Request,
	rsp interface{},
	opts ...client.CallOption,
) (err error) {
	if gc, err := wrapper.GinContextFromContext(ctx); err == nil {
		ctx = metadata.Set(ctx, "Authorization", gc.Request.Header.Get("Authorization"))
	} else {
		fmt.Println(err.Error())
	}

	defer logging.ClientRequestLog(ctx, req, rsp)

	return a.Client.Call(ctx, req, rsp, opts...)
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

func NewMarketServiceClient() marketV1.MarketService {
	return marketV1.NewMarketService(MARKET_SERVICE_NAME, service.Client())
}

func NewNotifyServiceClient() notifyV1.NotifyService {
	return notifyV1.NewNotifyService(NOTIFY_SERVICE_NAME, service.Client())
}
