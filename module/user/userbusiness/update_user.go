package userbusiness

import (
	"TKPM-Go/common"
	"TKPM-Go/module/user/usermodel"
	"context"
)

type UpdateUserInfoStore interface {
	UpdateUserInfo(context context.Context, id int, data *usermodel.UserUpdate) error
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type updateUserInfoBusiness struct {
	store UpdateUserInfoStore
}

func NewUpdateInfoUserBusiness(store UpdateUserInfoStore) *updateUserInfoBusiness {
	return &updateUserInfoBusiness{store: store}
}

func (business *updateUserInfoBusiness) UpdateUserInfo(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindUser(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return err
	}

	if result != nil && result.Status == 0 {
		return common.ErrEntityNotFound(usermodel.EntityName, nil)
	}

	if data.Email != "" {
		check, err := business.store.FindUser(ctx, map[string]interface{}{
			"email": data.Email,
		})

		if err != nil {
			return err
		}

		if check != nil {
			return usermodel.ErrEmailExisted
		}
	}

	if err := business.store.UpdateUserInfo(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
