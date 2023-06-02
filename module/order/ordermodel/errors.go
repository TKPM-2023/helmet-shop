package ordermodel

import (
	"TKPM-Go/common"
	"errors"
)

var (
	ErrOrderUserIdIsRequired = common.NewCustomError(
		errors.New("order UderId is required"),
		"order UserId is required",
		"ErrOrderUserIdIsRequired")

	ErrOrderTotalPriceIsRequired = common.NewCustomError(
		errors.New("order total price is required"),
		"order total price is required",
		"ErrOrderTotalPriceIsRequired")
)
