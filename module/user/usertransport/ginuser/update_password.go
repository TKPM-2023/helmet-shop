package ginuser

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/component/hasher"
	"TKPM-Go/module/user/userbusiness"
	"TKPM-Go/module/user/usermodel"
	"TKPM-Go/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdatePassword(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data usermodel.PasswordUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		business := userbusiness.NewUpdatePasswordStore(store, md5)
		if err := business.UpdatePassword(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
