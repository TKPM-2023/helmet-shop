package cartmodel

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product/productmodel"
)

const EntityName = "Carts"

type Cart struct {
	common.SQLModel `json:",inline"`
	TotalProduct    int           `json:"total_product" gorm:"column:total_product"`
	CartProducts    []CartProduct `json:"cart_products" gorm:"ForeignKey:CartId;AssociationForeignKey:Id"`
}

func (Cart) TableName() string {
	return "carts"
}

func (c *Cart) Mask() {
	c.GenUID(common.DbTypeCart)
}

type CartProduct struct {
	CartId     int                   `json:"-" gorm:"column:cart_id; primary_key"`
	CartUID    *common.UID           `json:"cart_id" gorm:"-"`
	ProductId  int                   `json:"-" gorm:"column:product_id; primary_key"`
	ProductUID *common.UID           `json:"product_id" gorm:"-"`
	Quantity   int                   `json:"quantity" gorm:"column:quantity; not null"`
	Product    *productmodel.Product `gorm:"ForeignKey:ProductId"`
}

func (CartProduct) TableName() string {
	return "cart_products"
}

func (c *CartProduct) Mask() {
	cartUID := common.NewUID(uint32(c.CartId), int(common.DbTypeCart), 1)
	c.CartUID = &cartUID

	productUID := common.NewUID(uint32(c.ProductId), int(common.DbTypeProduct), 1)
	c.ProductUID = &productUID
}

type CartCreate struct {
	common.SQLModel `json:",inline"`
	TotalProduct    int `json:"total_product" gorm:"column:total_product;"`
}

func (CartCreate) TableName() string {
	return Cart{}.TableName()
}

type CartProductDetail struct {
	ProductId  int         `json:"-" gorm:"column:product_id;"`
	ProductUID *common.UID `json:"product_id" gorm:"-"`
	Quantity   int         `json:"quantity" gorm:"column:quantity;"`
}

type RemoveCartProducts []CartProductDetail

func (c *CartProductDetail) GetProductID() int {
	return c.ProductId
}

type CartProductDetails []CartProductDetail

type ProductTotalUpdate struct {
	CartId   int `gorm:"column:cart_id;"`
	Quantity int `gorm:"column:quantity;"`
}

func (c *ProductTotalUpdate) GetCartID() int {
	return c.CartId
}

func (c *ProductTotalUpdate) GetQuantity() int {
	return c.Quantity
}
