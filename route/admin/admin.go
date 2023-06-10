package admin

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/middleware"
	"TKPM-Go/module/category/categorytransport/gincategory"
	"TKPM-Go/module/product/producttransport/ginproduct"
	"TKPM-Go/module/upload/uploadtransport/ginupload"
	"TKPM-Go/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
)

func AdminRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	admin := v1.Group("/admin",
		middleware.RequireAuth(appContext),
		middleware.RoleChecker(appContext,
			"admin"))
	admin.DELETE("/upload/remove/:id", ginupload.Remove(appContext))

	//user
	user := admin.Group("/users")
	user.GET("/", ginuser.ListUser(appContext))
	user.POST("/", ginuser.CreateUser(appContext))
	user.PATCH("/:id", ginuser.UpdateUser(appContext))
	user.DELETE("/:id", ginuser.DeleteUser(appContext))
	user.PATCH("/:id/password", ginuser.UpdatePassword(appContext))

	// category
	category := admin.Group("/categories")
	category.POST("/", gincategory.CreateCategory(appContext))
	category.GET("/:id", gincategory.GetCategory(appContext))
	category.GET("/", gincategory.ListCategory(appContext))
	category.PATCH("/:id", gincategory.UpdateCategory(appContext))
	category.DELETE("/:id", gincategory.DeleteCategory(appContext))

	//product
	product := admin.Group("/products")
	product.POST("/", ginproduct.CreateProduct(appContext))
	product.GET("/:id", ginproduct.GetProduct(appContext))
	product.GET("/", ginproduct.ListProduct(appContext))
	product.PATCH("/:id", ginproduct.UpdateProduct(appContext))
	product.DELETE("/:id", ginproduct.DeleteProduct(appContext))
}
