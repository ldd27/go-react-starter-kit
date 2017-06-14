package controllers

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

type SysUserController struct {
	BaseController
}

// @router /checkIsLogin [get]
func (this *SysUserController) CheckIsLogin() {
	req := new(models.SeaSysUser)
	req.Id = this.UserID
	this.AutoDataRs(req.GetLoginByUserID())
}

// @router /login [post]
func (this *SysUserController) Login() {
	req := new(models.SeaSysUser)
	if !this.ToJson(&req) {
		return
	}
	if req.LoginKey == "" || req.Password == "" {
		this.Fail(errCode.ErrorParams)
		return
	}

	this.AutoDataRs(req.WebLogin())
}

// @router /logout [post]
func (this *SysUserController) Logout() {
	this.AutoRs(nil)
}
