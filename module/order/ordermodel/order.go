package ordermodel

import (
	"TKPM-Go/common"

	"TKPM-Go/module/order_detail/orderdetailmodel"

	"github.com/go-playground/validator/v10"
)

const EntityName = "Orders"

type Order struct {
	common.SQLModel `json:",inline"`
	User_ID         int                            `json:"user_id" gorm:"column:user_id"`
	Total_Price     float64                        `json:"total_price" gorm:"column:total_price"`
	Products        []orderdetailmodel.OrderDetail `json:"products"`
}

func (Order) TableName() string {
	return "orders"
}

func (c *Order) Mask() {
	c.GenUID(common.DbTypeOrder)
}

func (c *Order) GetOrderID() int {
	return c.Id
}

type OrderCreate struct {
	common.SQLModel `json:",inline"`
	User_ID         int     `json:"user_id" validate:"required" gorm:"column:user_id"`
	Total_Price     float64 `json:"total_price" validate:"required" gorm:"column:total_price"`
}

func (OrderCreate) TableName() string {
	return Order{}.TableName()
}

func (c *OrderCreate) Mask() {
	c.GenUID(common.DbTypeOrder)
}

func (res *OrderCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.User_ID, "required"); err != nil {
		return ErrOrderUserIdIsRequired
	}

	if err := validate.Var(res.Total_Price, "required"); err != nil {
		return ErrOrderTotalPriceIsRequired
	}
	return nil
}

type OrderUpdate struct {
	common.SQLModel `json:",inline"`
	User_ID         int     `json:"user_id" validate:"required" gorm:"column:user_id"`
	Total_Price     float64 `json:"total_price" validate:"required" gorm:"column:total_price"`
}

func (res *OrderUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.User_ID, "required"); err != nil {
		return ErrOrderUserIdIsRequired
	}

	if err := validate.Var(res.Total_Price, "required"); err != nil {
		return ErrOrderTotalPriceIsRequired
	}
	return nil
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
