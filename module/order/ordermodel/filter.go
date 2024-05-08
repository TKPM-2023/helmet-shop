package ordermodel

import "github.com/orgball2608/helmet-shop-be/common"

type Filter struct {
	Status      int         `json:"status,omitempty" form:"status"`
	UserId      *common.UID `json:"user_id" form:"user_id"`
	OrderStatus int         `json:"order_status" form:"order_status"`
}
