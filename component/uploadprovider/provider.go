package uploadprovider

import (
	"TKPM-Go/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
