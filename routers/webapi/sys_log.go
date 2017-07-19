package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
)

func GetPagingSysLog(c *comStruct.CustomContext) error {
	req := new(models.SeaSysLog)

	return c.AutoPageDataRs(req.GetPaging())
}
