package usermodel

type Filter struct {
	Status int    `json:"status,omitempty" form:"status"`
	Role   string `json:"role,omitempty" form:"role"`
}
