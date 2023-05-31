package productbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"context"
)

type DeleteProductStore interface {
	FindProductWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
	DeleteProduct(context context.Context, id int) error
}

type deleteProductBusiness struct {
	store DeleteProductStore
}

func NewDeleteProductBusiness(store DeleteProductStore) *deleteProductBusiness {
	return &deleteProductBusiness{store: store}
}

func (business *deleteProductBusiness) DeleteProduct(context context.Context, id int) error {
	oldData, err := business.store.FindProductWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(productmodel.EntityName, err)
	}

	if err := business.store.DeleteProduct(context, id); err != nil {
		return common.ErrCannotDeleteEntity(productmodel.EntityName, err)
	}
	return nil
}
