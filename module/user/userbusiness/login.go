package userbusiness

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/tokenprovider"
	"LearnGo/module/user/usermodel"
	"context"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

type loginBusiness struct {
	appCtx             appctx.AppContext
	userStore          LoginStorage
	accessTokenExpiry  int
	refreshTokenExpiry int
	tokenProvider      tokenprovider.Provider
	hasher             Hasher
}

func NewLoginBusiness(appCtx appctx.AppContext,
	userStore LoginStorage,
	accessTokenExpiry int,
	refreshTokenExpiry int,
	tokenProvider tokenprovider.Provider,
	hasher Hasher) *loginBusiness {
	return &loginBusiness{
		appCtx:             appCtx,
		userStore:          userStore,
		accessTokenExpiry:  accessTokenExpiry,
		refreshTokenExpiry: refreshTokenExpiry,
		tokenProvider:      tokenProvider,
		hasher:             hasher,
		//tokenConfig:   tokenConfig,
	}
}

// 1. Find user, email
// 2. Hash pass from input & compare with pass in db
// 3. Provider: issue JWT token for Client
// 3.1 Access token & Refresh token
// 4. Return token(s)

func (biz *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*usermodel.Token, error) {
	user, err := biz.userStore.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	passwordHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passwordHashed {
		return nil, usermodel.ErrEmailOrPasswordInvalid
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	//biz.tokenConfig.GetAtExp() ===> biz.expiry
	accessToken, err := biz.tokenProvider.Generate(payload, biz.accessTokenExpiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	refreshToken, err := biz.tokenProvider.Generate(payload, biz.refreshTokenExpiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	token := usermodel.NewToken(accessToken, refreshToken)

	return token, nil
}
