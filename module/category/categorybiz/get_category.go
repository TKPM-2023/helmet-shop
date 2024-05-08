package categorybiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
)

type GetCategoryStore interface {
	FindCategoryWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type getCategoryBusiness struct {
	store GetCategoryStore
}

func NewGetCategoryBusiness(store GetCategoryStore) *getCategoryBusiness {
	return &getCategoryBusiness{store: store}
}

func (business *getCategoryBusiness) GetCategory(
	ctx context.Context, id int) (*categorymodel.Category, error) {
	result, err := business.store.FindCategoryWithCondition(ctx, map[string]interface{}{"id": id}, "Products")

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)

		}
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(categorymodel.EntityName, err)
	}

	return result, err
}
