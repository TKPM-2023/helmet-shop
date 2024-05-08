package ordermodel

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
)

var (
	ErrOrderUserIdIsRequired = common.NewCustomError(
		errors.New("order UserId is required"),
		"order UserId is required",
		"ErrOrderUserIdIsRequired")

	ErrOrderTotalPriceIsRequired = common.NewCustomError(
		errors.New("order total price is required"),
		"order total price is required",
		"ErrOrderTotalPriceIsRequired")

	ErrOrderContactIdIsRequired = common.NewCustomError(
		errors.New("order contact is required"),
		"order contact is required",
		"ErrOrderContactIdIsRequired")

	ErrOrderStatusInvalid = common.NewCustomError(
		errors.New("order status invalid"),
		"order status invalid",
		"ErrOrderStatusInvalid")
)
