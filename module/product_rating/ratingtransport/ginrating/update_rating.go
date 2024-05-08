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
