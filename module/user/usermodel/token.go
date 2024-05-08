package usermodel

import "github.com/orgball2608/helmet-shop-be/component/tokenprovider"

type Token struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token""`
}

func NewToken(at, rt *tokenprovider.Token) *Token {
	return &Token{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

type AccessTokenResponse struct {
	AccessToken *tokenprovider.Token `json:"access_token"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

func NewRefreshTokenResponse(at *tokenprovider.Token) *AccessTokenResponse {
	return &AccessTokenResponse{
		AccessToken: at,
	}
}
