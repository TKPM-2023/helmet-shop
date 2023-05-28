package categorybiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/category/categorymodel"
	"context"
)

type DeleteCategoryStore interface {
	DeleteCategory(context context.Context, id int) error
	FindCategoryWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*categorymodel.Category, error)
}

type deleteCategoryBusiness struct {
	store     DeleteCategoryStore
	requester common.Requester
}

func NewDeleteCategoryBusiness(store DeleteCategoryStore, requester common.Requester) *deleteCategoryBusiness {
	return &deleteCategoryBusiness{store: store, requester: requester}
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

	if oldData.Id != business.requester.GetUserId() {
		return common.ErrNoPermission(nil)
	}

	if err := business.store.DeleteCategory(context, id); err != nil {
		return common.ErrCannotDeleteEntity(categorymodel.EntityName, err)
	}
	return nil
}
