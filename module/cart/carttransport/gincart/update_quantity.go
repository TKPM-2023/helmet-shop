package gincart

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartbiz"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartmodel"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartstorage"
	"net/http"
)

func UpdateQuantity(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data cartmodel.CartProductDetail

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.ProductId = int(data.ProductUID.GetLocalID())

		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewUpdateQuantityBusiness(store)
		if err := business.UpdateQuantity(context.Request.Context(), requester.GetCartId(), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
