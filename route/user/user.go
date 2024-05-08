package user

import (
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/category/categorytransport/gincategory"
	"github.com/orgball2608/helmet-shop-be/module/product/producttransport/ginproduct"
	"github.com/orgball2608/helmet-shop-be/module/upload/uploadtransport/ginupload"
	"github.com/orgball2608/helmet-shop-be/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func UserRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	//User
	v1.POST("/register", ginuser.Register(appContext))
	v1.POST("/authenticate", ginuser.Login(appContext))
	v1.POST("/refresh", ginuser.RefreshToken(appContext))

	//upload
	v1.POST("/upload", ginupload.Upload(appContext))

	//Category
	category := v1.Group("/categories")
	category.GET("/:id", gincategory.GetCategory(appContext))
	category.GET("/", gincategory.ListCategory(appContext))

	//Product
	product := v1.Group("/products")
	product.GET("/:id", ginproduct.GetProduct(appContext))
	product.GET("/", ginproduct.ListProduct(appContext))
	product.GET("search", ginproduct.FindProductsByName(appContext))
}
