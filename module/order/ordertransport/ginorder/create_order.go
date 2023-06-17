package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order/orderstorage"
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
		/*
		if data.User_UID == nil {
			panic(common.ErrInvalidRequest(nil))
		}*/

		data.User_ID = requester.GetUserId()//int(data.User_UID.GetLocalID())
		data.Contact_ID=int(data.Contact_UID.GetLocalID())
		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewCreateOrderBusiness(store)

		if err := business.CreateOrder(context.Request.Context(), &data); err != nil {
			panic(err)
		}


		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
