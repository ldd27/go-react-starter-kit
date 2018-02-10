package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/code/errCode"
	"github.com/jdongdong/go-react-starter-kit/com"
	"github.com/jdongdong/go-react-starter-kit/models"
)

func Login(c *com.Context) error {
	req := new(models.SeaSysUser)

	err := c.BindEx(req)
	if err != nil {
		return err
	}

	if req.LoginKey == "" || req.Password == "" {
		return errCode.NewErrorParams()
	}

	//user, err := req.WebLogin()
	//if err != nil {
	//	return err
	//}
	//claims := &middleware.JwtCustomClaims{
	//	user.Id,
	//	jwt.StandardClaims{
	//		ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
	//	},
	//}
	//token, err := tool.GenToken(claims)
	//if err != nil {
	//	return err
	//}
	//user.Token = token
	//return c.Success(user)
	return c.RsData(req.WebLogin())
}

func CheckIsLogin(c *com.Context) error {
	req := new(models.SeaSysUser)
	req.Id = c.UserID
	return c.RsData(req.GetLoginByUserID())
}

func Logout(c *com.Context) error {
	return c.Success(nil)
}
