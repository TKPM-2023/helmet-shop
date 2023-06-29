package ratingmodel

import "TKPM-Go/common"

type Product struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Description     string         `json:"description" gorm:"column:description;"`
	Price           int            `json:"price" gorm:"column:price;"`
	Images          *common.Images `json:"images" gorm:"column:images;"`
}

func (Product) TableName() string {
	return "products"
}
