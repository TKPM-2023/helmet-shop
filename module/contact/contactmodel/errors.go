package contactmodel

import (
	"TKPM-Go/common"
	"errors"
)

var (
	ErrContactUserIdIsRequired = common.NewCustomError(
		errors.New("User id is required"),
		"User id is required",
		"ErrContactUserIdIsRequired")

	ErrContactNameIsRequired = common.NewCustomError(
		errors.New("Name is required"),
		"Name is required",
		"ErrContactNameIsRequired")

	ErrContactAddressIsRequired = common.NewCustomError(
		errors.New("Address is required"),
		"Address is required",
		"ErrContacAddressIsRequired")

	ErrContactPhoneIsRequired = common.NewCustomError(
		errors.New("Phone is required"),
		"Phone is required",
		"ErrContactPhoneIsRequired")

)
