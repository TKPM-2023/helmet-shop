package ordermodel
//?
type Filter struct {
	Status	int `json:"status,omitempty" form:"status"`
	User_Id int `json:"user_id" form:"user_id"`
}