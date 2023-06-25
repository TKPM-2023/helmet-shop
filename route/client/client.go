package client

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/middleware"
	"TKPM-Go/module/cart/carttransport/gincart"
	"TKPM-Go/module/category/categorytransport/gincategory"
	"TKPM-Go/module/contact/contacttransport/gincontact"
	"TKPM-Go/module/order/ordertransport/ginorder"
	"TKPM-Go/module/order_detail/orderdetailtransport/ginorderdetail"
	"TKPM-Go/module/product/producttransport/ginproduct"
	"TKPM-Go/module/product_rating/ratingtransport/ginrating"
	"TKPM-Go/module/user/usertransport/ginuser"

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

	//ProductRating
	clients.POST("/products/:id/rating", ginrating.CreateRating(appContext))
	clients.PATCH("products/rating/:id", ginrating.UpdateRating(appContext))
}
