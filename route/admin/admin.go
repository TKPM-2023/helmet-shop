package admin

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/middleware"
	"TKPM-Go/module/category/categorytransport/gincategory"
	"TKPM-Go/module/upload/uploadtransport/ginupload"
	"TKPM-Go/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func AdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequireAuth(appContext),
		middleware.RoleChecker(appContext,
			"admin"))

	admin.GET("/profile", ginuser.GetProfile(appContext))
	admin.DELETE("/upload/remove/:id", ginupload.Remove(appContext))

	// category
	category := admin.Group("/categories")
	category.POST("/", gincategory.CreateCategory(appContext))
	category.GET("/:id", gincategory.GetCategory(appContext))
	category.GET("/", gincategory.ListCategory(appContext))
	category.PATCH("/:id", gincategory.UpdateCategory(appContext))
	category.DELETE("/:id", gincategory.DeleteCategory(appContext))
}
