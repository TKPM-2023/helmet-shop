package gincategory

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/module/category/categorybiz"
	"LearnGo/module/category/categorymodel"
	"LearnGo/module/category/categorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data categorymodel.CategoryCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := categorystorage.NewSQLStore(db)
		business := categorybiz.NewCreateCategoryBusiness(store)

		if err := business.CreateCategory(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
