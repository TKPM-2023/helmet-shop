package ginproduct

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product/productbiz"
	"github.com/orgball2608/helmet-shop-be/module/product/productstorage"
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

		ratings := result.Ratings
		for i := range ratings {
			ratings[i].Mask()
			ratings[i].User.Mask()
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
