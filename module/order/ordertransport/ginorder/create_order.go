package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/contact/contactbiz"
	"TKPM-Go/module/contact/contactstorage"
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

		data.UserId = requester.GetUserId()
		data.ContactId = int(data.ContactUID.GetLocalID())
		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewCreateOrderBusiness(store)

		contact_store := contactstorage.NewSQLStore(db)
		contact_business := contactbiz.NewGetContactBusiness(contact_store)
		if _, err := contact_business.GetContact(context.Request.Context(), data.ContactId); err != nil {
			panic(err)
		}

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

		//for each Product

		for i := range data.Products {
			data.Products[i].OrderUID = data.FakeId
			data.Products[i].OrderId = int(data.Products[i].OrderUID.GetLocalID())

			//get product info from model products
			product, err := product_business.GetProduct(context.Request.Context(), int(data.Products[i].ProductOrigin.UID.GetLocalID()))

			if err != nil {
				panic(err)
			}

			//assign to product_origin
			data.Products[i].ProductOrigin.Description = product.Description
			data.Products[i].ProductOrigin.Name = product.Name
			data.Products[i].Price = (float64(product.Price) * float64(data.Products[i].Quantity)) - (float64(product.Price) * float64(data.Products[i].Discount))
			data.Products[i].ProductOrigin.Images = product.Images
			if err := orderdetail_business.CreateOrderDetail(context.Request.Context(), &data.Products[i]); err != nil {
				panic(err)
			}
		}
		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
