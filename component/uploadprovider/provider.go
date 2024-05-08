package uploadprovider

import (
	"context"
	"github.com/orgball2608/helmet-shop-be/common"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
