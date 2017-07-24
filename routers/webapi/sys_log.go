package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/common/comStruct"
)

func GetPagingSysLog(c *comStruct.CustomContext) error {
	req := new(models.SeaSysLog)
	req.Title = c.QueryParam("title")
	return c.AutoPageDtlRs(req)
}
