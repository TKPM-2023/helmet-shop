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
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewGetCartBusiness(store)
		result, err := business.GetCart(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		result.Mask()
		cartProducts := result.CartProducts
		for i := range cartProducts {
			cartProducts[i].Mask()
			cartProducts[i].Product.GenCategoryUID()
			cartProducts[i].Product.Mask()
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
