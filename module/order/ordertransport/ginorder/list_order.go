package ginorder

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/contact/contactbiz"
	"TKPM-Go/module/contact/contactstorage"
	"TKPM-Go/module/order/orderbiz"
	"TKPM-Go/module/order/ordermodel"
	"TKPM-Go/module/order/orderstorage"
	"TKPM-Go/module/user/userstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListOrder(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter ordermodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var result []ordermodel.Order
		store := orderstorage.NewSQLStore(db)
		business := orderbiz.NewListOrderBusiness(store)
		result, err := business.ListOrder(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		Contact_store := contactstorage.NewSQLStore(db)
		Contact_business := contactbiz.NewGetContactBusiness(Contact_store)

		User_store := userstorage.NewSQLStore(db)

		if err != nil {
			panic(err)
		}
		for i := range result {
			Contact_result, err := Contact_business.GetContact(c.Request.Context(), result[i].Contact_ID)
			if err == nil {
				Contact_result.GenUserUID()
				result[i].Contact = Contact_result
				result[i].Mask()
				result[i].GenUserUID()
			}
			User, err := User_store.FindUser(c.Request.Context(), map[string]interface{}{"id": result[i].User_ID})
			if err==nil {
				result[i].User=User
			}
			for j := range result[i].Products {
				result[i].Products[j].GenOrderUID()
				result[i].Products[j].GenUID(common.DbTypeOrder_Detail)
			}

		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
