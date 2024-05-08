package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartstorage"
	"github.com/orgball2608/helmet-shop-be/pubsub"
)

func DecreaseProductTotalAfterRemoveProductsFromCart(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Decrease total product after removes products from cart",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())
			data := message.Data().(UpdateProductTotal)
			return store.DecreaseTotalProduct(ctx, data.GetCartID(), data.GetQuantity())
		},
	}
}
