package categorybiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/category/categorymodel"
	"context"
)

type ListCategoryRepo interface {
	ListCategory(
		context context.Context,
		filter *categorymodel.Filter,
		paging *common.Paging,
	) ([]categorymodel.Category, error)
}

type listCategoryBusiness struct {
	repo ListCategoryRepo
}

func NewListCategoryBusiness(repo ListCategoryRepo) *listCategoryBusiness {
	return &listCategoryBusiness{repo: repo}
}

func (business *listCategoryBusiness) ListCategory(context context.Context,
	filter *categorymodel.Filter,
	paging *common.Paging,
) ([]categorymodel.Category, error) {
	result, err := business.repo.ListCategory(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
