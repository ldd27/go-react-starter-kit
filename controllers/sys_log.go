package controllers

import "github.com/jdongdong/go-react-starter-kit/models"

type SysLogController struct {
	BaseController
}

// @router /page [get]
func (this *SysLogController) GetPaging() {
	req := new(models.SeaSysLog)
	this.AutoPageDataRs(req.GetPaging())
}
