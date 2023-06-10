package subscriber

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product/productstorage"
	"TKPM-Go/pubsub"
	"context"
)

func DecreaseTotalRatingAfterUserRemoveRating(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease Rating count after user remove rating product",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
			ratingData := message.Data().(HasProductId)
			return store.DecreaseTotalRating(ctx, ratingData.GetProductID())
		},
	}
}
