package orderbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"context"
)

type CreateOrderStore interface {
	CreateOrder(ctx context.Context, data *ordermodel.OrderCreate) error
	FindOrderWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKey ...string,
	) (*ordermodel.Order, error)
}

type createOrderBusiness struct {
	store CreateOrderStore
}

func NewCreateOrderBusiness(store CreateOrderStore) *createOrderBusiness {
	return &createOrderBusiness{store: store}
}

func (business *createOrderBusiness) CreateOrder(context context.Context, data *ordermodel.OrderCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	_, err := business.store.FindOrderWithCondition(context, map[string]interface{}{"user_id": data.User_ID})

	if err != nil {
		return err
	}

	/*
	if result != nil {
		return common.ErrEntityExisted(ordermodel.EntityName, nil)
	}
	*/
	if err := business.store.CreateOrder(context, data); err != nil {
		return common.ErrCannotCreateEntity(ordermodel.EntityName, err)
	}
	return nil
}
