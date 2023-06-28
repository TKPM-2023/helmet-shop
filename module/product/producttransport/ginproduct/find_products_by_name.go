package ginproduct

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product/productbiz"
	"TKPM-Go/module/product/productstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FindProductsByName(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		name, ok := context.GetQuery("name")

		if !ok {
			context.JSON(http.StatusBadRequest, common.ErrInternal(nil))
			return
		}

		store := productstorage.NewSQLStore(db)
		business := productbiz.NewFindProductsBusiness(store)
		results, err := business.FindProductsByName(context.Request.Context(), name)

		if err != nil {
			panic(err)
		}

		for i := range results {
			results[i].Mask()
			ratings := results[i].Ratings
			for i := range ratings {
				ratings[i].Mask()
				ratings[i].User.Mask()
			}
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(results))
	}
}
