package ginuser

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/hasher"
	"LearnGo/component/tokenprovider/jwt"
	"LearnGo/module/user/userbusiness"
	"LearnGo/module/user/usermodel"
	"LearnGo/module/user/userstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbusiness.NewLoginBusiness(appCtx, store, 60*30, 60*60*24*30, tokenProvider, md5)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
