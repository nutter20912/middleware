package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"encoding/json"
	"fmt"
	"middleware/api/grpc"
	"middleware/graph/model"
	boardV1 "middleware/proto/board/v1"
	marketV1 "middleware/proto/market/v1"
	orderV1 "middleware/proto/order/v1"
	walletV1 "middleware/proto/wallet/v1"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/protobuf/types/known/emptypb"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	rsp, err := grpc.NewUserServiceClient().Get(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	rspString, _ := json.Marshal(rsp.Data)

	var user model.User
	if err = json.Unmarshal(rspString, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, cursor *string) (*model.Posts, error) {
	req := boardV1.PostServiceGetAllRequest{}

	if cursor != nil && *cursor != "" {
		req.Cursor = cursor
	}

	rsp, err := grpc.NewPostServiceClient().GetAll(ctx, &req)
	if err != nil {
		return nil, err
	}

	rspString, _ := json.Marshal(rsp.Data)
	var posts []*model.Post

	if err = json.Unmarshal(rspString, &posts); err != nil {
		return nil, err
	}

	postsResponse := model.Posts{
		Data: posts,
		Paginator: &model.Paginator{
			NextCursor: &rsp.Paginator.NextCursor,
		},
	}

	return &postsResponse, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	rsp, err := grpc.NewPostServiceClient().Get(ctx, &boardV1.PostServiceGetRequest{Id: id})
	if err != nil {
		return nil, err
	}

	rspString, _ := json.Marshal(rsp.Data)
	var post *model.Post

	if err = json.Unmarshal(rspString, &post); err != nil {
		return nil, err
	}

	return post, nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID string, cursor *string) (*model.Comments, error) {
	req := boardV1.CommentServiceGetAllRequest{PostId: postID}

	if cursor != nil && *cursor != "" {
		req.Cursor = cursor
	}

	rsp, err := grpc.NewCommentServiceClient().GetAll(ctx, &req)
	if err != nil {
		return nil, err
	}

	rspString, _ := json.Marshal(rsp.Data)
	var comments []*model.Comment

	if err = json.Unmarshal(rspString, &comments); err != nil {
		return nil, err
	}

	commentsResponse := model.Comments{
		Data: comments,
		Paginator: &model.Paginator{
			NextCursor: &rsp.Paginator.NextCursor,
		},
	}

	return &commentsResponse, nil
}

// Wallet is the resolver for the wallet field.
func (r *queryResolver) Wallet(ctx context.Context) (*model.Wallet, error) {
	rsp, err := grpc.NewWalletServiceClient().Get(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	} else {
		fmt.Println(rsp)
	}

	rspString, _ := json.Marshal(rsp.Data)
	var wallet *model.Wallet

	if err = json.Unmarshal(rspString, &wallet); err != nil {
		return nil, err
	}

	return wallet, nil
}

// WalletEvents is the resolver for the walletEvents field.
func (r *queryResolver) WalletEvents(ctx context.Context, page *int64, limit *int64, filter *model.WalletEventFilter) (*model.WalletEvents, error) {
	req := walletV1.GetEventRequest{
		StartDate: filter.StartDate,
		EndDate:   filter.EndDate,
	}

	if page != nil {
		req.Page = page
	}

	if limit != nil {
		req.Limit = limit
	}

	rsp, err := grpc.NewWalletServiceClient().GetEvent(ctx, &req)
	if err != nil {
		return nil, err
	}

	var events []*model.WalletEvent
	dataBytes, _ := json.Marshal(rsp.Data)
	json.Unmarshal(dataBytes, &events)

	var paginator *model.PagePaginator
	bytes, _ := json.Marshal(rsp.Paginator)
	json.Unmarshal(bytes, &paginator)

	response := model.WalletEvents{
		Data:      events,
		Paginator: paginator,
	}

	return &response, nil
}

// DepositOrder is the resolver for the depositOrder field.
func (r *queryResolver) DepositOrder(ctx context.Context, id string) (*model.DepositOrder, error) {
	orderRsp, err := grpc.NewOrderServiceClient().GetDeposit(ctx, &orderV1.GetDepositRequest{Id: id})
	if err != nil {
		return nil, err
	}

	order := &model.DepositOrder{
		ID:        orderRsp.Data.Id,
		UserID:    orderRsp.Data.UserId,
		Amount:    orderRsp.Data.Amount,
		Memo:      orderRsp.Data.Memo,
		CreatedAt: orderRsp.Data.CreatedAt,
		UpdatedAt: orderRsp.Data.UpdatedAt,
		Status: model.DepositStatus(
			strings.TrimPrefix(orderRsp.Data.Status.String(), "DEPOSIT_STATUS_"),
		),
	}

	return order, nil
}

