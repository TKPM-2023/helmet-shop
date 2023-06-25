package orderdetailmodel

import (
	"TKPM-Go/common"
	"TKPM-Go/module/contact/contactmodel"
	"TKPM-Go/module/user/usermodel"
)

type Order struct {
	common.SQLModel `json:",inline"`
	UserId          int                   `json:"-" gorm:"column:user_id"`
	UserUID         *common.UID           `json:"user_id" gorm:"-"`
	User            *usermodel.User       `json:"user"`
	TotalPrice      float64               `json:"total_price" gorm:"column:total_price"`
	OrderStatus     int                   `json:"order_status" gorm:"column:order_status;default:1"`
	ContactUID      *common.UID           `json:"contact_id" gorm:"-"`
	ContactId       int                   `json:"-" gorm:"column:contact_id"`
	Contact         *contactmodel.Contact `json:"contact"`
}

func (Order) TableName() string {
	return "orders"
}
