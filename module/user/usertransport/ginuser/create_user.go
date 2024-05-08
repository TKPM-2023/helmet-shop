package ginuser

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/component/hasher"
	"github.com/orgball2608/helmet-shop-be/module/user/userbusiness"
	"github.com/orgball2608/helmet-shop-be/module/user/usermodel"
	"github.com/orgball2608/helmet-shop-be/module/user/userstorage"
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
