package funcs

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"rj97807_work_serve/utils"
	"time"
)

type UserClaim struct {
	Id   int
	Uid  string
	Name string
	jwt.StandardClaims
}

// Md5 加密
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// YieldToken 生成Token
func YieldToken(id, exTime int, uid, name string) (string, error) {
	uc := UserClaim{
		Id:   id,
		Uid:  uid,
		Name: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(exTime)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	tokenString, err := token.SignedString([]byte(utils.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//token解析
func AnalyzeToken(token string) (*UserClaim, error) {
	uc := new(UserClaim)
	claims, err := jwt.ParseWithClaims(token, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(utils.JwtKey), nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return uc, errors.New("token已过期")
	}
	return uc, err
}
