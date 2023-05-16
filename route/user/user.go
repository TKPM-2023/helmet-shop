package user

import (
	"LearnGo/component/appctx"
	"LearnGo/middleware"
	"LearnGo/module/upload/uploadtransport/ginupload"
	"LearnGo/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func UserRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.Login(appContext))
	v1.GET("/refresh", middleware.RequireAuth(appContext), ginuser.RefreshToken(appContext))

	v1.GET("/profile", middleware.RequireAuth(appContext), ginuser.GetProfile(appContext))
	v1.POST("/upload", ginupload.Upload(appContext))
}
