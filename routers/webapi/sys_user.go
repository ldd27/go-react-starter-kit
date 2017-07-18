package webapi

import (
	"net/http"

	"fmt"

	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	Id       string
	UserName string
	jwt.StandardClaims
}

func Login(c echo.Context) error {
	req := new(models.SeaSysUser)

	c.Bind(req)

	fmt.Println(req.LoginKey, req.Password, 2222)
	if req.LoginKey == "" || req.Password == "" {
		return errCode.ErrorParams
	}

	item, err := req.WebLogin()
	if err != nil {
		return err
	}

	// Set custom claims
	claims := &JwtCustomClaims{
		item.Id,
		item.UserName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	item.Token = t
	return c.JSON(http.StatusOK, item)
}
