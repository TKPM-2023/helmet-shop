package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order/orderstorage"
//	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/order_detail/orderdetailstorage"
//	"TKPM-Go/module/product/productbiz"
	"TKPM-Go/module/product/productstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data ordermodel.OrderCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId() //int(data.User_UID.GetLocalID())
		data.ContactId = int(data.ContactUID.GetLocalID())

		store := orderstorage.NewSQLStore(db)
		product_store := productstorage.NewSQLStore(db)
		orderdetail_store := orderdetailstorage.NewSQLStore(db)
		business := orderbiz.NewCreateOrderBusiness(store,product_store, orderdetail_store)

		if err := business.CreateOrder(context.Request.Context(), &data); err != nil {
			panic(err)
		}
		
		data.Mask()
		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
