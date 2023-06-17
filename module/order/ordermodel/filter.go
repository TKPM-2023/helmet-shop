package ordermodel

import "TKPM-Go/common"

// ?
type Filter struct {
	Status  int         `json:"status,omitempty" form:"status"`
	User_Id *common.UID `json:"user_id" form:"user_id"`
}
