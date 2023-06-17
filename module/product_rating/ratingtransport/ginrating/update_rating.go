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

func UpdateRating(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data ratingmodel.RatingUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := ratingstorage.NewSQLStore(db)
		business := ratingbiz.NewUpdateRatingBusiness(store)
		if err := business.UpdateRating(context.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
