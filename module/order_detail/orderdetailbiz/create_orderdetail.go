package orderdetailbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"context"
)

type CreateOrderDetailStore interface {
	CreateOrderDetail(ctx context.Context, data *orderdetailmodel.OrderDetailCreate) error
	FindOrderDetailWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*orderdetailmodel.OrderDetail, error)
}

type createOrderDetailBusiness struct {
	store CreateOrderDetailStore
}

func NewCreateOrderDetailBusiness(store CreateOrderDetailStore) *createOrderDetailBusiness {
	return &createOrderDetailBusiness{store: store}
}

func (business *createOrderDetailBusiness) CreateOrderDetail(context context.Context, data *orderdetailmodel.OrderDetailCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	/*
		_, err := business.store.FindOrderDetailWithCondition(context, map[string]interface{}{"user_id": data.User_ID})

		if err != nil {
			return err
		}

		if result != nil {
			return common.ErrEntityExisted(orderdetailmodel.EntityName, nil)
		}
	*/
	if err := business.store.CreateOrderDetail(context, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.EntityName, err)
	}
	return nil
}
