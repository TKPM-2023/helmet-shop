package ratingbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"TKPM-Go/module/product_rating/ratingmodel"
	"TKPM-Go/pubsub"
	"context"
	"fmt"
)

type CreateRatingStore interface {
	CreateRating(ctx context.Context, data *ratingmodel.RatingCreate) error
}

type GetOrderDetailStore interface {
	FindOrderDetailWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*orderdetailmodel.OrderDetail, error)
}

type createRatingBusiness struct {
	store      CreateRatingStore
	orderStore GetOrderDetailStore
	pubsub     pubsub.Pubsub
}

func NewCreateRatingBusiness(store CreateRatingStore, orderStore GetOrderDetailStore, pubsub pubsub.Pubsub) *createRatingBusiness {
	return &createRatingBusiness{store: store, orderStore: orderStore, pubsub: pubsub}
}

func (business *createRatingBusiness) CreateRating(context context.Context, data *ratingmodel.RatingCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	fmt.Println(data.OrderDetailId)

	detail, err := business.orderStore.FindOrderDetailWithCondition(context, map[string]interface{}{"id": data.OrderDetailId}, "Order")

	if err != nil {
		return common.ErrCannotGetEntity(ordermodel.EntityName, err)
	}

	fmt.Println(detail.Order)

	if detail.Status == 0 {
		return common.ErrEntityDeleted(ordermodel.EntityName, err)
	}

	if detail.Order.OrderStatus != 2 {
		return ratingmodel.ErrCannotCreateRating
	}

	if err := business.store.CreateRating(context, data); err != nil {
		return common.ErrCannotCreateEntity(ratingmodel.EntityName, err)
	}

	business.pubsub.Publish(context, common.TopicUserRatingProduct, pubsub.NewMessage(data))

	return nil

}
