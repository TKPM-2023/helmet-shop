package ginproduct

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product/productbiz"
	"github.com/orgball2608/helmet-shop-be/module/product/productmodel"
	"github.com/orgball2608/helmet-shop-be/module/product/productstorage"
	"net/http"
)

func ListProduct(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter productmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var results []productmodel.Product
		store := productstorage.NewSQLStore(db)
		business := productbiz.NewListProductBusiness(store)
		results, err := business.ListProduct(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range results {
			results[i].Mask()
			ratings := results[i].Ratings
			for i := range ratings {
				ratings[i].Mask()
				ratings[i].User.Mask()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(results, pagingData, filter))

	}
}
