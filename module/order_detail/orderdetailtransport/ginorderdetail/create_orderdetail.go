package ginorderdetail

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/order_detail/orderdetailmodel"
	"TKPM-Go/module/order_detail/orderdetailstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data orderdetailmodel.OrderDetailCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.Order_UID == nil {
			panic(common.ErrInvalidRequest(nil))
		}

		data.Order_ID = int(data.Order_UID.GetLocalID())

		store := orderdetailstorage.NewSQLStore(db)
		business := orderdetailbiz.NewCreateOrderDetailBusiness(store)

		if err := business.CreateOrderDetail(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
