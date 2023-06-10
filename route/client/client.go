package client

import (
	"TKPM-Go/component/appctx"
	"TKPM-Go/middleware"
	"TKPM-Go/module/order/ordertransport/ginorder"
	"TKPM-Go/module/user/usertransport/ginuser"
	"TKPM-Go/module/order_detail/orderdetailtransport/ginorderdetail"
	"github.com/gin-gonic/gin"
)

func ClientRoute(appContext appctx.AppContext, v1 *gin.RouterGroup) {
	clients := v1.Group("client", middleware.RequireAuth(appContext))
	clients.GET("/refresh", middleware.RequireAuth(appContext), ginuser.RefreshToken(appContext))

	order := clients.Group("/orders")
	order.POST("/", ginorder.CreateOrder(appContext))
	order.GET("/:id", ginorder.GetOrder(appContext))
	order.PATCH("/:id", ginorder.UpdateOrder(appContext))
	order.DELETE(":id", ginorder.DeleteOrder(appContext))
	order.GET("/", ginorder.ListOrder(appContext))
	
	orderdetail := clients.Group("/orderdetails")
	orderdetail.POST("/", ginorderdetail.CreateOrderDetail(appContext))
	orderdetail.GET("/:id",ginorderdetail.GetOrderDetail(appContext))
}
