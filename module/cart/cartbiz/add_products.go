package cartbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartmodel"
	"github.com/orgball2608/helmet-shop-be/pubsub"
)

type AddProductsStorage interface {
	AddProductsToCart(ctx context.Context, cartID int, data cartmodel.CartProductDetails) error
}

type addProductsBusiness struct {
	store  AddProductsStorage
	pubsub pubsub.Pubsub
}

func NewAddProductsBusiness(store AddProductsStorage, pubsub pubsub.Pubsub) *addProductsBusiness {
	return &addProductsBusiness{store: store, pubsub: pubsub}
}

func (business *addProductsBusiness) AddProducts(context context.Context, id int, data cartmodel.CartProductDetails) error {
	if err := business.store.AddProductsToCart(context, id, data); err != nil {
		return err
	}

	business.pubsub.Publish(context, common.TopicAddProductsToCart, pubsub.NewMessage(&cartmodel.ProductTotalUpdate{CartId: id, Quantity: len(data)}))

	return nil
}
