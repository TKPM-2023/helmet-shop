package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/contact/contactbiz"
	"TKPM-Go/module/contact/contactstorage"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/orderstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrder(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewGetOrderBusiness(store)
		result, err := business.GetOrder(context.Request.Context(), int(uid.GetLocalID()))

		
		if err != nil {
			panic(err)
		}

		Contact_store := contactstorage.NewSQLStore(db)
		Contact_business := contactbiz.NewGetContactBusiness(Contact_store)
		Contact_result, err := Contact_business.GetContact(context.Request.Context(),result.Contact_ID)

		result.Mask()
		result.GenUserUID()
		result.GenContactUID()
		Contact_result.GenUID(common.DbTypeContact)
		Contact_result.GenUserUID()
		result.Contact=Contact_result
		products := result.Products
		for i := range products {
			products[i].Mask()
			products[i].GenOrderUID()
		}

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
