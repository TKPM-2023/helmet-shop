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

func CreateUser(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewCreateUserBusiness(store, md5)

		if err := biz.CreateUser(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
