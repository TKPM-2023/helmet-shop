package uploadmodel

import "github.com/orgball2608/helmet-shop-be/common"

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewErrorResponse(err, "can not save file", err.Error(), "ERROR_CANNOT_SAVE_FILE")
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewErrorResponse(err, "file is not image", err.Error(), "ERROR_FILE_IS_NOT_IMAGE")
}
