package orderdetailmodel

import (
	"TKPM-Go/common"
	"errors"
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
)