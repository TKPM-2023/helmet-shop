package productmodel

import (
	"TKPM-Go/common"
	"errors"
)

var (
	ErrProductNameExisted = common.NewCustomError(
		errors.New("category name has already exits"),
		"category name has already exits",
		"ErrCategoryNameExisted")

	ErrProductNameIsRequired = common.NewCustomError(
		errors.New("category name is required"),
		"category name is required",
		"ErrCategoryNameIsRequired")

	ErrProductNameLengthIsInvalid = common.NewCustomError(
		errors.New("product name length is invalid"),
		"product name length is invalid",
		"ErrProductNameLengthIsInvalid")

	ErrProductDescriptionIsRequired = common.NewCustomError(
		errors.New("category description is required"),
		"category description is required",
		"ErrCategoryDescriptionIsRequired")

	ErrProductDescriptionLengthIsInvalid = common.NewCustomError(
		errors.New("product description is invalid"),
		"product description is invalid",
		"ErrProductDescriptionLengthIsInvalid")

	ErrProductPriceIsRequired = common.NewCustomError(
		errors.New("product price is required"),
		"product price is required",
		"ErrProductPriceIsRequired")

	ErrProductPriceMustBeGreaterThanZero = common.NewCustomError(
		errors.New("product price must be greater than zero"),
		"product price must be greater than zero",
		"ErrProductPriceMustBeGreaterThanZero")

	ErrProductQuantityIsRequired = common.NewCustomError(
		errors.New("product quantity is required"),
		"product quantity is required",
		"ErrProductQuantityIsRequired")

	ErrProductQuantityMustBeAtLeastOne = common.NewCustomError(
		errors.New("product quantity must be at least one"),
		"product quantity must be at least one",
		"ErrProductQuantityMustBeAtLeastOne")

	ErrProductQuantityMustBeAtLeastZero = common.NewCustomError(
		errors.New("product quantity must be at least zero"),
		"product quantity must be at least zero",
		"ErrProductQuantityMustBeAtLeastZero")

	ErrProductImagesIsRequired = common.NewCustomError(
		errors.New("product images is required"),
		"product images is required",
		"ErrProductImagesIsRequired")

	ErrProductCategoryIdIsRequired = common.NewCustomError(
		errors.New("product CategoryId is required"),
		"product CategoryId is required",
		"ErrProductCategoryIdIsRequired")

	ErrProductCategoryIdMustBeAtLeastOne = common.NewCustomError(
		errors.New("product CategoryId must be at least one"),
		"product CategoryId must be at least one",
		"ErrProductCategoryIdMustBeAtLeastOne")
)
