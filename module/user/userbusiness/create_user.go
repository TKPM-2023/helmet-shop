package userbusiness

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/user/usermodel"
)

type CreateUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type createUserBusiness struct {
	createUserStore CreateUserStore
	hasher          Hasher
}

func NewCreateUserBusiness(createUserStore CreateUserStore, hasher Hasher) *createUserBusiness {
	return &createUserBusiness{
		createUserStore: createUserStore,
		hasher:          hasher,
	}
}

func (biz *createUserBusiness) CreateUser(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, _ := biz.createUserStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt

	if err := biz.createUserStore.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
