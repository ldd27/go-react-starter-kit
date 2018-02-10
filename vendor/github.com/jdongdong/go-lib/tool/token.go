package tool

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "dsfhw&*kjahsd"

func SetTokenSecretKey(i string) {
	if i != "" {
		secretKey = i
	}
}

func GenToken(i interface{}) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": i,
		"exp":  time.Now().UnixNano(),
	})

	return token.SignedString([]byte(secretKey))
}

func ParseToken(tokenStr string) (interface{}, error) {
	if tokenStr == "" {
		return nil, errors.New("parse token error : no token string")
	}
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("parse token error : Unexpected token signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("parse token error : invalid token")
	}

	if value, ok := token.Claims.(jwt.MapClaims); ok {
		//vlu, isOk := value["exp"].(float64)
		//if !isOk {
		//	return nil, errors.New("parse token error : invalid exp date")
		//}
		//if int64(vlu) < time.Now().Unix() {
		//	return nil, errors.New("parse token error : token expired")
		//}
		return value["data"], nil
	} else {
		return nil, errors.New("parse token error : invalid token data")
	}
}
