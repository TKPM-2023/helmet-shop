package categorybiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
)

type DeleteCategoryStore interface {
	DeleteCategory(context context.Context, id int) error
	FindCategoryWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type deleteCategoryBusiness struct {
	store DeleteCategoryStore
}

func NewDeleteCategoryBusiness(store DeleteCategoryStore) *deleteCategoryBusiness {
	return &deleteCategoryBusiness{store: store}
}

func (business *deleteCategoryBusiness) DeleteCategory(context context.Context, id int) error {
	oldData, err := business.store.FindCategoryWithCondition(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(categorymodel.EntityName, err)
	}

	if err := business.store.DeleteCategory(context, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}
	return nil
}
