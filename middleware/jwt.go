package middleware

import (
	"fmt"

	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/common/comStruct"
	"github.com/jdongdong/go-react-starter-kit/common/errCode"
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	UserID  string
	jwt.StandardClaims
}

func JwtHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().RequestURI == "/webApi/sysUser/login" {
				return next(c)
			}
			tokenStr := c.Request().Header.Get(echo.HeaderAuthorization)
			if tokenStr == "" {
				return errCode.ErrorInvalidToken
			}
			data, err := tool.ParseToken(tokenStr)
			if err = errCode.CheckErrorInvalidToken(err); err != nil {
				slog.Error(fmt.Sprintf("token str is %s", tokenStr), err)
				return err
			} else {
				sysToken := new(models.SysToken)
				sysToken.Token = tokenStr
				if err := sysToken.CheckTokenExpireTime(20); err != nil {
					slog.Error(fmt.Sprintf("token str is %s", tokenStr), err)
					return err
				}
			}

			cc := c.(*comStruct.CustomContext)
			if vlu, ok := data.(map[string]interface{}); ok {
				cc.UserID = tool.ToString(vlu["Id"])
			}

			if cc.UserID == "" {
				return errCode.ErrorInvalidToken
			}

			return next(c)
		}
	}
}
