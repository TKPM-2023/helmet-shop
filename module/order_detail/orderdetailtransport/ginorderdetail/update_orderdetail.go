package ginorderdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailbiz"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailstorage"
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
