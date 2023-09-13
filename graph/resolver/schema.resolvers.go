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

	"google.golang.org/protobuf/types/known/emptypb"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

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

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
