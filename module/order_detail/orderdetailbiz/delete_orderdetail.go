package orderdetailbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
)

type DeleteOrderDetailStore interface {
	FindOrderDetailWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*orderdetailmodel.OrderDetail, error)
	DeleteOrderDetail(context context.Context, id int) error
}

type deleteOrderDetailBusiness struct {
	store DeleteOrderDetailStore
}

func NewDeleteOrderDetailBusiness(store DeleteOrderDetailStore) *deleteOrderDetailBusiness {
	return &deleteOrderDetailBusiness{store: store}
}

func (business *deleteOrderDetailBusiness) DeleteOrderDetail(context context.Context, id int) error {
	oldData, err := business.store.FindOrderDetailWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(orderdetailmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(orderdetailmodel.EntityName, err)
	}

	if err := business.store.DeleteOrderDetail(context, id); err != nil {
		return common.ErrCannotDeleteEntity(orderdetailmodel.EntityName, err)
	}
	return nil
}
