package orderbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
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

	if data.OrderStatus != 0 && data.OrderStatus >= result.OrderStatus {
		if data.OrderStatus == 4 && (result.OrderStatus == 2 || result.OrderStatus == 3) {
			return common.ErrCannotUpdateEntity(ordermodel.EntityName, nil)
		}

		if data.OrderStatus == 3 && result.OrderStatus == 1 {
			return common.ErrCannotUpdateEntity(ordermodel.EntityName, nil)
		}
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(ordermodel.EntityName, err)
	}

	if err := business.store.UpdateOrder(context, id, data); err != nil {
		return common.ErrCannotUpdateEntity(ordermodel.EntityName, err)
	}
	return nil

}