// SpotOrder is the resolver for the spotOrder field.
func (r *queryResolver) SpotOrders(ctx context.Context, page *int64, limit *int64, filter *model.SpotOrderFilter) (*model.SpotOrders, error) {
	req := orderV1.GetSpotRequest{
		StartDate: filter.StartDate,
		EndDate:   filter.EndDate,
	}
	if page != nil {
		req.Page = page
	}

	if limit != nil {
		req.Limit = limit
	}
	rsp, err := grpc.NewOrderServiceClient().GetSpot(ctx, &req)
	if err != nil {
		return nil, err
	}

	var data []*model.SpotOrder
	dataBytes, _ := json.Marshal(rsp.Data)
	json.Unmarshal(dataBytes, &data)

	var paginator *model.PagePaginator
	bytes, _ := json.Marshal(rsp.Paginator)
	json.Unmarshal(bytes, &paginator)

	response := model.SpotOrders{
		Data:      data,
		Paginator: paginator,
	}

	return &response, nil
}

// SpotOrderEvents is the resolver for the spotOrderEvents field.
func (r *queryResolver) SpotOrderEvents(ctx context.Context, orderID string) ([]*model.SpotOrderEvent, error) {
	req := orderV1.GetSpotEventRequest{OrderId: orderID}

	rsp, err := grpc.NewOrderServiceClient().GetSpotEvent(ctx, &req)
	if err != nil {
		return nil, err
	}

	var data []*model.SpotOrderEvent
	dataBytes, _ := json.Marshal(rsp.Data)
	json.Unmarshal(dataBytes, &data)

	return data, nil
}

// SpotPositions is the resolver for the spotPositions field.
func (r *queryResolver) SpotPositions(ctx context.Context, page *int64, limit *int64, symbol string) (*model.SpotPositions, error) {
	req := &orderV1.GetSpotPositionRequest{Symbol: symbol}
	if page != nil {
		req.Page = page
	}

	if limit != nil {
		req.Limit = limit
	}

	rsp, err := grpc.NewOrderServiceClient().GetSpotPosition(ctx, req)
	if err != nil {
		return nil, err
	}

	var data []*model.SpotPosition
	dataBytes, _ := json.Marshal(rsp.Data)
	json.Unmarshal(dataBytes, &data)

	fmt.Println(string(dataBytes))
	var paginator *model.PagePaginator
	bytes, _ := json.Marshal(rsp.Paginator)
	json.Unmarshal(bytes, &paginator)

	response := model.SpotPositions{
		Data:      data,
		Paginator: paginator,
	}

	return &response, nil
}

// SpotPositionClosed is the resolver for the spotPositionClosed field.
func (r *queryResolver) SpotPositionClosed(ctx context.Context, page *int64, limit *int64, filter model.SpotPositionClosedFilter) (*model.SpotPositionCloseds, error) {
	req := &orderV1.GetSpotPositionClosedRequest{
		StartDate: filter.StartDate,
		EndDate:   filter.EndDate,
		Symbol:    filter.Symbol}

	if page != nil {
		req.Page = page
	}

	if limit != nil {
		req.Limit = limit
	}

	rsp, err := grpc.NewOrderServiceClient().GetSpotPositionClosed(ctx, req)
	if err != nil {
		return nil, err
	}

	var data []*model.SpotPositionClosed
	dataBytes, _ := json.Marshal(rsp.Data)
	json.Unmarshal(dataBytes, &data)

	var paginator *model.PagePaginator
	bytes, _ := json.Marshal(rsp.Paginator)
	json.Unmarshal(bytes, &paginator)

	response := &model.SpotPositionCloseds{
		Data:      data,
		Paginator: paginator}

	return response, nil
}

// Wallet is the resolver for the wallet field.
func (r *subscriptionResolver) Wallet(ctx context.Context, eventCursor *string) (<-chan *model.WalletStream, error) {
	req := walletV1.GetWalletStreamResquest{}
	if eventCursor != nil && *eventCursor != "" {
		req.EventCursor = eventCursor
	}

	stream, err := grpc.NewWalletServiceClient().GetWalletStream(ctx, &req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *model.WalletStream)

	go func() {
		defer close(ch)

		for {
			rsp, err := stream.Recv()
			if err != nil {
				transport.AddSubscriptionError(ctx, gqlerror.Wrap(err))
				return
			}

			var data *model.WalletStream
			bytes, _ := json.Marshal(rsp)
			json.Unmarshal(bytes, &data)

			var events []*model.WalletEvent
			for _, item := range rsp.Events {
				events = append(events, &model.WalletEvent{
					ID:      item.Id,
					UserID:  item.UserId,
					OrderID: item.OrderId,
					Time:    item.Time,
					Change:  item.Change,
					Memo:    item.Memo,
					Type: model.WalletEventType(
						strings.TrimPrefix(item.Type.String(), "WALLET_EVENT_TYPE_"),
					),
				})
			}
			data.Events = events

			select {
			case <-ctx.Done():
				return
			case ch <- data:
			}
		}
	}()

	return ch, nil
}

