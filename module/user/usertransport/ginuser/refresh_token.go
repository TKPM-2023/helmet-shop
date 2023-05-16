package ginuser

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/hasher"
	"LearnGo/component/tokenprovider/jwt"
	"LearnGo/module/user/userbusiness"
	"LearnGo/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RefreshToken(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		data := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey()) //appctx.SecretKey()

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbusiness.NewRefreshBusiness(appCtx, store, 30*60, tokenProvider, md5)
		account, err := biz.Refresh(c.Request.Context(), data)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
