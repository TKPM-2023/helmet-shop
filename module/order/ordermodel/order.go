package ordermodel

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"TKPM-Go/module/order_detail/orderdetailmodel"

	"github.com/go-playground/validator/v10"
)

const EntityName = "Orders"

type Order struct {
	common.SQLModel `json:",inline"`
	User_ID         int                            `json:"-" gorm:"column:user_id"`
	User_UID        *common.UID                    `json:"user_id" gorm:"-"`
	Total_Price     float64                        `json:"total_price" gorm:"column:total_price"`
	Order_Status    string                         `json:"order_status" gorm:"column:order_status;default:chưa xử lý"`
	Products        []orderdetailmodel.OrderDetail `json:"products"`
	Contact_UID     *common.UID                    `json:"contact_id" gorm:"-"`
	Contact_ID      int                            `json:"-" gorm:"column:contact_id"`
	Contact         *contactmodel.Contact          `json:"contact"`
}

func (c *Order) GenUserUID() {
	uid := common.NewUID(uint32(c.User_ID), int(common.DbTypeUser), 1)
	c.User_UID = &uid
}

func (c *Order) GenContactUID() {
	uid := common.NewUID(uint32(c.Contact_ID), int(common.DbTypeContact), 1)
	c.Contact_UID = &uid
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
	User_ID         int         `json:"-" validate:"required" gorm:"column:user_id"`
	User_UID        *common.UID `json:"user_id" gorm:"-"`
	Total_Price     float64     `json:"total_price" gorm:"column:total_price"`
	Order_Status    string      `json:"-" gorm:"column:order_status;default:chưa xử lý"`
	Contact_UID     *common.UID `json:"contact_id" gorm:"-"`
	Contact_ID      int         `json:"-" gorm:"column:contact_id"`
	Products        []orderdetailmodel.OrderDetailCreate `json:"products" gorm:"-"`
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


	if err := validate.Var(res.Contact_ID, "required"); err != nil {
		return ErrOrderContactIdIsRequired
	}
	return nil
}

type OrderUpdate struct {
	common.SQLModel `json:",inline"`
	User_ID         int         `json:"-" validate:"required" gorm:"column:user_id"`
	User_UID        *common.UID `json:"user_id" gorm:"-"`
	Total_Price     float64     `json:"total_price" validate:"required" gorm:"column:total_price"`
	Order_Status    string      `json:"order_status" gorm:"column:order_status;default:chưa xử lý"`
	Contact_UID     *common.UID `json:"contact_id" gorm:"-"`
	Contact_ID      int         `json:"-" gorm:"column:contact_id"`
}

func (res *OrderUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.User_ID, "required"); err != nil {
		return ErrOrderUserIdIsRequired
	}

	if err := validate.Var(res.Total_Price, "required"); err != nil {
		return ErrOrderTotalPriceIsRequired
	}
	if err := validate.Var(res.Contact_ID, "required"); err != nil {
		return ErrOrderContactIdIsRequired
	}
	return nil
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
