package gincontact

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/contact/contactbiz"
	"TKPM-Go/module/contact/contactmodel"
	"TKPM-Go/module/contact/contactstorage"
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

		/*
		if data.User_UID == nil {
			panic(common.ErrInvalidRequest(nil))
		}*/

		data.User_ID = requester.GetUserId()//int(data.User_UID.GetLocalID())

		store := contactstorage.NewSQLStore(db)
		business := contactbiz.NewCreateContactBusiness(store)

		if err := business.CreateContact(context.Request.Context(), &data); err != nil {
			panic(err)
		}


		data.Mask()

		context.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
