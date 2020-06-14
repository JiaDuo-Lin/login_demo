package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const (
	SecretKey      = "this is a secretKey"            // 密钥
	ExpireDuration = time.Second * time.Duration(300) // 过期时间长度
)

type JWT struct{}

type Token struct {
	Token string `json:"token"`
}

// CreateToken
func CreateToken(userName string, id int64) (*Token, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(ExpireDuration).Unix()
	claims["iat"] = time.Now().Unix()
	claims["username"] = userName
	claims["id"] = id
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	response := &Token{tokenString}
	return response, err
}

// ParseToken 解析Token
func ParseToken(tokenStr string) (token *jwt.Token, err error) {
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	return
}

func CheckToken(tokenStr string) bool {
	token, err := ParseToken(tokenStr)
	// 过期、无法识别等其他错误
	if err != nil || !token.Valid {
		writeUnauthorizedStatus(w, Unauthorized)
		return false
	}
	return true
}
