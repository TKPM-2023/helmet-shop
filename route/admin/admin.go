package admin

import (
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/middleware"
	"github.com/orgball2608/helmet-shop-be/module/category/categorytransport/gincategory"
	"github.com/orgball2608/helmet-shop-be/module/product/producttransport/ginproduct"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingtransport/ginrating"
	"github.com/orgball2608/helmet-shop-be/module/statistic/statistictransport/ginstatistic"
	"github.com/orgball2608/helmet-shop-be/module/upload/uploadtransport/ginupload"
	"github.com/orgball2608/helmet-shop-be/module/user/usertransport/ginuser"

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
	product.GET("search", ginproduct.FindProductsByName(appContext))

	//Product Rating
	product.DELETE("/rating/:id", ginrating.DeleteRating(appContext))

	//Statistic
	statistic := admin.Group("/statistic")
	statistic.GET("/:year", ginstatistic.GetStatistic(appContext))
}
