package categorymodel

import (
	"LearnGo/common"
	"github.com/go-playground/validator/v10"
)

const EntityName = "Categories"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description" gorm:"column:description;"`
	Icon            *common.Image `json:"icon" gorm:"column:icon;"`
	TotalProduct    int           `json:"total_product" gorm:"column:total_product;"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Mask() {
	c.GenUID(common.DbTypeUser)
}

func (c *Category) GetCategoryID() int {
	return c.Id
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" validate:"required" gorm:"column:name;"`
	Description     string        `json:"description" validate:"required" gorm:"column:description;"`
	Icon            *common.Image `json:"icon" validate:"required" gorm:"column:icon;"`
}

func (res *CategoryCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Name, "required"); err != nil {
		return ErrCategoryNameIsRequired
	}

	if err := validate.Var(res.Description, "required"); err != nil {
		return ErrCategoryDescriptionIsRequired
	}

	if err := validate.Var(res.Icon, "required"); err != nil {
		return ErrCategoryIconIsRequired
	}

	return nil
}

func (CategoryCreate) TableName() string {
	return Category{}.TableName()
}

func (c *CategoryCreate) Mask() {
	c.GenUID(common.DbTypeUser)
}
