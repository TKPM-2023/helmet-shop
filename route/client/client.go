package client

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/middleware"
	"TKPM-Go/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func ClientRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	clients := v1.Group("client", middleware.RequireAuth(appContext))
	clients.GET("/refresh", middleware.RequireAuth(appContext), ginuser.RefreshToken(appContext))
}
