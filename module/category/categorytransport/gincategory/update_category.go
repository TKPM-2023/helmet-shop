package gincategory

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorybiz"
	"TKPM-Go/module/category/categorymodel"
	"TKPM-Go/module/category/categorystorage"
	"github.com/gin-gonic/gin"
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
		context.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
