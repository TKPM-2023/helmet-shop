package ginuser

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/hasher"
	"LearnGo/module/user/userbusiness"
	"LearnGo/module/user/usermodel"
	"LearnGo/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewRegisterBusiness(store, md5)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
