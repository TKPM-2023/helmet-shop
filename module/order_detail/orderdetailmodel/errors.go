package orderdetailmodel

import (
	"errors"
	"github.com/orgball2608/helmet-shop-be/common"
)

var (
	ErrOrderDetailOrderIdIsRequired = common.NewCustomError(
		errors.New("OrderId is required"),
		"OrderId is required",
		"ErrOrderDetailOrderIdIsRequired")

	ErrOrderDetailPriceIsRequired = common.NewCustomError(
		errors.New("product price for order is required"),
		"order total price is required",
		"ErrOrderDetailPriceIsRequired")

	ErrOrderDetailQuantityIsRequired = common.NewCustomError(
		errors.New("product quantity for order is required"),
		"product quantity for order is required",
		"ErrOrderDetailQuantityIsRequired")

	ErrOrderDetailOrderIDNotFound = common.NewCustomError(
		errors.New("order id not found"),
		"order id not found",
		"ErrOrderDetailOrderIDNotFound")
)
