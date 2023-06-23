package subscriber

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/cart/cartstorage"
	"TKPM-Go/pubsub"
	"context"
)

type UpdateProductTotal interface {
	GetCartID() int
	GetQuantity() int
}

func IncreaseProductTotalAfterAddProductsToCart(
	appCtx appctx.AppContext) consumerJob {
	return consumerJob{
		Title: "Increase total product after add products to cart",
		Hld: func(ctx context.Context, message *pubsub.Message) error {
			store := cartstorage.NewSQLStore(appCtx.GetMainDBConnection())
			data := message.Data().(UpdateProductTotal)
			return store.IncreaseTotalProduct(ctx, data.GetCartID(), data.GetQuantity())
		},
	}
}
