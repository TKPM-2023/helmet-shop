package subscriber

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/cart/cartstorage"
	"TKPM-Go/pubsub"
	"context"
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
