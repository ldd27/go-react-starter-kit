package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

func Login(c *comStruct.CustomContext) error {
	req := new(models.SeaSysUser)

	err := c.Bind(req)
	if err != nil {
		return err
	}

	if req.LoginKey == "" || req.Password == "" {
		return errCode.ErrorParams
	}

	return c.AutoDataRs(req.WebLogin())
}
