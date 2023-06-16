package orderdetailbiz

import (
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
	"errors"
)

type UpdateOrderDetailStore interface {
	UpdateOrderDetail(context context.Context, id int, data *orderdetailmodel.OrderDetailUpdate) error
	FindOrderDetailWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*orderdetailmodel.OrderDetail, error)
}

type updateOrderDetailBusiness struct {
	store UpdateOrderDetailStore
}


func NewUpdateOrderDetailBusiness(store UpdateOrderDetailStore) *updateOrderDetailBusiness {
	return &updateOrderDetailBusiness{store: store}
}

func (business *updateOrderDetailBusiness) UpdateOrderDetail(context context.Context, id int,
	data *orderdetailmodel.OrderDetailUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindOrderDetailWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := business.store.UpdateOrderDetail(context, id, data); err != nil {
		return err
	}
	return nil

}
