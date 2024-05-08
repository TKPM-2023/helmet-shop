package ratingmodel

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
)

var (
	ErrRatingPointIsRequired = common.NewCustomError(
		errors.New("rating point is required"),
		"rating point is required",
		"ErrRatingPointIsRequired")

	ErrRatingPointIsInvalid = common.NewCustomError(
		errors.New("rating point is invalid"),
		"rating point is invalid",
		"ErrRatingPointIsInvalid")

	ErrCommentIsRequired = common.NewCustomError(
		errors.New("comment is required"),
		"comment is required",
		"ErrCommentIsRequired")

	ErrCommentIsInvalid = common.NewCustomError(
		errors.New("comment is invalid"),
		"comment is invalid",
		"ErrCommentIsInvalid")

	ErrUserIdIsRequired = common.NewCustomError(
		errors.New("UserId is required"),
		"UserId is required",
		"ErrUserIdIsRequired")

	ErrProductIdIsRequired = common.NewCustomError(
		errors.New("ProductId is required"),
		"ProductId is required",
		"ErrProductIdIsRequired")

	ErrCannotCreateRating = common.NewCustomError(
		errors.New("cannot create rating"),
		"cannot create rating",
		"ErrCannotCreateRating")
)
