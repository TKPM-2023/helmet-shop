package ratingbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"TKPM-Go/pubsub"
	"context"
)

type DeleteRatingStore interface {
	FindDataWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*ratingmodel.Rating, error)
	DeleteRating(context context.Context, id int) error
}

type deleteRatingBusiness struct {
	store  DeleteRatingStore
	pubsub pubsub.Pubsub
}

func NewDeleteRatingBusiness(store DeleteRatingStore, pubsub pubsub.Pubsub) *deleteRatingBusiness {
	return &deleteRatingBusiness{store: store, pubsub: pubsub}
}

func (business *deleteRatingBusiness) DeleteRating(context context.Context, id int) error {
	oldData, err := business.store.FindDataWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(ratingmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(ratingmodel.EntityName, err)
	}

	if err := business.store.DeleteRating(context, id); err != nil {
		return common.ErrCannotDeleteEntity(ratingmodel.EntityName, err)
	}

	business.pubsub.Publish(context, common.TopicUserDeleteRatingProduct, pubsub.NewMessage(&ratingmodel.Rating{ProductID: oldData.ProductID}))

	return nil
}
