package controllers

import (
	"fmt"

	"pcs_server/modules/code"

	"reflect"
	"time"

	"github.com/jdongdong/go-lib/base"
	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/apiCode"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
	"github.com/pquerna/ffjson/ffjson"
)

type BaseController struct {
	base.BaseController
	UserID string
}

type ResModel struct {
	R       interface{} `json:"r"`
	Success string      `json:"success"`
	Code    int         `json:"code"`
	Info    string      `json:"info"`
}

type PageResModel struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func (this *BaseController) toResModel(err_code int, info string, r interface{}) ResModel {
	if r == nil {
		r = ""
	}
	success := "T"
	if err_code != apiCode.Success {
		success = "F"
	}

	res := ResModel{Success: success, Code: err_code, Info: info, R: r}
	defer func() {
		vlu, err := ffjson.Marshal(r)
		if err == nil {
			slog.Debug(fmt.Sprintf("url:%s success:%s code:%d info:%s r:%s", this.Ctx.Request.RequestURI, success, err_code, info, string(vlu)))
		} else {
			slog.Debug(fmt.Sprintf("url:%s success:%s code:%d info:%s r:%+v", this.Ctx.Request.RequestURI, success, err_code, info, r))
		}
	}()
	return res
}

func (this *BaseController) Success(r interface{}) {
	this.Data["json"] = this.toResModel(code.Success, "", r)
	this.ServeJSON()
}

func (this *BaseController) Fail(err error) {
	this.Data["json"] = this.toResModel(apiCode.FormatApiCode(err), err.Error(), nil)
	this.ServeJSON()
}

func (this *BaseController) ToJson(i interface{}) error {
	err := errCode.CheckErrorInvalidJson(ffjson.Unmarshal(this.Ctx.Input.RequestBody, &i))
	if err != nil {
		this.Fail(err)
	}
	return err
}

func (this *BaseController) ToInt(s string) (int, error) {
	i, err := this.GetInt(s)
	return i, errCode.CheckErrorParams(err)
}

func (this *BaseController) ToInt64(s string) (int64, error) {
	i, err := this.GetInt64(s)
	return i, errCode.CheckErrorParams(err)
}

func (this *BaseController) HasError(err error) bool {
	if err != nil {
		this.Fail(err)
	}
	return err != nil
}

func (this *BaseController) AutoPageDataRs(i interface{}, total int64, err error) {
	if this.HasError(err) {
		return
	}
	this.Success(PageResModel{Data: i, Total: total})
}

func (this *BaseController) AutoDataRs(i interface{}, err error) {
	if this.HasError(err) {
		return
	}
	this.Success(i)
}

func (this *BaseController) AutoRs(err error) {
	if this.HasError(err) {
		return
	}
	this.Success(nil)
}

func (this *BaseController) ParseHeaderToken(neeRs bool, userType string) bool {
	token := this.Ctx.Request.Header.Get("token")
	data, err := tool.ParseToken(token)
	if err = errCode.CheckErrorInvalidToken(err); err != nil {
		slog.Error(fmt.Sprintf("token str is %s", token), err)
		if neeRs {
			this.Fail(err)
		}
		return false
	} else {
		sysToken := new(models.SysToken)
		sysToken.Token = token
		if err := sysToken.CheckTokenExpireTime(20); err != nil {
			slog.Error(fmt.Sprintf("token str is %s", token), err)
			if neeRs {
				this.Fail(err)
			}
			return false
		}
	}

	if vlu, ok := data.(map[string]interface{}); ok {
		this.UserID = tool.ToString(vlu["id"])
	}
	if this.UserID == "" {
		return false
	}
	return true
}

func (this *BaseController) WriteLog(t string, title string, info interface{}) {
	log := new(models.SysLog)
	log.Type = t
	log.Title = title
	vlu, err := ffjson.Marshal(info)
	if err == nil {
		log.Info = string(vlu)
	} else {
		log.Info = fmt.Sprintf("%+v", info)
	}
	log.IpAddr = this.Ctx.Input.IP()
	log.OpTime = time.Now()
	log.OpUser = this.UserID
	log.ReqUri = this.Ctx.Request.RequestURI
	log.Insert()
}

func (this *BaseController) InsertRs(t string, logTag string, req models.InsertInterface) {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("CreateBy").CanSet() {
		temp.FieldByName("CreateBy").Set(reflect.ValueOf(this.UserID))
	}
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(t, logTag, req)
	this.AutoRs(req.Insert())
}

func (this *BaseController) UpdateRs(t string, logTag string, req models.UpdateByIdInterface) {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(t, logTag, req)
	this.AutoRs(req.UpdateById())
}

func (this *BaseController) DeleteRs(t string, logTag string, req models.DeleteByIdInterface) {
	this.WriteLog(t, logTag, req)
	this.AutoRs(req.DeleteById())
}

func (this *BaseController) InsertTransRs(t string, logTag string, req models.InsertTransInterface) {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("CreateBy").CanSet() {
		temp.FieldByName("CreateBy").Set(reflect.ValueOf(this.UserID))
	}
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(t, logTag, req)
	this.AutoRs(req.InsertTrans())
}

func (this *BaseController) UpdateTransRs(t string, logTag string, req models.UpdateTransInterface) {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("CreateBy").CanSet() {
		temp.FieldByName("CreateBy").Set(reflect.ValueOf(this.UserID))
	}
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(t, logTag, req)
	this.AutoRs(req.UpdateTrans())
}

func (this *BaseController) DeleteTransRs(t string, logTag string, req models.DeleteTransInterface) {
	this.WriteLog(t, logTag, req)
	this.AutoRs(req.DeleteTrans())
}
