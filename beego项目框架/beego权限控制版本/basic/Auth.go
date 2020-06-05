package basic

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/dgrijalva/jwt-go"
)

var (
	EncryptKey     = []byte("wyl")
	IssUser        = "wyl"
	ExpireInterval = 20
)

func init() {
	EncryptKey = []byte(beego.AppConfig.DefaultString("jwt.encrypt", string(EncryptKey)))
}

type CustClaims struct {
	Sequence int64 `json:"seq"`

	// recommended having
	jwt.StandardClaims
}

func CreateToken(seq int64, id string) (string, int64) {
	expireTime := time.Now().Add(time.Hour * time.Duration(ExpireInterval)).Unix()
	claims := CustClaims{
		seq,
		jwt.StandardClaims{
			Id:        id,
			ExpiresAt: expireTime,
			Issuer:    IssUser,
		},
	}

	// Create the token using your claims
	c_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Signs the token with a secret.
	signedToken, _ := c_token.SignedString(EncryptKey)

	return signedToken, expireTime
}

func ParseToken(tokenString string, key string) (interface{}, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Println(err, ok)
		return "", false
	}
}
