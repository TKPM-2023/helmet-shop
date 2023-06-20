package cartbiz

import (
	"TKPM-Go/module/cart/cartmodel"
	"context"
)

type AddProductsStorage interface {
	AddProductsToCart(ctx context.Context, cartID int, data cartmodel.CartProductDetails) error
}

type addProductsBusiness struct {
	store AddProductsStorage
}

func NewAddProductsBusiness(store AddProductsStorage) *addProductsBusiness {
	return &addProductsBusiness{store: store}
}

func (business *addProductsBusiness) AddProducts(context context.Context, id int, data cartmodel.CartProductDetails) error {
	if err := business.store.AddProductsToCart(context, id, data); err != nil {
		return err
	}
	return nil
}
