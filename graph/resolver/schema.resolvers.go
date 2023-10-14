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
	orderV1 "middleware/proto/order/v1"
	walletV1 "middleware/proto/wallet/v1"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"
)

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	rsp, err := grpc.NewUserServiceClient().Get(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	} else {
		fmt.Println(rsp)
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
	} else {
		fmt.Println(rsp)
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
	} else {
		fmt.Println(rsp)
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
	} else {
		fmt.Println(rsp)
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
func (r *queryResolver) WalletEvents(ctx context.Context, cursor *string) (*model.WalletEvents, error) {
	req := walletV1.GetEventRequest{}

	if cursor != nil && *cursor != "" {
		req.Cursor = cursor
	}

	rsp, err := grpc.NewWalletServiceClient().GetEvent(ctx, &req)
	if err != nil {
		return nil, err
	} else {
		fmt.Println(rsp)
	}

	var events []*model.WalletEvent
	for _, item := range rsp.Data {
		events = append(events, &model.WalletEvent{
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

	response := model.WalletEvents{
		Data: events,
		//Paginator: &model.Paginator{
		//	NextCursor: &rsp.Paginator.NextCursor,
		//},
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
		ID:     orderRsp.Data.Id,
		UserID: orderRsp.Data.UserId,
		Amount: orderRsp.Data.Amount,
		Memo:   orderRsp.Data.Memo,
		Status: model.DepositStatus(
			strings.TrimPrefix(orderRsp.Data.Status.String(), "DEPOSIT_STATUS_"),
		),
	}

	return order, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
