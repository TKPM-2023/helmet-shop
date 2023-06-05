package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/orderstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewGetOrderBusiness(store)
		result, err := business.GetOrder(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		result.Mask()

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}