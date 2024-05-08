package ginorderdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailbiz"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailstorage"
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
