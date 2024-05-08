package client

import (
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/middleware"
	"github.com/orgball2608/helmet-shop-be/module/cart/carttransport/gincart"
	"github.com/orgball2608/helmet-shop-be/module/category/categorytransport/gincategory"
	"github.com/orgball2608/helmet-shop-be/module/contact/contacttransport/gincontact"
	"github.com/orgball2608/helmet-shop-be/module/order/ordertransport/ginorder"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailtransport/ginorderdetail"
	"github.com/orgball2608/helmet-shop-be/module/product/producttransport/ginproduct"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingtransport/ginrating"
	"github.com/orgball2608/helmet-shop-be/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func ClientRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	clients := v1.Group("client", middleware.RequireAuth(appContext))

	//Order
	order := clients.Group("/orders")
	order.POST("/", ginorder.CreateOrder(appContext))
	order.GET("/:id", ginorder.GetOrder(appContext))
	order.PATCH("/:id", ginorder.UpdateOrder(appContext))
	order.DELETE(":id", ginorder.DeleteOrder(appContext))
	order.GET("/", ginorder.ListOrder(appContext))

	//OderDetail
	orderDetail := clients.Group("/order-details")
	orderDetail.POST("/", ginorderdetail.CreateOrderDetail(appContext))
	orderDetail.GET("/:id", ginorderdetail.GetOrderDetail(appContext))
	orderDetail.PATCH("/:id", ginorderdetail.UpdateOrderDetail(appContext))
	orderDetail.DELETE("/:id", ginorderdetail.DeleteOrderDetail(appContext))

	//contact
	contact := clients.Group("/contact")
	contact.POST("/", gincontact.CreateContact(appContext))
	contact.GET("/:id", gincontact.GetContact(appContext))
	contact.PATCH("/:id", gincontact.UpdateContact(appContext))
	contact.DELETE("/:id", gincontact.DeleteContact(appContext))
	contact.GET("/", gincontact.ListContact(appContext))

	//Cart
	cart := clients.Group("/carts")
	cart.GET("", gincart.GetCart(appContext))
	cart.PATCH("", gincart.AddProducts(appContext))
	cart.PATCH("/quantity", gincart.UpdateQuantity(appContext))
	cart.DELETE("", gincart.RemoveProducts(appContext))

	//User
	user := clients.Group("/users")
	user.PATCH("/:id", ginuser.UpdateUser(appContext))
	user.PATCH("/:id/password", ginuser.UpdatePassword(appContext))
	user.GET("/profile", ginuser.GetProfile(appContext))

	//Category
	category := clients.Group("/categories")
	category.GET("/:id", gincategory.GetCategory(appContext))
	category.GET("/", gincategory.ListCategory(appContext))

	//product
	product := clients.Group("/products")
	product.GET("/:id", ginproduct.GetProduct(appContext))
	product.GET("/", ginproduct.ListProduct(appContext))
	product.GET("search", ginproduct.FindProductsByName(appContext))

	//ProductRating
	clients.POST("/products/:id/rating", ginrating.CreateRating(appContext))
	clients.PATCH("products/rating/:id", ginrating.UpdateRating(appContext))
	clients.GET("products/rating/", ginrating.ListRating(appContext))
}
