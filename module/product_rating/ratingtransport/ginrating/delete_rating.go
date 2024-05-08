package ginrating

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingbiz"
	"github.com/orgball2608/helmet-shop-be/module/product_rating/ratingstorage"
	"net/http"
)

func DeleteRating(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		pubsub := appCtx.GetPubSub()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := ratingstorage.NewSQLStore(db)
		business := ratingbiz.NewDeleteRatingBusiness(store, pubsub)

		if err := business.DeleteRating(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
