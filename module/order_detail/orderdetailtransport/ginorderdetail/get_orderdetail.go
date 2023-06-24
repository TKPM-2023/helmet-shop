package ginorderdetail

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/order_detail/orderdetailstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrderDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := orderdetailstorage.NewSQLStore(db)
		business := orderdetailbiz.NewGetOrderDetailBusiness(store)
		result, err := business.GetOrderDetail(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		result.Mask()

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
