package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product/productstorage"
	"github.com/orgball2608/helmet-shop-be/pubsub"
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
