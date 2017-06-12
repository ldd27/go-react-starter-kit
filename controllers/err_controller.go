package controllers

import "github.com/jdongdong/go-react-starter-kit/modules/errCode"

type ErrorController struct {
	BaseController
}

func (this *ErrorController) Error404() {
	this.Fail(errCode.Error404)
}

func (this *ErrorController) Error500() {
	this.Fail(errCode.Error500)
}
