package cartbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"TKPM-Go/pubsub"
	"context"
)

type RemoveProductsStorage interface {
	RemoveProductsFromCart(ctx context.Context, cartID int, data cartmodel.RemoveCartProducts) error
}

type removeProductsBusiness struct {
	store  RemoveProductsStorage
	pubsub pubsub.Pubsub
}

func NewRemoveProductsBusiness(store RemoveProductsStorage, pubsub pubsub.Pubsub) *removeProductsBusiness {
	return &removeProductsBusiness{store: store, pubsub: pubsub}
}

func (business *removeProductsBusiness) RemoveProducts(context context.Context, id int, data cartmodel.RemoveCartProducts) error {
	if err := business.store.RemoveProductsFromCart(context, id, data); err != nil {
		return err
	}

	business.pubsub.Publish(context, common.TopicRemoveProductsFromCart, pubsub.NewMessage(&cartmodel.ProductTotalUpdate{CartId: id, Quantity: len(data)}))

	return nil
}
