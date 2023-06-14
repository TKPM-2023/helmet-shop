package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order/orderstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOrder(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data ordermodel.OrderUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.User_UID == nil {
			panic(common.ErrInvalidRequest(nil))
		}

		data.User_ID = int(data.User_UID.GetLocalID())

		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewUpdateOrderBusiness(store)
		if err := business.UpdateOrder(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
