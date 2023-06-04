package userbusiness

import (
	"TKPM-Go/common"
	"TKPM-Go/module/user/usermodel"
	"context"
)

type ListUserStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.User, error)
}

type listUserBusiness struct {
	store ListUserStore
}

func NewListUserBusiness(store ListUserStore) *listUserBusiness {
	return &listUserBusiness{store: store}
}

func (business *listUserBusiness) ListUser(context context.Context, filter *usermodel.Filter,
	paging *common.Paging) ([]usermodel.User, error) {
	result, err := business.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, err
	}
	return result, nil
}
