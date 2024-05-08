package categorymodel

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
)

var (
	ErrCategoryNameExisted = common.NewCustomError(
		errors.New("category name has already exits"),
		"category name has already exits",
		"ErrCategoryNameExisted")

	ErrCategoryNameIsRequired = common.NewCustomError(
		errors.New("category name is required"),
		"category name is required",
		"ErrCategoryNameIsRequired")

	ErrCategoryNameLengthIsInvalid = common.NewCustomError(
		errors.New("category name length is invalid"),
		"category name length is invalid",
		"ErrCategoryNameLengthIsInvalid")

	ErrCategoryDescriptionIsRequired = common.NewCustomError(
		errors.New("category description is required"),
		"category description is required",
		"ErrCategoryDescriptionIsRequired")

	ErrCategoryDescriptionLengthIsInvalid = common.NewCustomError(
		errors.New("category description length is invalid"),
		"category description length is invalid",
		"ErrCategoryDescriptionLengthIsInvalid")

	ErrCategoryIconIsRequired = common.NewCustomError(
		errors.New("category icon is required"),
		"category icon is required",
		"ErrCategoryIconIsRequired")
)
