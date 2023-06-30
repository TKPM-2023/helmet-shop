package ginrating

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product_rating/ratingbiz"
	"TKPM-Go/module/product_rating/ratingmodel"
	"TKPM-Go/module/product_rating/ratingstorage"
	"github.com/gin-gonic/gin"
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
			result[i].User.Mask()
			result[i].OrderDetail.Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))

	}
}
