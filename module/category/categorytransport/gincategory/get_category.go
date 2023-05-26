package gincategory

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/module/category/categorybiz"
	"LearnGo/module/category/categorystorage"
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

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))

	}
}
