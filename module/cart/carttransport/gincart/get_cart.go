package gincart

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/cart/cartbiz"
	"TKPM-Go/module/cart/cartstorage"
	"github.com/gin-gonic/gin"
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
