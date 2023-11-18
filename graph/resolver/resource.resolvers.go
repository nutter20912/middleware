package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"context"
	"fmt"
	"middleware/api/grpc"
	"middleware/graph/loaders"
	"middleware/graph/model"
	orderV1 "middleware/proto/order/v1"
	"strings"
)

// User is the resolver for the user field.
func (r *commentResolver) User(ctx context.Context, obj *model.Comment) (*model.User, error) {
	result, err := loaders.GetUser(ctx, obj.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

// Events is the resolver for the events field.
func (r *depositOrderResolver) Events(ctx context.Context, obj *model.DepositOrder) ([]*model.DepositOrderEvent, error) {
	rsp, err := grpc.NewOrderServiceClient().GetDepositEvent(ctx, &orderV1.GetDepositEventRequest{OrderId: obj.ID})
	if err != nil {
		return nil, err
	}

	var events []*model.DepositOrderEvent
	for _, item := range rsp.Data {
		events = append(events, &model.DepositOrderEvent{
			UserID:  item.UserId,
			OrderID: item.OrderId,
			Amount:  item.Amount,
			Memo:    item.Memo,
			Time:    item.Time,
			Status: model.DepositStatus(
				strings.TrimPrefix(item.Status.String(), "DEPOSIT_STATUS_"),
			),
		})
	}

	return events, nil
}

// User is the resolver for the user field.
func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	result, err := loaders.GetUser(ctx, obj.UserId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result, nil
}

// Comment returns CommentResolver implementation.
func (r *Resolver) Comment() CommentResolver { return &commentResolver{r} }

// DepositOrder returns DepositOrderResolver implementation.
func (r *Resolver) DepositOrder() DepositOrderResolver { return &depositOrderResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

type commentResolver struct{ *Resolver }
type depositOrderResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