// Position is the resolver for the position field.
func (r *subscriptionResolver) Position(ctx context.Context, symbol *string) (<-chan *model.PositionStream, error) {
	req := orderV1.GetPositionStreamResquest{Symbol: *symbol}

	stream, err := grpc.NewOrderServiceClient().GetPositionStream(ctx, &req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *model.PositionStream)

	go func() {
		defer close(ch)

		for {
			rsp, err := stream.Recv()
			if err != nil {
				transport.AddSubscriptionError(ctx, gqlerror.Wrap(err))
				return
			}

			var data *model.PositionStream
			bytes, _ := json.Marshal(rsp)
			json.Unmarshal(bytes, &data)

			select {
			case <-ctx.Done():
				return
			case ch <- data:
			}
		}
	}()

	return ch, nil
}

// Trade is the resolver for the trade field.
func (r *subscriptionResolver) Trade(ctx context.Context, symbol *string) (<-chan *model.TradeStream, error) {
	req := marketV1.GetTradeStreamResquest{Symbol: symbol}
	stream, err := grpc.NewMarketServiceClient().GetTradeStream(ctx, &req)
	if err != nil {
		return nil, err
	}

	ch := make(chan *model.TradeStream)

	go func() {
		defer close(ch)

		for {
			rsp, err := stream.Recv()
			if err != nil {
				transport.AddSubscriptionError(ctx, gqlerror.Wrap(err))
				return
			}

			data := &model.TradeStream{}

			switch val := rsp.Data.(type) {
			case *marketV1.GetTradeStreamResponse_AggTradeData:
				data.AggTrade = &model.AggTradeData{
					EventType:       val.AggTradeData.EventType,
					EventTime:       val.AggTradeData.EventTime,
					Symbol:          val.AggTradeData.Symbol,
					Price:           val.AggTradeData.Price,
					Quantity:        val.AggTradeData.Quantity,
					TransactionTime: val.AggTradeData.TransactionTime,
					IsSell:          val.AggTradeData.IsSell,
				}

			case *marketV1.GetTradeStreamResponse_KlineData:
				data.Kline = &model.KlineData{
					EventType: val.KlineData.EventType,
					EventTime: val.KlineData.EventTime,
					Symbol:    val.KlineData.Symbol,
					Kline: &model.Kline{
						StartTime: val.KlineData.Kline.StartTime,
						EndTime:   val.KlineData.Kline.EndTime,
						Symbol:    val.KlineData.Kline.Symbol,
						Interval:  val.KlineData.Kline.Interval,
						Open:      val.KlineData.Kline.Open,
						Close:     val.KlineData.Kline.Close,
						High:      val.KlineData.Kline.High,
						Low:       val.KlineData.Kline.Low,
					},
				}

			case *marketV1.GetTradeStreamResponse_DepthData:
				data.Depth = &model.DepthData{
					Asks: [][]string{},
					Bids: [][]string{},
				}

				for _, v := range val.DepthData.Asks {
					data.Depth.Asks = append(data.Depth.Asks, v.PriceAndQty)
				}

				for _, v := range val.DepthData.Bids {
					data.Depth.Bids = append(data.Depth.Bids, v.PriceAndQty)
				}
			}

			select {
			case <-ctx.Done():
				return
			case ch <- data:
			}
		}
	}()

	return ch, nil
}

// Notify is the resolver for the notify field.
func (r *subscriptionResolver) Notify(ctx context.Context) (<-chan interface{}, error) {
	stream, err := grpc.NewNotifyServiceClient().GetStream(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	ch := make(chan interface{})

	go func() {
		defer close(ch)

		for {
			rsp, err := stream.Recv()
			if err != nil {
				transport.AddSubscriptionError(ctx, gqlerror.Wrap(err))
				return
			}

			var data interface{}
			json.Unmarshal(rsp.Data.Payload, &data)

			select {
			case <-ctx.Done():
				return
			case ch <- data:
			}
		}
	}()

	return ch, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
