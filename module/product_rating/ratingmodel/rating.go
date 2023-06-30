package ratingmodel

import (
	"TKPM-Go/common"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"TKPM-Go/module/user/usermodel"
	"github.com/go-playground/validator/v10"
)

const EntityName = "Product Ratings"

type Rating struct {
	common.SQLModel `json:",inline"`
	Point           float32                       `json:"point" gorm:"column:point"`
	Comment         string                        `json:"comment" gorm:"column:comment;"`
	UserId          int                           `json:"-" gorm:"column:user_id"`
	UserUID         *common.UID                   `json:"user_id" gorm:"-"`
	ProductId       int                           `json:"-" gorm:"column:product_id"`
	ProductUID      *common.UID                   `json:"product_id" gorm:"-"`
	User            *usermodel.User               `gorm:"foreignKey:UserId"`
	OrderDetailId   int                           `json:"-" gorm:"column:order_id"`
	OrderDetailUID  *common.UID                   `json:"detail_id"`
	Product         *Product                      `json:"product" gorm:"foreignKey:ProductId"`
	OrderDetail     *orderdetailmodel.OrderDetail `json:"order-detail" gorm:"foreignKey:OrderDetailId"`
}

func (Rating) TableName() string {
	return "product_ratings"
}

func (r *Rating) Mask() {
	r.GenUID(common.DbTypeProductRating)

	userUID := common.NewUID(uint32(r.UserId), int(common.DbTypeUser), 1)
	r.UserUID = &userUID

	productUID := common.NewUID(uint32(r.ProductId), int(common.DbTypeProduct), 1)
	r.ProductUID = &productUID

	detailUID := common.NewUID(uint32(r.OrderDetailId), int(common.DbTypeOrderDetail), 1)
	r.OrderDetailUID = &detailUID
}

func (r *Rating) GetProductID() int {
	return r.ProductId
}

type RatingCreate struct {
	common.SQLModel `json:",inline"`
	Point           float32            `json:"point" gorm:"column:point"`
	Comment         string             `json:"comment" gorm:"column:comment;"`
	UserId          int                `json:"-" gorm:"column:user_id"`
	User            *common.SimpleUser `json:"user"`
	ProductId       int                `json:"-" gorm:"column:product_id"`
	OrderDetailId   int                `json:"-" gorm:"column:detail_id"`
	OrderDetailUID  *common.UID        `json:"detail_id" gorm:"-"`
}

func (RatingCreate) TableName() string {
	return Rating{}.TableName()
}

func (r *RatingCreate) Mask() {
	r.GenUID(common.DbTypeProductRating)
}

func (r *RatingCreate) GetUserID() int {
	return r.UserId
}

func (r *RatingCreate) GetProductID() int {
	return r.ProductId
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

	if err := validate.Var(res.UserId, "required"); err != nil {
		return ErrUserIdIsRequired
	}

	if err := validate.Var(res.ProductId, "required"); err != nil {
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
