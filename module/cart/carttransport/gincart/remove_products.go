package gincart

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/cart/cartbiz"
	"TKPM-Go/module/cart/cartmodel"
	"TKPM-Go/module/cart/cartstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RemoveProducts(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		pubsub := ctx.GetPubSub()
		var data cartmodel.RemoveCartProducts
		requester := context.MustGet(common.CurrentUser).(common.Requester)

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewRemoveProductsBusiness(store, pubsub)
		if err := business.RemoveProducts(context.Request.Context(), requester.GetCartId(), data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
