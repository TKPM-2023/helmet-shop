package ginproduct

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/product/productbiz"
	"TKPM-Go/module/product/productmodel"
	"TKPM-Go/module/product/productstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProduct(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		pubsub := appCtx.GetPubSub()
		var data productmodel.ProductCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.CategoryUID == nil {
			panic(common.ErrInvalidRequest(nil))
		}

		data.CategoryId = int(data.CategoryUID.GetLocalID())

		store := productstorage.NewSQLStore(db)
		business := productbiz.NewCreateProductBusiness(store, pubsub)

		if err := business.CreateProduct(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
