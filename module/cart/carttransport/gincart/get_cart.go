package gincart

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartbiz"
	"github.com/orgball2608/helmet-shop-be/module/cart/cartstorage"
	"net/http"
)

func GetCart(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewGetCartBusiness(store)
		result, err := business.GetCart(context.Request.Context(), requester.GetCartId())

		if err != nil {
			panic(err)
		}
		result.Mask()
		cartProducts := result.CartProducts
		for i := range cartProducts {
			cartProducts[i].Mask()
			cartProducts[i].Product.Mask()
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
