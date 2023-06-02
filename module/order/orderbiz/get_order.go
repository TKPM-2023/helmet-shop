package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/category/categorymodel"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type GetOrderStore interface {
	FindOrderWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ordermodel.Order, error)
}

type getOrderBusiness struct {
	store GetOrderStore
}

func NewGetOrderBusiness(store GetOrderStore) *getOrderBusiness {
	return &getOrderBusiness{store: store}
}

func (business *getOrderBusiness) GetOrder(context context.Context, id int) (*ordermodel.Order, error) {
	result, err := business.store.FindOrderWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(categorymodel.EntityName, err)
	}
	return result, err
}
