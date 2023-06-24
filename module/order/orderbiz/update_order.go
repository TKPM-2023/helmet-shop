package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type UpdateOrderStore interface {
	UpdateOrder(context context.Context, id int, data *ordermodel.OrderUpdate) error
	FindOrderWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ordermodel.Order, error)
}

type updateOrderBusiness struct {
	store UpdateOrderStore
}

func NewUpdateOrderBusiness(store UpdateOrderStore) *updateOrderBusiness {
	return &updateOrderBusiness{store: store}
}

func (business *updateOrderBusiness) UpdateOrder(context context.Context, id int,
	data *ordermodel.OrderUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindOrderWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(ordermodel.EntityName, err)
	}

	if err := business.store.UpdateOrder(context, id, data); err != nil {
		return common.ErrCannotUpdateEntity(ordermodel.EntityName, err)
	}
	return nil

}
