package contactmodel

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
)

var (
	ErrContactUserIdIsRequired = common.NewCustomError(
		errors.New("user id is required"),
		"User id is required",
		"ErrContactUserIdIsRequired")

	ErrContactNameIsRequired = common.NewCustomError(
		errors.New("name is required"),
		"Name is required",
		"ErrContactNameIsRequired")

	ErrContactAddressIsRequired = common.NewCustomError(
		errors.New("address is required"),
		"Address is required",
		"ErrContactAddressIsRequired")

	ErrContactPhoneIsRequired = common.NewCustomError(
		errors.New("phone is required"),
		"Phone is required",
		"ErrContactPhoneIsRequired")
)
