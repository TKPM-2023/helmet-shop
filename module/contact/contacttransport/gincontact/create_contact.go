package gincontact

import (
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactbiz"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactmodel"
	"github.com/orgball2608/helmet-shop-be/module/contact/contactstorage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateContact(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		db := appCtx.GetMainDBConnection()
		requester := context.MustGet(common.CurrentUser).(common.Requester)
		var data contactmodel.ContactCreate

		if err := context.ShouldBind(&data); err != nil {
			panic(err)
		}

		data.UserId = requester.GetUserId()

		store := contactstorage.NewSQLStore(db)
		business := contactbiz.NewCreateContactBusiness(store)

		if err := business.CreateContact(context.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
