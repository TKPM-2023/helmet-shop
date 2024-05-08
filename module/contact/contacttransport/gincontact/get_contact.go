package gincontact

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactbiz"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactstorage"
	"net/http"
)

func GetContact(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := contactstorage.NewSQLStore(db)
		business := contactbiz.NewGetContactBusiness(store)
		result, err := business.GetContact(context.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		result.Mask()

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
