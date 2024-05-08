package orderbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
)

type DeleteOrderStore interface {
	FindOrderWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ordermodel.Order, error)
	DeleteOrder(context context.Context, id int) error
}

type deleteOrderBusiness struct {
	store DeleteOrderStore
}

func NewDeleteOrderBusiness(store DeleteOrderStore) *deleteOrderBusiness {
	return &deleteOrderBusiness{store: store}
}

func (business *deleteOrderBusiness) DeleteOrder(context context.Context, id int) error {
	oldData, err := business.store.FindOrderWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(ordermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(ordermodel.EntityName, err)
	}

	if err := business.store.DeleteOrder(context, id); err != nil {
		return common.ErrCannotDeleteEntity(ordermodel.EntityName, err)
	}
	return nil
}
