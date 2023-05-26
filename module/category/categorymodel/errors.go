package categorymodel

import (
	"LearnGo/common"
	"errors"
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

	ErrCategoryDescriptionIsRequired = common.NewCustomError(
		errors.New("category description is required"),
		"category description is required",
		"ErrCategoryDescriptionIsRequired")

	ErrCategoryIconIsRequired = common.NewCustomError(
		errors.New("category icon is required"),
		"category icon is required",
		"ErrCategoryIconIsRequired")
)
