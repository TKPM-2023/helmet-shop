package categorybiz

import (
	"LearnGo/common"
	"LearnGo/module/category/categorymodel"
	"context"
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
	result, err := business.store.FindCategoryWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)

		}
		// OR: able to throw err `sth went wrong with server`
		return nil, common.ErrCannotGetEntity(categorymodel.EntityName, err)
	}

	// for case soft deleted (mean: can't retrieve record when status == 0)
	if result.Status == 0 {
		return nil, common.ErrEntityDeleted(categorymodel.EntityName, err)
	}
	return result, err
}
