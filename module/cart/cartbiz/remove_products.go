package cartbiz

import (
	"TKPM-Go/module/cart/cartmodel"
	"context"
)

type RemoveProductsStorage interface {
	RemoveProductsFromCart(ctx context.Context, cartID int, data cartmodel.RemoveCartProducts) error
}

type removeProductsBusiness struct {
	store RemoveProductsStorage
}

func NewRemoveProductsBusiness(store RemoveProductsStorage) *removeProductsBusiness {
	return &removeProductsBusiness{store: store}
}

func (business *removeProductsBusiness) RemoveProducts(context context.Context, id int, data cartmodel.RemoveCartProducts) error {
	if err := business.store.RemoveProductsFromCart(context, id, data); err != nil {
		return err
	}
	return nil
}
