package ginorder

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order/orderbiz"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
	"github.com/orgball2608/helmet-shop-be/module/order/orderstorage"
	"net/http"
)

func UpdateOrder(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data ordermodel.OrderUpdate
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.ContactUID != nil {
			data.ContactId = int(data.ContactUID.GetLocalID())
		}
		data.UserId = requester.GetUserId()

		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewUpdateOrderBusiness(store)
		if err := business.UpdateOrder(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
