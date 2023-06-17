package cartbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"context"
)

type UpdateQuantityStore interface {
	FindCartWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*cartmodel.Cart, error)
	UpdateCartItemQuantity(ctx context.Context, cartID int, data *cartmodel.CartProductDetail) error
}

type updateQuantityBusiness struct {
	store UpdateQuantityStore
}

func NewUpdateQuantityBusiness(store UpdateQuantityStore) *updateQuantityBusiness {
	return &updateQuantityBusiness{store: store}
}

func (business *updateQuantityBusiness) UpdateQuantity(context context.Context, cartID int, data *cartmodel.CartProductDetail) error {
	result, err := business.store.FindCartWithCondition(context, map[string]interface{}{
		"id": cartID,
	})

	if err != nil {
		return err
	}

	if result == nil {
		return common.ErrEntityNotFound(cartmodel.EntityName, nil)
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(cartmodel.EntityName, nil)
	}

	if err := business.store.UpdateCartItemQuantity(context, cartID, data); err != nil {
		return err
	}
	return nil
}
