package loaders

import (
	"context"
	"encoding/json"
	"middleware/api/grpc"
	"middleware/graph/model"
	userV1 "middleware/proto/user/v1"

	"github.com/graph-gophers/dataloader/v7"
)

type userReader struct{}

func (u *userReader) getUsers(ctx context.Context, userIds []string) []*dataloader.Result[*model.User] {
	rsp, err := grpc.NewUserServiceClient().GetList(ctx, &userV1.GetListRequest{UserId: userIds})
	if err != nil {
		return handleError[*model.User](len(userIds), err)
	}

	result := make([]*dataloader.Result[*model.User], 0, len(userIds))

	for _, v := range rsp.Data {
		var user model.User
		userBytes, _ := json.Marshal(v)

		if err := json.Unmarshal(userBytes, &user); err != nil {
			result = append(result, &dataloader.Result[*model.User]{Error: err})
			continue
		}
		result = append(result, &dataloader.Result[*model.User]{Data: &user})
	}

	return result
}

func GetUser(ctx context.Context, userID string) (*model.User, error) {
	loaders := GetLoadersFromContext(ctx)
	return loaders.UserLoader.Load(ctx, userID)()
}

func GetUsers(ctx context.Context, userIDs []string) ([]*model.User, []error) {
	loaders := GetLoadersFromContext(ctx)
	return loaders.UserLoader.LoadMany(ctx, userIDs)()
}
