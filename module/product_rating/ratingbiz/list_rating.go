package ratingbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"context"
)

type ListRatingStore interface {
	ListRatingWithCondition(
		ctx context.Context,
		filter *ratingmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]ratingmodel.Rating, error)
}

type listRatingBusiness struct {
	store ListRatingStore
}

func NewListRatingBusiness(store ListRatingStore) *listRatingBusiness {
	return &listRatingBusiness{store: store}
}

func (business *listRatingBusiness) ListRating(context context.Context,
	filter *ratingmodel.Filter,
	paging *common.Paging,
) ([]ratingmodel.Rating, error) {
	result, err := business.store.ListRatingWithCondition(context, filter, paging,"User","Product")
	if err != nil {
		return nil, common.ErrCannotListEntity(ratingmodel.EntityName, err)
	}

	return result, nil
}
