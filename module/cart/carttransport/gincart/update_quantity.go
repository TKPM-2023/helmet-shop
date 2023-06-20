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

func UpdateQuantity(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data cartmodel.CartProductDetail
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.ProductId = int(data.ProductUID.GetLocalID())

		store := cartstorage.NewSQLStore(db)
		business := cartbiz.NewUpdateQuantityBusiness(store)
		if err := business.UpdateQuantity(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
