package productmodel

type Filter struct {
	Status     int `json:"status,omitempty" form:"status"`
	CategoryId int `json:"category_id" form:"category_id""`
}
