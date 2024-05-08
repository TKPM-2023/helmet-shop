package categorymodel

import (
	"github.com/go-playground/validator/v10"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
)

const EntityName = "Categories"

type Category struct {
	common.SQLModel `json:",inline"`
	Name            string                 `json:"name" gorm:"column:name;"`
	Description     string                 `json:"description" gorm:"column:description;"`
	Icon            *common.Image          `json:"icon" gorm:"column:icon;"`
	TotalProduct    int                    `json:"total_product" gorm:"column:total_product;"`
	Products        []productmodel.Product `json:"products"`
}

func (Category) TableName() string {
	return "categories"
}

func (c *Category) Mask() {
	c.GenUID(common.DbTypeCategory)
}

func (c *Category) GetCategoryID() int {
	return c.Id
}

type CategoryCreate struct {
	common.SQLModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	Description     string        `json:"description"  gorm:"column:description;"`
	Icon            *common.Image `json:"icon" gorm:"column:icon;"`
}

func (res *CategoryCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Name, "required"); err != nil {
		return ErrCategoryNameIsRequired
	}

	if err := validate.Var(res.Name, "min=5"); err != nil {
		return ErrCategoryNameLengthIsInvalid
	}

	if err := validate.Var(res.Description, "required"); err != nil {
		return ErrCategoryDescriptionIsRequired
	}

	if err := validate.Var(res.Description, "min=5"); err != nil {
		return ErrCategoryDescriptionLengthIsInvalid
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

type CategoryUpdate struct {
	Name        string        `json:"name" gorm:"column:name;"`
	Description string        `json:"description" gorm:"column:description;"`
	Icon        *common.Image `json:"icon" gorm:"column:icon;"`
}

func (CategoryUpdate) TableName() string {
	return Category{}.TableName()
}

func (res *CategoryUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Name, "omitempty,min=5,max=100"); err != nil {
		return ErrCategoryNameLengthIsInvalid
	}

	if err := validate.Var(res.Description, "omitempty,min=5,max=100"); err != nil {
		return ErrCategoryDescriptionLengthIsInvalid
	}

	return nil
}
