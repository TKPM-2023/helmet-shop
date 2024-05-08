package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/category/categorystorage"
	"github.com/orgball2608/helmet-shop-be/pubsub"
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
			productData := message.Data().(HasCategoryId)
			return store.IncreaseTotalProduct(ctx, productData.GetCategoryID())
		},
	}
}
