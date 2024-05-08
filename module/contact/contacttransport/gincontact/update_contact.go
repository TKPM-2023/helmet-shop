package gincontact

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactbiz"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactstorage"
	"net/http"
)

func UpdateContact(ctx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := ctx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data contactmodel.ContactUpdate
		uid, err := common.FromBase58(context.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := contactstorage.NewSQLStore(db)
		business := contactbiz.NewUpdateContactBusiness(store)
		if err := business.UpdateContact(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
