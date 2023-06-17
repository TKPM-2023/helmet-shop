package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Role      string `json:"role" gorm:"column:role;"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	CartId    int    `json:"-" gorm:"column:cart_id;"`
	CartUID   *UID   `json:"cart_id" gorm:"-"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)

	uid := NewUID(uint32(u.CartId), int(DbTypeCart), 1)
	u.CartUID = &uid
}
