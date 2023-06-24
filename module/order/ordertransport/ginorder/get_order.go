package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/orderstorage"
	"net/http"

	"github.com/gin-gonic/gin"
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
		result.GenUserUID()
		result.GenContactUID()
		result.Contact.GenUserUID()
		products := result.Products
		for i := range products {
			products[i].Mask()
			products[i].GenOrderUID()
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
