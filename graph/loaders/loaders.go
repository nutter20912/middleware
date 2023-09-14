package loaders

import (
	"context"
	"middleware/graph/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/dataloader/v7"
)

type ctxKey string

const (
	LOADER_KEY = ctxKey("dataloaders")
)

type Loaders struct {
	UserLoader *dataloader.Loader[string, *model.User]
}

func GetLoadersFromContext(ctx context.Context) *Loaders {
	return ctx.Value(LOADER_KEY).(*Loaders)
}

func NewLoaders() *Loaders {
	ur := &userReader{}

	return &Loaders{
		UserLoader: dataloader.NewBatchedLoader(
			ur.getUsers,
			dataloader.WithWait[string, *model.User](time.Millisecond),
		),
	}
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loader := NewLoaders()

		ctx := context.WithValue(c.Request.Context(), LOADER_KEY, loader)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func handleError[T any](itemsLength int, err error) []*dataloader.Result[T] {
	result := make([]*dataloader.Result[T], itemsLength)
	for i := 0; i < itemsLength; i++ {
		result[i] = &dataloader.Result[T]{Error: err}
	}
	return result
}
