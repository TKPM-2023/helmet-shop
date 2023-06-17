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
		var data cartmodel.RemoveCartProducts
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewRemoveProductsBusiness(store)
		if err := business.RemoveProducts(context.Request.Context(), int(uid.GetLocalID()), data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
