package ginorderdetail

import (
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailbiz"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailmodel"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailstorage"
	"github.com/orgball2608/helmet-shop-be/module/product/productbiz"
	"github.com/orgball2608/helmet-shop-be/module/product/productstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrderDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data orderdetailmodel.OrderDetailCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.OrderUID == nil {
			panic(common.ErrInvalidRequest(nil))
		}

		data.OrderId = int(data.OrderUID.GetLocalID())

		productStore := productstorage.NewSQLStore(db)
		productBusiness := productbiz.NewGetProductBusiness(productStore)
		product, err := productBusiness.GetProduct(context.Request.Context(), int(data.ProductOrigin.UID.GetLocalID()))

		if err != nil {
			panic(err)
		}

		data.ProductOrigin.Description = product.Description
		data.ProductOrigin.Name = product.Name
		data.Price = (float64(product.Price) * float64(data.Quantity)) - (float64(product.Price) * float64(data.Discount))
		store := orderdetailstorage.NewSQLStore(db)
		business := orderdetailbiz.NewCreateOrderDetailBusiness(store)

		if err := business.CreateOrderDetail(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
