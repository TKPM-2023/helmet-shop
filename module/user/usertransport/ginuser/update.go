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

func UpdateUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data usermodel.UserUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		business := userbusiness.NewUpdateInfoUserBusiness(store)
		if err := business.UpdateUserInfo(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
