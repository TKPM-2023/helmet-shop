package client

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/middleware"
	"TKPM-Go/module/cart/carttransport/gincart"
	"TKPM-Go/module/contact/contacttransport/gincontact"
	"TKPM-Go/module/order/ordertransport/ginorder"
	"TKPM-Go/module/order_detail/orderdetailtransport/ginorderdetail"
	"TKPM-Go/module/product_rating/ratingtransport/ginrating"
	"TKPM-Go/module/user/usertransport/ginuser"

	"github.com/gin-gonic/gin"
)

func ClientRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	clients := v1.Group("client", middleware.RequireAuth(appContext))
	clients.GET("/refresh", middleware.RequireAuth(appContext), ginuser.RefreshToken(appContext))

	//Order
	order := clients.Group("/orders")
	order.POST("/", ginorder.CreateOrder(appContext))
	order.GET("/:id", ginorder.GetOrder(appContext))
	order.PATCH("/:id", ginorder.UpdateOrder(appContext))
	order.DELETE(":id", ginorder.DeleteOrder(appContext))
	order.GET("/", ginorder.ListOrder(appContext))

	//OderDetail
	orderdetail := clients.Group("/orderdetails")
	orderdetail.POST("/", ginorderdetail.CreateOrderDetail(appContext))
	orderdetail.GET("/:id", ginorderdetail.GetOrderDetail(appContext))

	//ProductRating
	clients.POST("/products/:id/rating", ginrating.CreateRating(appContext))
	clients.PATCH("products/rating/:id", ginrating.UpdateRating(appContext))

	//contact
	contact := clients.Group("/contacts")
	contact.POST("/", gincontact.CreateContact(appContext))
	contact.GET("/:id", gincontact.GetContact(appContext))
	contact.PATCH("/:id", gincontact.UpdateContact(appContext))
	contact.DELETE("/:id", gincontact.DeleteContact(appContext))

	//Cart
	cart := clients.Group("/carts")
	cart.GET("/:id", gincart.GetCart(appContext))
	cart.PATCH("/:id", gincart.AddProducts(appContext))
	cart.PATCH("/:id/quantity", gincart.UpdateQuantity(appContext))
	cart.DELETE("/:id", gincart.RemoveProducts(appContext))
}
