package orderdetailbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
)

type GetOrderDetailStore interface {
	FindOrderDetailWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*orderdetailmodel.OrderDetail, error)
}

type getOrderDetailBusiness struct {
	store GetOrderDetailStore
}

func NewGetOrderDetailBusiness(store GetOrderDetailStore) *getOrderDetailBusiness {
	return &getOrderDetailBusiness{store: store}
}

func (business *getOrderDetailBusiness) GetOrderDetail(
	ctx context.Context, id int) (*orderdetailmodel.OrderDetail, error) {
	result, err := business.store.FindOrderDetailWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(orderdetailmodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(orderdetailmodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(orderdetailmodel.EntityName, err)
	}

	return result, err
}
