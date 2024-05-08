package ratingbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingmodel"
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
	result, err := business.store.ListRatingWithCondition(context, filter, paging, "User", "OrderDetail", "Product")
	if err != nil {
		return nil, common.ErrCannotListEntity(ratingmodel.EntityName, err)
	}

	return result, nil
}
