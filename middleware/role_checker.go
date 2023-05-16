package middleware

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"errors"
	"github.com/gin-gonic/gin"
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
