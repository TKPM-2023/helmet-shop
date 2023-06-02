package ordermodel
//?
type Filter struct {
	Status	int `json:"status,omitempty" form:"status"`
	OrderId int `json:"order_id" form:"order_id"`
}