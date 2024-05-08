package gincategory

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/category/categorybiz"
	"github.com/orgball2608/helmet-shop-be/module/category/categorymodel"
	"github.com/orgball2608/helmet-shop-be/module/category/categorystorage"
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
