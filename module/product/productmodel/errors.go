package productmodel

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
)

var (
	ErrProductNameExisted = common.NewCustomError(
		errors.New("product name has already exits"),
		"product name has already exits",
		"ErrProductNameExisted")

	ErrProductNameIsRequired = common.NewCustomError(
		errors.New("product name is required"),
		"product name is required",
		"ErrProductNameIsRequired")

	ErrProductNameLengthIsInvalid = common.NewCustomError(
		errors.New("product name length is invalid"),
		"product name length is invalid",
		"ErrProductNameLengthIsInvalid")

	ErrProductDescriptionIsRequired = common.NewCustomError(
		errors.New("category description is required"),
		"category description is required",
		"ErrCategoryDescriptionIsRequired")

	ErrProductDescriptionLengthIsInvalid = common.NewCustomError(
		errors.New("product description length is invalid"),
		"product description length is invalid",
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
