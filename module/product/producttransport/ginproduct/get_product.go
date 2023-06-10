package ginproduct

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product/productbiz"
	"TKPM-Go/module/product/productstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := productstorage.NewSQLStore(db)
		business := productbiz.NewGetProductBusiness(store)
		result, err := business.GetProduct(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		result.Mask()
		result.GenCategoryUID()

		ratings := result.Ratings
		for i := range ratings {
			ratings[i].Mask()
			ratings[i].GenUserUID()
			ratings[i].GenProductUID()
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
