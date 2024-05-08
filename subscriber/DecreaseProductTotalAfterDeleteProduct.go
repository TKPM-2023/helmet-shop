package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/category/categorystorage"
	"github.com/orgball2608/helmet-shop-be/pubsub"
)

func DecreaseProductTotalAfterDeleteProduct(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease total product after delete product",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := categorystorage.NewSQLStore(appCtx.GetMainDBConnection())
			productData := message.Data().(HasCategoryId)
			return store.DecreaseTotalProduct(ctx, productData.GetCategoryID())
		},
	}
}
