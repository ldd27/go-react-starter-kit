package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/com"
	"github.com/jdongdong/go-react-starter-kit/models"
)

func GetPagingSysLog(c *com.Context) error {
	req := new(models.SeaSysLog)
	req.Title = c.QueryParam("title")
	return c.AutoRsPageDtl(req)
}
