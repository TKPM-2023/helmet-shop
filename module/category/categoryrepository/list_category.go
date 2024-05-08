package categoryrepository

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
)

type ListCategoryStorage interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *categorymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]categorymodel.Category, error)
}

type listCategoryRepo struct {
	store ListCategoryStorage
}

func NewListCategoryRepo(store ListCategoryStorage) *listCategoryRepo {
	return &listCategoryRepo{store: store}
}

func (repo *listCategoryRepo) ListCategory(
	context context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
) ([]categorymodel.Category, error) {
	result, err := repo.store.ListDataWithCondition(context, filter, paging, "Products")
	if err != nil {
		return nil, err
	}

	return result, nil
}
