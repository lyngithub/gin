package user

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"xx/app/middlewares"
	"xx/app/models"
)

func CreateJwtToken(uid uint, nickname string, role uint, expiresAt int64) (string, error) {
	j := middlewares.NewJWT()
	nowTime := time.Now().Unix()
	claims := models.CustomClaims{
		ID:          uid,
		NickName:    nickname,
		AuthorityId: role,
		StandardClaims: jwt.StandardClaims{
			NotBefore: nowTime,   //签名的生效时间
			ExpiresAt: expiresAt, //30天过期
			Issuer:    "x-token",
		},
	}

	return j.CreateToken(claims)
}
