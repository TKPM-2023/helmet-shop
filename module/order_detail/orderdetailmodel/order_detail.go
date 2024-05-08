package orderdetailmodel

import (
	"github.com/go-playground/validator/v10"
	"github.com/orgball2608/helmet-shop-be/common"
)

const EntityName = "Order Details"

type OrderDetail struct {
	common.SQLModel `json:",inline"`
	OrderUID        *common.UID    `json:"order_id" gorm:"-"`
	OrderId         int            `json:"-" gorm:"column:order_id"`
	ProductOrigin   *ProductOrigin `json:"product_origin" gorm:"product_origin"`
	Price           float64        `json:"price" gorm:"column:price"`
	Quantity        int            `json:"quantity" gorm:"column:quantity"`
	Discount        float32        `json:"discount" gorm:"column:discount"`
	Order           *Order         `json:"order" gorm:"foreignKey:OrderId"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}

func (d *OrderDetail) Mask() {
	d.GenUID(common.DbTypeOrderDetail)

	uid := common.NewUID(uint32(d.OrderId), int(common.DbTypeOrder), 1)
	d.OrderUID = &uid
}

func (d *OrderDetail) GetOrderDetailID() int {
	return d.Id
}

type OrderDetailCreate struct {
	common.SQLModel `json:",inline"`
	OrderId         int            `json:"-" validate:"required" gorm:"column:order_id"`
	OrderUID        *common.UID    `json:"order_id" gorm:"-"`
	ProductOrigin   *ProductOrigin `json:"product_origin" gorm:"product_origin"`
	Price           float64        `json:"price" validate:"required" gorm:"column:price"`
	Quantity        int            `json:"quantity" validate:"required" gorm:"column:quantity"`
	Discount        float32        `json:"discount" gorm:"column:discount"`
}

func (res *OrderDetailCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.OrderId, "required"); err != nil {
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

func (d *OrderDetailCreate) Mask() {
	d.GenUID(common.DbTypeUser)
}

type OrderDetailUpdate struct {
	OrderId       int            `json:"-" gorm:"column:order_id"`
	OrderUID      *common.UID    `json:"order_id" gorm:"-"`
	ProductOrigin *ProductOrigin `json:"product_origin" gorm:"product_origin"`
	Price         float64        `json:"price" gorm:"column:price"`
	Quantity      int            `json:"quantity" gorm:"column:quantity"`
	Discount      float32        `json:"discount" gorm:"column:discount"`
}

func (OrderDetailUpdate) TableName() string {
	return OrderDetail{}.TableName()
}

func (res *OrderDetailUpdate) Validate() error {

	validate := validator.New()

	if err := validate.Var(res.OrderId, "required"); err != nil {
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
