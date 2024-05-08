package ginproduct

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product/productbiz"
	"github.com/orgball2608/helmet-shop-be/module/product/productstorage"
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
