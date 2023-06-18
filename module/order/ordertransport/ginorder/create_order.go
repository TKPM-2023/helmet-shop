package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order/orderstorage"
	"TKPM-Go/module/order_detail/orderdetailbiz"
	"TKPM-Go/module/order_detail/orderdetailstorage"
	"TKPM-Go/module/product/productbiz"
	"TKPM-Go/module/product/productstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data ordermodel.OrderCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.User_ID = requester.GetUserId() //int(data.User_UID.GetLocalID())
		data.Contact_ID = int(data.Contact_UID.GetLocalID())
		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewCreateOrderBusiness(store)

		if err := business.CreateOrder(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		//fetch product_origin

		//Create store
		product_store := productstorage.NewSQLStore(db)
		product_business := productbiz.NewGetProductBusiness(product_store)

		orderdetail_store := orderdetailstorage.NewSQLStore(db)
		orderdetail_business := orderdetailbiz.NewCreateOrderDetailBusiness(orderdetail_store)

		//for each Products
		for i := range data.Products {
			data.Products[i].Order_UID=data.FakeId
			data.Products[i].Order_ID = int(data.Products[i].Order_UID.GetLocalID())

			//get product info from model products
			product, err := product_business.GetProduct(context.Request.Context(), int(data.Products[i].Product_Origin.UID.GetLocalID()))

			if err != nil {
				panic(err)
			}

			//assign to prodcut_origin
			data.Products[i].Product_Origin.Description = product.Description
			data.Products[i].Product_Origin.Name = product.Name
			data.Products[i].Price = (float64(product.Price) * float64(data.Products[i].Quantity)) - (float64(product.Price) * float64(data.Products[i].Discount))

			if err := orderdetail_business.CreateOrderDetail(context.Request.Context(), &data.Products[i]); err != nil {
				panic(err)
			}
		}
		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
