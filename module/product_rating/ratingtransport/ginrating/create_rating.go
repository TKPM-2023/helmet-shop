package ginrating

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/order_detail/orderdetailstorage"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingbiz"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingmodel"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingstorage"
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

		data.UserId = requester.GetUserId()
		id := int(uid.GetLocalID())
		data.ProductId = id

		if data.OrderDetailUID == nil {
			panic(common.ErrInvalidRequest(nil))
		}

		data.OrderDetailId = int(data.OrderDetailUID.GetLocalID())

		store := ratingstorage.NewSQLStore(db)
		detailStore := orderdetailstorage.NewSQLStore(db)
		business := ratingbiz.NewCreateRatingBusiness(store, detailStore, pubsub)

		if err := business.CreateRating(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
