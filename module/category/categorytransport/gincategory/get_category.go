package gincategory

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorybiz"
	"TKPM-Go/module/category/categorystorage"
	"github.com/gin-gonic/gin"
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
