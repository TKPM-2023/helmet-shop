package subscriber

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorystorage"
	"TKPM-Go/pubsub"
	"context"
)

func DecreaseProductTotalAfterDeleteProduct(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease total product after delete product",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
			likeData := message.Data().(HasCategoryId)
			return store.DecreaseTotalProduct(ctx, likeData.GetCategoryID())
		},
	}
}
