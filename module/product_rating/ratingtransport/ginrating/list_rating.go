package ginrating

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingbiz"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingmodel"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingstorage"
	"net/http"
)

func ListRating(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMainDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBind(&pagingData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInternal(err))
			return
		}

		var filter ratingmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		pagingData.Fulfill()

		var result []ratingmodel.Rating
		store := ratingstorage.NewSQLStore(db)
		business := ratingbiz.NewListRatingBusiness(store)
		result, err := business.ListRating(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
			if result[i].User != nil {
				result[i].User.Mask()
			}
			if result[i].OrderDetail != nil {
				result[i].OrderDetail.Mask()
			}
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
