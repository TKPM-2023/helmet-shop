package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product/productstorage"
	"github.com/orgball2608/helmet-shop-be/pubsub"
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
