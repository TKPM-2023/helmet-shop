package ginuser

import (
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
	"github.com/orgball2608/helmet-shop-be/module/user/userbusiness"
	"github.com/orgball2608/helmet-shop-be/module/user/userstorage"
	"net/http"
)

func DeleteUser(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		business := userbusiness.NewDeleteUserBusiness(store)

		if err := business.DeleteUser(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
