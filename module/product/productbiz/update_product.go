package productbiz

import (
	"context"
	"fmt"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
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

	if result == nil {
		return common.ErrEntityNotFound(productmodel.EntityName, nil)
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(productmodel.EntityName, nil)
	}

	if data.Name != "" {
		fmt.Println(data)
		result, err := business.store.FindProductWithCondition(context, map[string]interface{}{
			"name": data.Name,
		})

		if err != nil {
			return err
		}

		if result != nil {
			return productmodel.ErrProductNameExisted
		}
	}

	if err := business.store.UpdateProduct(context, id, data); err != nil {
		return err
	}
	return nil
}
