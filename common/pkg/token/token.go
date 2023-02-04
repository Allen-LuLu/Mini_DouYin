package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const(
	secret = "secret"
	Expire = 3600
)

type Claims struct{
	UserId int64 `json:user_id`
	ExpireAt int64 `json:"expire_at"`
	Else jwt.MapClaims
	jwt.RegisteredClaims
}

// GenToken 生成token
// iat:生成时间
func GenToken(iat time.Time,userID int64,payloads map[string]interface{})(string,error){
	claims := Claims{
		UserId:           userID,
		ExpireAt:         iat.Add(time.Second*Expire).Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "user",
		},
	}

	for k ,v := range payloads{
		claims.Else[k] = v
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析token
func ParseToken(tokenString string)(*Claims,error){
	claims,err := jwt.ParseWithClaims(tokenString,&Claims{},func(token *jwt.Token)(interface{},error){
		return []byte(secret),nil
	})

	if claims != nil{
		if claim,ok := claims.Claims.(*Claims);ok&& claims.Valid{
			return claim,nil
		}
	}
	return nil,err
}