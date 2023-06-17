package ginorderdetail

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/order_detail/orderdetailstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteOrderDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderdetailstorage.NewSQLStore(db)
		business := orderdetailbiz.NewDeleteOrderDetailBusiness(store)

		if err := business.DeleteOrderDetail(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
