package subscriber

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartstorage"
	"github.com/orgball2608/helmet-shop-be/pubsub"
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
