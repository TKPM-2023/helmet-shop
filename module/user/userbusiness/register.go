package userbusiness

import (
	"TKPM-Go/common"
	"TKPM-Go/module/cart/cartmodel"
	"TKPM-Go/module/user/usermodel"
	"TKPM-Go/pubsub"
	"context"
	"fmt"
)

type RegisterStorage interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type CreatCartStorage interface {
	CreateCart(ctx context.Context, data *cartmodel.CartCreate) (int, error)
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage   RegisterStorage
	createCartStorage CreatCartStorage
	hasher            Hasher
	pubsub            pubsub.Pubsub
}

func NewRegisterBusiness(registerStorage RegisterStorage, createCartStorage CreatCartStorage, hasher Hasher, pubsub pubsub.Pubsub) *registerBusiness {
	return &registerBusiness{
		registerStorage:   registerStorage,
		createCartStorage: createCartStorage,
		hasher:            hasher,
		pubsub:            pubsub,
	}
}

func (biz *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	user, _ := biz.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = biz.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"

	cartId, err := biz.createCartStorage.CreateCart(ctx, &cartmodel.CartCreate{TotalProduct: 0})
	if err != nil {
		return common.ErrCannotCreateEntity(cartmodel.EntityName, err)
	}

	fmt.Println(cartId)

	data.CartId = cartId

	if err := biz.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}
