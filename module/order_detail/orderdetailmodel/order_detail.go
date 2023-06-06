package orderdetailmodel

import (
	"TKPM-Go/common"

	"github.com/go-playground/validator/v10"
)

const EntityName = "Order Details"

type OrderDetail struct {
	common.SQLModel `json:",inline"`
	Order_UID       *common.UID     `json:"order_id" gorm:"-"`
	Order_ID        int             `json:"-" gorm:"column:order_id"`
	Product_Origin  *Product_Origin `json:"product_origin" gorm:"product_origin"`
	Price           float64         `json:"price" gorm:"column:price"`
	Quantity        int             `json:"quantiy" gorm:"column:quantity"`
	Discount        float32         `json:"discount" gorm:"column:discount"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

func (c *OrderDetail) Mask() {
	c.GenUID(common.DbTypeOrder_Detail)
}

func (c *OrderDetail) GetOrderDetailID() int {
	return c.Id
}

type OrderDetailCreate struct {
	common.SQLModel `json:",inline"`
	Order_ID        int             `json:"-" validate:"required" gorm:"column:order_id"`
	Order_UID       *common.UID     `json:"order_id" gorm:"-"`
	Product_Origin  *Product_Origin `json:"product_origin" gorm:"product_origin"`
	Price           float64         `json:"price" validate:"required" gorm:"column:price"`
	Quantity        int             `json:"quantity" validate:"required" gorm:"column:quantity"`
	Discount        float32         `json:"discount" gorm:"column:discount"`
}

func (res *OrderDetailCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Order_ID, "required"); err != nil {
		return ErrOrderDetailOrderIdIsRequired
	}

	if err := validate.Var(res.Price, "required"); err != nil {
		return ErrOrderDetailPriceIsRequired
	}

	if err := validate.Var(res.Quantity, "required"); err != nil {
		return ErrOrderDetailQuantityIsRequired
	}
	return nil
}

func (OrderDetailCreate) TableName() string {
	return OrderDetail{}.TableName()
}

func (c *OrderDetailCreate) Mask() {
	c.GenUID(common.DbTypeUser)
}

type OrderDetailUpdate struct {
	Order_ID       int             `json:"order_id" gorm:"column:order_id"`
	Product_Origin *Product_Origin `json:"product_origin" gorm:"product_origin"`
	Price          float64         `json:"price" gorm:"column:price"`
	Quantity       int             `json:"quantiy" gorm:"column:quantity"`
	Discount       float32         `json:"discount" gorm:"column:discount"`
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetail{}.TableName()
}

func (res *OrderDetailUpdate) Validate() error {

	// chờ implement, lười quá
	return nil
}
