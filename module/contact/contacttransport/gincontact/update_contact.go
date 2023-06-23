package gincontact

import (
	"TKPM-Go/common"
	"TKPM-Go/component/appctx"
	"TKPM-Go/module/contact/contactbiz"
	"TKPM-Go/module/contact/contactmodel"
	"TKPM-Go/module/contact/contactstorage"
	"github.com/gin-gonic/gin"
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

		/*
		if data.User_UID == nil {
			panic(common.ErrInvalidRequest(nil))
		}*/

		data.User_ID = requester.GetUserId()//int(data.User_UID.GetLocalID())

		store := contactstorage.NewSQLStore(db)
		business := contactbiz.NewUpdateContactBusiness(store)
		if err := business.UpdateContact(context.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
