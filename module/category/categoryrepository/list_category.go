package categoryrepository

import (
	"TKPM-Go/common"
	"TKPM-Go/module/category/categorymodel"
	"context"
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
