package gincategory

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/category/categorybiz"
	"github.com/orgball2608/helmet-shop-be/module/category/categorystorage"
	"net/http"
)

func GetCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := categorystorage.NewSQLStore(db)
		business := categorybiz.NewGetCategoryBusiness(store)
		result, err := business.GetCategory(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask()
		products := result.Products
		for i := range products {
			products[i].Mask()
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}
}
