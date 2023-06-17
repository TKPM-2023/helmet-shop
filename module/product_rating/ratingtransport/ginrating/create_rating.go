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

func CreateRating(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		pubsub := appCtx.GetPubSub()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data ratingmodel.RatingCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		uid, err := common.FromBase58(context.Param("id"))

		if err != nil {
			panic(err)
		}

		data.UserID = requester.GetUserId()
		id := int(uid.GetLocalID())
		data.ProductID = id

		store := ratingstorage.NewSQLStore(db)
		business := ratingbiz.NewCreateRatingBusiness(store, pubsub)

		if err := business.CreateRating(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
