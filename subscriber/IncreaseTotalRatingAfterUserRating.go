package subscriber

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product/productstorage"
	"TKPM-Go/pubsub"
	"context"
)

type HasProductId interface {
	GetProductID() int
}

func IncreaseTotalRatingAfterUserRating(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase Rating count after user rating product",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := productstorage.NewSQLStore(appCtx.GetMainDBConnection())
			ratingData := message.Data().(HasProductId)
			return store.IncreaseTotalRating(ctx, ratingData.GetProductID())
		},
	}
}
