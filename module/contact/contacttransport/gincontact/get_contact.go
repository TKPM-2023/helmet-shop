package gincontact

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/contact/contactbiz"
	"TKPM-Go/module/contact/contactstorage"
	"github.com/gin-gonic/gin"
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
		result.GenUserUID()
	

		context.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
