package ratingmodel

import (
	"TKPM-Go/common"
	"TKPM-Go/module/user/usermodel"
	"github.com/go-playground/validator/v10"
)

const EntityName = "ProductRatings"

type Rating struct {
	common.SQLModel `json:",inline"`
	Point           float32         `json:"point" gorm:"column:point"`
	Comment         string          `json:"comment" gorm:"column:comment;"`
	UserID          int             `json:"-" gorm:"column:user_id"`
	UserUID         *common.UID     `json:"user_id" gorm:"-"`
	ProductID       int             `json:"-" gorm:"column:product_id"`
	ProductUID      *common.UID     `json:"product_id" gorm:"-"`
	User            *usermodel.User `gorm:"foreignKey:UserID"`
}

func (r *Rating) GenUserUID() {
	uid := common.NewUID(uint32(r.UserID), int(common.DbTypeUser), 1)
	r.UserUID = &uid
}

func (r *Rating) GenProductUID() {
	uid := common.NewUID(uint32(r.ProductID), int(common.DbTypeProduct), 1)
	r.ProductUID = &uid
}

func (Rating) TableName() string {
	return "product_ratings"
}

func (r *Rating) Mask() {
	r.GenUID(common.DbTypeProductRating)
}

func (r *Rating) GetProductID() int {
	return r.ProductID
}

type RatingCreate struct {
	common.SQLModel `json:",inline"`
	Point           float32            `json:"point" gorm:"column:point"`
	Comment         string             `json:"comment" gorm:"column:comment;"`
	UserID          int                `json:"-" gorm:"column:user_id"`
	User            *common.SimpleUser `json:"user"`
	ProductID       int                `json:"-" gorm:"column:product_id"`
}

func (RatingCreate) TableName() string {
	return Rating{}.TableName()
}

func (r *RatingCreate) Mask() {
	r.GenUID(common.DbTypeProductRating)
}

func (c *RatingCreate) GetUserID() int {
	return c.UserID
}

func (c *RatingCreate) GetProductID() int {
	return c.ProductID
}

func (res *RatingCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Point, "required"); err != nil {
		return ErrRatingPointIsRequired
	}

	if err := validate.Var(res.Point, "min=0,max=5"); err != nil {
		return ErrRatingPointIsInvalid
	}

	if err := validate.Var(res.Comment, "required"); err != nil {
		return ErrCommentIsRequired
	}

	if err := validate.Var(res.UserID, "required"); err != nil {
		return ErrUserIdIsRequired
	}

	if err := validate.Var(res.ProductID, "required"); err != nil {
		return ErrProductIdIsRequired
	}

	return nil
}

type RatingUpdate struct {
	Point   float32 `json:"point" gorm:"column:point"`
	Comment string  `json:"comment" gorm:"column:comment;"`
}

func (RatingUpdate) TableName() string {
	return Rating{}.TableName()
}

func (res *RatingUpdate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Point, "omitempty,min=0,max=5"); err != nil {
		return ErrRatingPointIsInvalid
	}

	return nil
}
