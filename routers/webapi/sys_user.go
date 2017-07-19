package webapi

import (
	"fmt"
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

func Login(c *comStruct.CustomContext) error {
	req := new(models.SeaSysUser)

	err := c.ToJson(req)
	fmt.Println(err)
	if err != nil {
		return errCode.ErrorParams
	}

	if req.LoginKey == "" || req.Password == "" {
		return errCode.ErrorParams
	}

	return c.AutoDataRs(req.WebLogin())
}
