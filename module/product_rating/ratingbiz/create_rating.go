package ratingbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"TKPM-Go/pubsub"
	"context"
)

type CreateRatingStore interface {
	CreateRating(ctx context.Context, data *ratingmodel.RatingCreate) error
}

type createRatingBusiness struct {
	store  CreateRatingStore
	pubsub pubsub.Pubsub
}

func NewCreateRatingBusiness(store CreateRatingStore, pubsub pubsub.Pubsub) *createRatingBusiness {
	return &createRatingBusiness{store: store, pubsub: pubsub}
}

func (business *createRatingBusiness) CreateRating(context context.Context, data *ratingmodel.RatingCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.store.CreateRating(context, data); err != nil {
		return common.ErrCannotCreateEntity(ratingmodel.EntityName, err)
	}

	business.pubsub.Publish(context, common.TopicUserRatingProduct, pubsub.NewMessage(data))

	return nil

}
