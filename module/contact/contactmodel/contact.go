package contactmodel

import (
	"TKPM-Go/common"

	"github.com/go-playground/validator/v10"
)

const EntityName = "contacts"

type Contact struct {
	common.SQLModel `json:",inline"`
	User_ID         int         `json:"-" gorm:"column:user_id"`
	User_UID        *common.UID `json:"user_id" gorm:"-"`
	Name            string      `json:"name" gorm:"column:name"`
	Addr            string      `json:"addr" gorm:"column:addr"`
	Phone           string      `json:"phone" gorm:"column:phone"`
}

func (c *Contact) GenUserUID() {
	uid := common.NewUID(uint32(c.User_ID), int(common.DbTypeUser), 1)
	c.User_UID = &uid
}

func (Contact) TableName() string {
	return "contacts"
}

func (c *Contact) Mask() {
	c.GenUID(common.DbTypeContact)
}

func (c *Contact) GetUserID() int {
	return c.User_ID
}

type ContactCreate struct {
	common.SQLModel `json:",inline"`
	User_ID         int         `json:"-" gorm:"column:user_id"`
	//User_UID        *common.UID `json:"user_id" gorm:"-"`
	Name            string      `json:"name" gorm:"column:name"`
	Addr            string      `json:"addr" gorm:"column:addr"`
	Phone           string      `json:"phone" gorm:"column:phone"`
}

func (ContactCreate) TableName() string {
	return Contact{}.TableName()
}

func (c *ContactCreate) Mask() {
	c.GenUID(common.DbTypeContact)
}

func (c *ContactCreate) GetUserID() int {
	return c.User_ID
}

func (res *ContactCreate) Validate() error {
	validate := validator.New()

	if err:=validate.Var(res.User_ID, "required") ;err != nil {
		return ErrContactUserIdIsRequired
	}

	if err:=validate.Var(res.Name, "required") ;err != nil {
		return ErrContactNameIsRequired
	}
	
	if err:=validate.Var(res.Addr, "required") ;err != nil {
		return ErrContactAddressIsRequired
	}

	if err:=validate.Var(res.Phone, "required") ;err != nil {
		return ErrContactPhoneIsRequired
	}
	return nil
}

type ContactUpdate struct {
	common.SQLModel `json:",inline"`
	User_ID         int         `json:"-" gorm:"column:user_id"`
	//User_UID        *common.UID `json:"user_id" gorm:"-"`
	Name            string      `json:"name" gorm:"column:name"`
	Addr            string      `json:"addr" gorm:"column:addr"`
	Phone           string      `json:"phone" gorm:"column:phone"`
}

func (res *ContactUpdate) Validate() error {
	//validate := validator.New()
	
	return nil
}

func (ContactUpdate) TableName() string {
	return Contact{}.TableName()
}