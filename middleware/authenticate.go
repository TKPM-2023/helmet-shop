package middleware

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/tokenprovider/jwt"
	"LearnGo/module/user/userstore"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//Authorization : Bearn{token}
	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}
	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequireAuth(appCtx appctx.AppContext) func(ctx *gin.Context) {

	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()

		store := userstore.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)

		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			//c.AbortWithStatusJSON(http.StatusUnauthorized, err)
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}

}
