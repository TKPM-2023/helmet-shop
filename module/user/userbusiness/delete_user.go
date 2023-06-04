package userbusiness

import (
	"TKPM-Go/common"
	"TKPM-Go/module/user/usermodel"
	"context"
)

type DeleteUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
	DeleteUser(context context.Context, id int) error
}

type deleteUserBusiness struct {
	store DeleteUserStore
}

func NewDeleteUserBusiness(store DeleteUserStore) *deleteUserBusiness {
	return &deleteUserBusiness{store: store}
}

func (business *deleteUserBusiness) DeleteUser(context context.Context, id int) error {
	oldData, err := business.store.FindUser(context, map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, err)
	}

	if err := business.store.DeleteUser(context, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, err)
	}
	return nil
}
