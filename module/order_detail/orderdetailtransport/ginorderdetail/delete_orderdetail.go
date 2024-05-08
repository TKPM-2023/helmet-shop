package ginorderdetail

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailbiz"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailstorage"
	"net/http"
)

func DeleteOrderDetail(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := orderdetailstorage.NewSQLStore(db)
		business := orderdetailbiz.NewDeleteOrderDetailBusiness(store)

		if err := business.DeleteOrderDetail(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
