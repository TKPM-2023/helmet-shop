package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/orgball2608/helmet-shop-be/common"
	"github.com/orgball2608/helmet-shop-be/component/appctx"
)

func RoleChecker(appCtx appctx.AppContext, role ...string) func(c *gin.Context) {

	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		var flat = false

		for _, item := range role {
			if u.GetUserRole() == item {
				flat = true
			}
		}

		if !flat {
			panic(common.ErrNoPermission(errors.New("invalid role client")))
		}

		c.Next()
	}
}
