package uploadbusiness

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
)

type RemoveImageStore interface {
	DeleteImage(ctx context.Context, imageId int) error
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.Image, error)
}

type removeImageBusiness struct {
	store RemoveImageStore
}

func NewDeleteImageBusiness(store RemoveImageStore) *removeImageBusiness {
	return &removeImageBusiness{store: store}
}

func (business *removeImageBusiness) RemoveImage(context context.Context, id int) error {
	_, err := business.store.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if err := business.store.DeleteImage(context, id); err != nil {
		return common.ErrCannotDeleteEntity("Image", err)
	}
	return nil
}
