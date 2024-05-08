package orderdetailbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
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

	if err := business.store.CreateOrderDetail(context, data); err != nil {
		return common.ErrCannotCreateEntity(orderdetailmodel.EntityName, err)
	}
	return nil
}
