package userbusiness

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/user/usermodel"
)

type UpdatePasswordStore interface {
	UpdatePassword(context context.Context, id int, data *usermodel.PasswordUpdate) error
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}

type updatePasswordStore struct {
	store  UpdatePasswordStore
	hasher Hasher
}

func NewUpdatePasswordStore(store UpdatePasswordStore, hasher Hasher) *updatePasswordStore {
	return &updatePasswordStore{store: store, hasher: hasher}
}

func (business *updatePasswordStore) UpdatePassword(ctx context.Context, id int, data *usermodel.PasswordUpdate) error {
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

	passwordHashed := business.hasher.Hash(data.Password + result.Salt)

	if result.Password != passwordHashed {
		return usermodel.PasswordIncorrect
	}

	data.NewPassword = business.hasher.Hash(data.NewPassword + result.Salt)

	if err := business.store.UpdatePassword(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}
	return nil
}
