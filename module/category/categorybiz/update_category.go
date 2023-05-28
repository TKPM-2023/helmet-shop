package categorybiz

import (
	"TKPM-Go/module/category/categorymodel"
	"context"
	"errors"
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

	if result.Status == 0 {
		return errors.New("data deleted")
	}

	if err := business.store.UpdateCategory(context, id, data); err != nil {
		return err
	}
	return nil

}
