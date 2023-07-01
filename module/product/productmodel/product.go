package productmodel

import (
	"TKPM-Go/common"
	"TKPM-Go/module/product_rating/ratingmodel"
	"github.com/go-playground/validator/v10"
)

const EntityName = "Products"

type Product struct {
	common.SQLModel `json:",inline"`
	Name            string               `json:"name" gorm:"column:name;"`
	Description     string               `json:"description" gorm:"column:description;"`
	Price           int                  `json:"price" gorm:"column:price;"`
	Quantity        int                  `json:"quantity" gorm:"column:quantity;"`
	Images          *common.Images       `json:"images" gorm:"column:images;"`
	TotalRating     int                  `json:"total_rating" gorm:"column:total_rating;"`
	CategoryId      int                  `json:"-" gorm:"column:category_id"`
	CategoryUID     *common.UID          `json:"category_id" gorm:"-"`
	Ratings         []ratingmodel.Rating `json:"ratings"`
}

func (Product) TableName() string {
	return "products"
}

func (p *Product) Mask() {
	p.GenUID(common.DbTypeProduct)

	uid := common.NewUID(uint32(p.CategoryId), int(common.DbTypeCategory), 1)
	p.CategoryUID = &uid
}

func (p *Product) GetProductID() int {
	return p.Id
}

type ProductCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Description     string         `json:"description" gorm:"column:description;"`
	Price           int            `json:"price" gorm:"column:price;"`
	Quantity        int            `json:"quantity" gorm:"column:quantity;"`
	Images          *common.Images `json:"images"  gorm:"column:images;"`
	CategoryId      int            `json:"-"  gorm:"column:category_id"`
	CategoryUID     *common.UID    `json:"category_id" gorm:"-"`
}

func (ProductCreate) TableName() string {
	return Product{}.TableName()
}

func (p *ProductCreate) Mask() {
	p.GenUID(common.DbTypeProduct)
}

func (p *ProductCreate) GetCategoryID() int {
	return p.CategoryId
}

func (p *ProductCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(p.Name, "required"); err != nil {
		return ErrProductNameIsRequired
	}

	if err := validate.Var(p.Name, "min=5,max=200"); err != nil {
		return ErrProductNameLengthIsInvalid
	}

	if err := validate.Var(p.Description, "required"); err != nil {
		return ErrProductDescriptionIsRequired
	}

	if err := validate.Var(p.Description, "min=5,max=2000"); err != nil {
		return ErrProductDescriptionLengthIsInvalid
	}

	if err := validate.Var(p.Price, "required,gt=0"); err != nil {
		return ErrProductPriceIsRequired
	}

	if err := validate.Var(p.Price, "gt=0"); err != nil {
		return ErrProductPriceMustBeGreaterThanZero
	}

	if err := validate.Var(p.Quantity, "required"); err != nil {
		return ErrProductQuantityIsRequired
	}

	if err := validate.Var(p.Quantity, "min=1"); err != nil {
		return ErrProductQuantityMustBeAtLeastOne
	}

	if err := validate.Var(p.Images, "required"); err != nil {
		return ErrProductImagesIsRequired
	}

	if err := validate.Var(p.CategoryId, "required"); err != nil {
		return ErrProductCategoryIdIsRequired
	}

	if err := validate.Var(p.CategoryId, "min=1"); err != nil {
		return ErrProductCategoryIdMustBeAtLeastOne
	}

	return nil
}

type ProductUpdate struct {
	Name        string         `json:"name" gorm:"column:name;"`
	Description string         `json:"description" gorm:"column:description;"`
	Price       int            `json:"price" gorm:"column:price;"`
	Quantity    int            `json:"quantity" gorm:"column:quantity;"`
	Images      *common.Images `json:"images" gorm:"column:images;"`
	CategoryId  int            `json:"-" gorm:"column:category_id"`
	CategoryUID *common.UID    `json:"category_id" gorm:"-"`
}

func (res *ProductUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Name, "omitempty,min=5,max=100"); err != nil {
		return ErrProductNameLengthIsInvalid
	}

	if err := validate.Var(res.Description, "omitempty,min=5,max=100"); err != nil {
		return ErrProductDescriptionLengthIsInvalid
	}

	if err := validate.Var(res.Price, "omitempty,gt=0"); err != nil {
		return ErrProductPriceMustBeGreaterThanZero
	}

	if err := validate.Var(res.Quantity, "omitempty,min=0"); err != nil {
		return ErrProductQuantityMustBeAtLeastZero
	}

	return nil
}

func (ProductUpdate) TableName() string {
	return Product{}.TableName()
}
