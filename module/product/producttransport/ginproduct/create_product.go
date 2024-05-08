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
