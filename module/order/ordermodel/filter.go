package ordermodel

import "TKPM-Go/common"

type Filter struct {
	Status      int         `json:"status,omitempty" form:"status"`
	UserId      *common.UID `json:"user_id" form:"user_id"`
	OrderStatus int         `json:"order_status" form:"order_status"`
}
