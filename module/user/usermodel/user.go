package usermodel

import (
	"LearnGo/common"
	"github.com/go-playground/validator/v10"
)

const EntityName = "client"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" gorm:"column:email;"`
	Password        string        `json:"-" gorm:"column:password;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"role" gorm:"column:role;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

func (u *User) GetUserID() int {
	return u.Id
}

func (u *User) GetUserEmail() string {
	return u.Email
}
func (u *User) GetUserRole() string {
	return u.Role
}

func (u *User) GetUserId() int {
	return u.Id
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string        `json:"email" form:"email" validate:"required,email" gorm:"column:email;"`
	Password        string        `json:"password" form:"password" validate:"required,min=8,max=20" gorm:"column:password;"`
	LastName        string        `json:"last_name" form:"last_name" validate:"required" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" form:"first_name" validate:"required" gorm:"column:first_name;"`
	Phone           string        `json:"phone" gorm:"column:phone;"`
	Role            string        `json:"-" gorm:"column:role;"`
	Salt            string        `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (res *UserCreate) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Email, "required,email"); err != nil {
		return InvalidEmailFormat
	}

	if err := validate.Var(res.Password, "required,min=8,max=20"); err != nil {
		return InvalidPasswordFormat
	}

	if err := validate.Var(res.FirstName, "required"); err != nil {
		return InvalidFirstNameFormat
	}

	if err := validate.Var(res.LastName, "required"); err != nil {
		return InvalidLastNameFormat
	}
	return nil
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserLogin struct {
	Email    string `json:"email" form:"email" validate:"required,email" gorm:"column:email;"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=20" gorm:"column:password;"`
}

func (res *UserLogin) Validate() error {
	validate := validator.New()

	if err := validate.Var(res.Email, "required,email"); err != nil {
		return InvalidEmailFormat
	}

	if err := validate.Var(res.Password, "required,min=8,max=20"); err != nil {
		return InvalidPasswordFormat
	}
	return nil
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
