package productbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
)

type FindProductsStore interface {
	FindProductsByName(ctx context.Context,
		name string,
		moreKeys ...string,
	) ([]productmodel.Product, error)
}

type findProductsBusiness struct {
	store FindProductsStore
}

func NewFindProductsBusiness(store FindProductsStore) *findProductsBusiness {
	return &findProductsBusiness{store: store}
}

func (business *findProductsBusiness) FindProductsByName(context context.Context, name string) ([]productmodel.Product, error) {
	result, err := business.store.FindProductsByName(context, name, "Ratings", "Ratings.User")

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(productmodel.EntityName, err)
	}
	return result, err
}
