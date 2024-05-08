package ordermodel

import (
	"github.com/go-playground/validator/v10"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
	"github.com/orgball2608/helmet-shop-be/module/user/usermodel"
)

const EntityName = "Orders"

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int                            `json:"-" gorm:"column:user_id"`
	UserUID         *common.UID                    `json:"user_id" gorm:"-"`
	User            *usermodel.User                `json:"user"`
	TotalPrice      float64                        `json:"total_price" gorm:"column:total_price"`
	OrderStatus     int                            `json:"order_status" gorm:"column:order_status;default:1"`
	Products        []orderdetailmodel.OrderDetail `json:"products"`
	ContactUID      *common.UID                    `json:"contact_id" gorm:"-"`
	ContactId       int                            `json:"-" gorm:"column:contact_id"`
	Contact         *contactmodel.Contact          `json:"contact"`
}

func (Order) TableName() string {
	return "orders"
}

func (c *Order) Mask() {
	c.GenUID(common.DbTypeOrder)

	contactUID := common.NewUID(uint32(c.ContactId), int(common.DbTypeContact), 1)
	c.ContactUID = &contactUID

	userUID := common.NewUID(uint32(c.UserId), int(common.DbTypeUser), 1)
	c.UserUID = &userUID
}

func (c *Order) GetOrderID() int {
	return c.Id
}

type OrderCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int                                  `json:"-" validate:"required" gorm:"column:user_id"`
	UserUID         *common.UID                          `json:"user_id" gorm:"-"`
	TotalPrice      float64                              `json:"total_price" gorm:"column:total_price"`
	OrderStatus     int                                  `json:"-" gorm:"column:order_status;default:1"`
	ContactUID      *common.UID                          `json:"contact_id" gorm:"-"`
	ContactId       int                                  `json:"-" gorm:"column:contact_id"`
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

	if err := validate.Var(res.UserId, "required"); err != nil {
		return ErrOrderUserIdIsRequired
	}

	if err := validate.Var(res.ContactId, "required"); err != nil {
		return ErrOrderContactIdIsRequired
	}
	return nil
}

type OrderUpdate struct {
	common.SQLModel `json:",inline"`
	UserId          int         `json:"-" validate:"required" gorm:"column:user_id"`
	UserUID         *common.UID `json:"user_id" gorm:"-"`
	TotalPrice      float64     `json:"total_price" validate:"required" gorm:"column:total_price"`
	OrderStatus     int         `json:"order_status" gorm:"column:order_status;default:1"`
	ContactUID      *common.UID `json:"contact_id" gorm:"-"`
	ContactId       int         `json:"-" gorm:"column:contact_id"`
}

func (res *OrderUpdate) Validate() error {
	validate := validator.New()
	if err := validate.Var(res.OrderStatus, "omitempty,min=1,max=4"); err != nil {
		return ErrOrderStatusInvalid
	}
	return nil
}

func (OrderUpdate) TableName() string {
	return Order{}.TableName()
}
