package util

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret []byte

type Claims struct {
	Id int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}


func GenerateToken(id int, username, password string) (string, error) {
	jwtSecret = []byte("a")
	nowTime := time.Now()
	expireTime := nowTime.AddDate(0,0,7)

	m := md5.New()
	m.Write([]byte(password))

	mPassword := hex.EncodeToString(m.Sum(nil))

	claims := Claims{
		id,
		username,
		mPassword,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}


func ParseToken(token string) (*Claims, error) {
	jwtSecret = []byte("a")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {

		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}