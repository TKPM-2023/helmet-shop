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

func UpdateProduct(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		var data productmodel.ProductUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		if data.CategoryUID != nil {
			data.CategoryId = int(data.CategoryUID.GetLocalID())
		}

		store := productstorage.NewSQLStore(db)
		business := productbiz.NewUpdateProductBusiness(store)
		if err := business.UpdateProduct(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
