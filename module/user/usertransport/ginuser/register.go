package ginuser

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/component/hasher"
	"TKPM-Go/module/cart/cartstorage"
	"TKPM-Go/module/user/userbusiness"
	"TKPM-Go/module/user/usermodel"
	"TKPM-Go/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appctx.AppContext) func(*gin.Context) {
	return func(ctx *gin.Context) {
		db := appCtx.GetMainDBConnection()
		pubsub := appCtx.GetPubSub()
		var data usermodel.UserCreate

		if err := ctx.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		cartStore := cartstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbusiness.NewRegisterBusiness(store, cartStore, md5, pubsub)

		if err := biz.Register(ctx.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		ctx.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
