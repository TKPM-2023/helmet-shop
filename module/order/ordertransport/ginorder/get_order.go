package ginorder

import (
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order/orderbiz"
	"github.com/orgball2608/helmet-shop-be/module/order/orderstorage"
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
		result.Contact.Mask()
		result.User.Mask()
		products := result.Products
		for i := range products {
			products[i].Mask()
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
