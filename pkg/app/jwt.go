package app

//生成token和解析token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin/blog-service/global"
	"github.com/gin/blog-service/pkg/util"
	"time"
)

type Claims struct {
	AppKey    string `json:"app_key"`
	AppSecret string `json:"app_secret"`
	jwt.StandardClaims
}

//jwt.StandardClaims 结构体，它是 jwt-go 库中预定义的，也是 JWT 的规范，其涉及字段如下： 对应的其实是Payload 的相关字段
//type StandardClaims struct {
//	Audience  string `json:"aud,omitempty"`
//	ExpiresAt int64  `json:"exp,omitempty"`
//	Id        string `json:"jti,omitempty"`
//	IssuedAt  int64  `json:"iat,omitempty"`
//	Issuer    string `json:"iss,omitempty"`
//	NotBefore int64  `json:"nbf,omitempty"`
//	Subject   string `json:"sub,omitempty"`
//}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(appKey, appSecret string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		AppKey:    util.EncodeMD5(appKey),
		AppSecret: util.EncodeMD5(appSecret),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
