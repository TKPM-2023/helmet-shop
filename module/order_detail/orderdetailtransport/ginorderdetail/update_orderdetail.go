package ginorderdetail

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"TKPM-Go/module/order_detail/orderdetailstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateOrderDetail(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data orderdetailmodel.OrderDetailUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.OrderUID == nil {
			panic(common.ErrInvalidRequest(nil))
		}

		data.OrderId = int(data.OrderUID.GetLocalID())

		store := orderdetailstorage.NewSQLStore(db)
		business := orderdetailbiz.NewUpdateOrderDetailBusiness(store)
		if err := business.UpdateOrderDetail(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
