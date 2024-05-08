package categorybiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
)

type UpdateCategoryStore interface {
	UpdateCategory(context context.Context, id int, data *categorymodel.CategoryUpdate) error
	FindCategoryWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type updateCategoryBusiness struct {
	store UpdateCategoryStore
}

func NewUpdateCategoryBusiness(store UpdateCategoryStore) *updateCategoryBusiness {
	return &updateCategoryBusiness{store: store}
}

func (business *updateCategoryBusiness) UpdateCategory(context context.Context, id int,
	data *categorymodel.CategoryUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindCategoryWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return err
	}

	if result == nil {
		return common.ErrEntityNotFound(categorymodel.EntityName, nil)
	}

	if result.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, nil)
	}

	if data.Name != "" {
		result, err := business.store.FindCategoryWithCondition(context, map[string]interface{}{
			"name": data.Name,
		})

		if err != nil {
			return err
		}

		if result != nil {
			return categorymodel.ErrCategoryNameExisted
		}
	}

	if err := business.store.UpdateCategory(context, id, data); err != nil {
		return common.ErrCannotUpdateEntity(categorymodel.EntityName, err)
	}
	return nil

}
