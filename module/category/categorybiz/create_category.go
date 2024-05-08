package categorybiz

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
)

type CreateCategoryStore interface {
	CreateCategory(ctx context.Context, data *categorymodel.CategoryCreate) error
}

type createCategoryBusiness struct {
	store CreateCategoryStore
}

func NewCreateCategoryBusiness(store CreateCategoryStore) *createCategoryBusiness {
	return &createCategoryBusiness{store: store}
}

func (business *createCategoryBusiness) CreateCategory(
	context context.Context,
	data *categorymodel.CategoryCreate,
) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.store.CreateCategory(context, data); err != nil {
		return common.ErrCannotCreateEntity(categorymodel.EntityName, err)
	}
	return nil
}
