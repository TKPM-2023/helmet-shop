package ginuser

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/user/userbusiness"
	"TKPM-Go/module/user/usermodel"
	"TKPM-Go/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging
		if err := context.ShouldBind(&pagingData); err != nil {
			context.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter usermodel.Filter
		if err := context.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var results []usermodel.User
		store := userstorage.NewSQLStore(db)
		business := userbusiness.NewListUserBusiness(store)
		results, err := business.ListUser(context.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range results {
			results[i].Mask()
		}

		context.JSON(http.StatusOK, common.NewSuccessResponse(results, pagingData, filter))
	}
}
