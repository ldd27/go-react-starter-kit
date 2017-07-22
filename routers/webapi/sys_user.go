package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
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

	return c.AutoDataRs(req.WebLogin())
}

func CheckIsLogin(c *comStruct.CustomContext) error {
	req := new(models.SeaSysUser)
	req.Id = c.UserID
	return c.AutoDataRs(req.GetLoginByUserID())
}

func Logout(c *comStruct.CustomContext) error {
	return c.Success(nil)
}
