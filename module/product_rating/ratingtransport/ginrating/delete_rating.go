package ginrating

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product_rating/ratingbiz"
	"TKPM-Go/module/product_rating/ratingstorage"
	"github.com/gin-gonic/gin"
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
