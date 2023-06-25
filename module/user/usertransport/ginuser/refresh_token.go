package ginuser

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/component/hasher"
	"TKPM-Go/component/tokenprovider/jwt"
	"TKPM-Go/module/user/userbusiness"
	"TKPM-Go/module/user/usermodel"
	"TKPM-Go/module/user/userstorage"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())
		var data usermodel.RefreshToken

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		payload, err := tokenProvider.Validate(data.RefreshToken)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		biz := userbusiness.NewRefreshBusiness(appCtx, store, 30*60, tokenProvider, md5)
		account, err := biz.Refresh(c.Request.Context(), user)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
