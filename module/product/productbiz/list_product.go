package productbiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
)

type ListProductStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *productmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}

type listProductBusiness struct {
	store ListProductStore
}

func NewListProductBusiness(store ListProductStore) *listProductBusiness {
	return &listProductBusiness{store: store}
}

func (business *listProductBusiness) ListProduct(context context.Context,
	filter *productmodel.Filter,
	paging *common.Paging,
) ([]productmodel.Product, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging, "Ratings", "Ratings.User")
	if err != nil {
		return nil, common.ErrCannotListEntity(productmodel.EntityName, err)
	}

	return result, nil
}
