package orderbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
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

func (business *getOrderBusiness) GetOrder(
	ctx context.Context, id int) (*ordermodel.Order, error) {
	result, err := business.store.FindOrderWithCondition(ctx, map[string]interface{}{"id": id}, "Products", "Contact", "User")

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(ordermodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(ordermodel.EntityName, err)
	}

	return result, err
}
