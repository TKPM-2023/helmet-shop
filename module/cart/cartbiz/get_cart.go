package cartbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"context"
	"fmt"
)

type GetCartStorage interface {
	FindCartWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*cartmodel.Cart, error)
}

type getCartBusiness struct {
	store GetCartStorage
}

func NewGetCartBusiness(store GetCartStorage) *getCartBusiness {
	return &getCartBusiness{store: store}
}

func (business *getCartBusiness) GetCart(context context.Context, id int) (*cartmodel.Cart, error) {
	fmt.Println(id)
	result, err := business.store.FindCartWithCondition(context, map[string]interface{}{"id": id}, "CartProducts.Product", "CartProducts")

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(cartmodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(cartmodel.EntityName, err)
	}

	fmt.Println(result)

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(cartmodel.EntityName, err)
	}
	return result, err
}
