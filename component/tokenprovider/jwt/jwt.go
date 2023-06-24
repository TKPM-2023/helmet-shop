package jwt

import (
	"TKPM-Go/component/tokenprovider"
	"github.com/golang-jwt/jwt"
	"time"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenprovider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	// generate the JWT
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(), //Local() => UTC(): for standard
			IssuedAt:  time.Now().Local().Unix(),                                          //Local() => UTC(): for standard
		},
	})

	myToken, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	// return the token
	return &tokenprovider.Token{
		Token:   myToken,
		Expiry:  expiry,
		Created: time.Now(),
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenprovider.TokenPayload, error) {
	token, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				return nil, tokenprovider.ErrInvalidToken
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, tokenprovider.ErrTokenExpired
			}
		}
		return nil, tokenprovider.ErrInvalidToken
	}

	claims, ok := token.Claims.(*myClaims)

	// validate the token
	if !token.Valid {
		//Check token expire
		currentTimestamp := time.Now().Unix()
		if claims.StandardClaims.ExpiresAt <= currentTimestamp {
			return nil, tokenprovider.ErrTokenExpired
		}

		if !ok {
			return nil, tokenprovider.ErrInvalidToken
		}

		return nil, tokenprovider.ErrInvalidToken
	}

	// return the token
	return &claims.Payload, nil
}

func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
