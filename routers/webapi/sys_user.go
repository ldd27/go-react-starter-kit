package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/common/comStruct"
	"github.com/jdongdong/go-react-starter-kit/common/errCode"
	"github.com/jdongdong/go-react-starter-kit/models"
)

func Login(c *comStruct.CustomContext) error {
	req := new(models.SeaSysUser)

	err := c.BindEx(req)
	if err != nil {
		return err
	}

	if req.LoginKey == "" || req.Password == "" {
		return errCode.ErrorParams
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
	return c.DataRs(req.WebLogin())
}

func CheckIsLogin(c *comStruct.CustomContext) error {
	req := new(models.SeaSysUser)
	req.Id = c.UserID
	return c.DataRs(req.GetLoginByUserID())
}

func Logout(c *comStruct.CustomContext) error {
	return c.Success(nil)
}
