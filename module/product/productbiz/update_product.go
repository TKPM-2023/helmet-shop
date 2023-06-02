package productbiz

import (
	"TKPM-Go/module/product/productmodel"
	"context"
	"errors"
)

type UpdateProductStore interface {
	UpdateProduct(context context.Context, id int, data *productmodel.ProductUpdate) error
	FindProductWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type updateProductBusiness struct {
	store UpdateProductStore
}

func NewUpdateProductBusiness(store UpdateProductStore) *updateProductBusiness {
	return &updateProductBusiness{store: store}
}

func (business *updateProductBusiness) UpdateProduct(context context.Context, id int, data *productmodel.ProductUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindProductWithCondition(context, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if result != nil && result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := business.store.UpdateProduct(context, id, data); err != nil {
		return err
	}
	return nil
}
