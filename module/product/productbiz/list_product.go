package productbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
)

type ListProductStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}

type listProdcutBusiness struct {
	store ListProductStore
}

func NewListProductBusiness(store ListProductStore) *listProdcutBusiness {
	return &listProdcutBusiness{store: store}
}

func (business *listProdcutBusiness) ListProduct(context context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.Product, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging, "Ratings", "Ratings.User")
	if err != nil {
		return nil, err
	}

	return result, nil
}
