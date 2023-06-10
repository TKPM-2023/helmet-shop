package subscriber

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorystorage"
	"TKPM-Go/pubsub"
	"context"
)

type HasCategoryId interface {
	GetCategoryID() int
}

func IncreaseProductTotalAfterAddProduct(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase Product count after add product",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasCategoryId)
			return store.IncreaseTotalProduct(ctx, likeData.GetCategoryID())
		},
	}
}
