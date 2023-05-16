package userbusiness

import (
	"LearnGo/common"
	"LearnGo/component/appctx"
	"LearnGo/component/tokenprovider"
	"LearnGo/module/user/usermodel"
	"context"
)

type RefreshStorage interface {
}

type refreshBusiness struct {
	appCtx            appctx.AppContext
	userStore         LoginStorage
	accessTokenExpiry int // expiry will replace for type TokenConfig
	tokenProvider     tokenprovider.Provider
	hasher            Hasher
}

func NewRefreshBusiness(appCtx appctx.AppContext,
	userStore LoginStorage,
	accessTokenExpiry int,
	tokenProvider tokenprovider.Provider,
	hasher Hasher) *refreshBusiness {
	return &refreshBusiness{
		appCtx:            appCtx,
		userStore:         userStore,
		accessTokenExpiry: accessTokenExpiry,
		tokenProvider:     tokenProvider,
		hasher:            hasher,
	}
}

// 1. Hash pass from input & compare with pass in db
// 2. Provider: issue JWT token for Client
// 3 Access token
// 4. Return token(s)

func (biz *refreshBusiness) Refresh(ctx context.Context, user interface{}) (*usermodel.AccessTokenResponse, error) {

	data, ok := user.(*usermodel.User)

	if !ok {
		return nil, common.ErrInternal(nil)
	}

	payload := tokenprovider.TokenPayload{
		UserId: data.GetUserId(),
		Role:   data.GetUserRole(),
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.accessTokenExpiry)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	token := usermodel.NewRefreshTokenResponse(accessToken)

	return token, nil
}
