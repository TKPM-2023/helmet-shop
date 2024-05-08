package ginorder

import (
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order/orderbiz"
	"github.com/orgball2608/helmet-shop-be/module/order/ordermodel"
	"github.com/orgball2608/helmet-shop-be/module/order/orderrepository"
	"github.com/orgball2608/helmet-shop-be/module/order/orderstorage"
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
		repo := orderrepository.NewListOrderRepo(store)
		business := orderbiz.NewListOrderBusiness(repo)
		result, err := business.ListOrder(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			if err == nil {
				result[i].Mask()
				result[i].Contact.Mask()
				result[i].User.Mask()
				products := result[i].Products
				for i := range products {
					products[i].Mask()
				}
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
