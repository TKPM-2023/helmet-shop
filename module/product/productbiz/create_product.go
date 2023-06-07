package productbiz

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
	"TKPM-Go/pubsub"
	"context"
)

type CreateProductStore interface {
	CreateCategory(ctx context.Context, data *productmodel.ProductCreate) error
	FindProductWithCondition(ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*productmodel.Product, error)
}

type createProductBusiness struct {
	store  CreateProductStore
	pubsub pubsub.Pubsub
}

func NewCreateProductBusiness(store CreateProductStore, pubsub pubsub.Pubsub) *createProductBusiness {
	return &createProductBusiness{store: store, pubsub: pubsub}
}

func (business *createProductBusiness) CreateProduct(context context.Context, data *productmodel.ProductCreate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	result, err := business.store.FindProductWithCondition(context, map[string]interface{}{"name": data.Name})
	if err != nil {
		return err
	}
	if result != nil {
		return common.ErrEntityExisted(productmodel.EntityName, nil)
	}

	if err := business.store.CreateCategory(context, data); err != nil {
		return common.ErrCannotCreateEntity(productmodel.EntityName, err)
	}

	business.pubsub.Publish(context, common.TopicUserAddProduct, pubsub.NewMessage(data))

	return nil

}
