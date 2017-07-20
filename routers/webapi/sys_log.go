package webapi

import (
	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
)

func GetPagingSysLog(c *comStruct.CustomContext) error {
	req := new(models.SeaSysLog)
	req.GetDtlPaging()
	slog.Info(11111)
	return c.AutoPageDataRs(req.GetPaging())
}
