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

func UpdateCategory(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data categorymodel.CategoryUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := categorystorage.NewSQLStore(db)
		business := categorybiz.NewUpdateCategoryBusiness(store)
		if err := business.UpdateCategory(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
