package gincategory

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/category/categorybiz"
	"TKPM-Go/module/category/categorystorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteCategory(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := categorystorage.NewSQLStore(db)
		business := categorybiz.NewDeleteCategoryBusiness(store, requester)

		if err := business.DeleteCategory(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
