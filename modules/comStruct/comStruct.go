package comStruct

import (
	"fmt"
	"net/http"

	"time"

	"reflect"

	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/apiCode"
	"github.com/labstack/echo"
	"github.com/pquerna/ffjson/ffjson"
)

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

type CusHandlerFunc func(*CustomContext) error

type CustomContext struct {
	UserID string
	echo.Context
}

func (this *CustomContext) toResModel(err_code int, info string, r interface{}) ResModel {
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
			slog.Debug(fmt.Sprintf("url:%s success:%s code:%d info:%s r:%s", this.Request().RequestURI, success, err_code, info, string(vlu)))
		} else {
			slog.Debug(fmt.Sprintf("url:%s success:%s code:%d info:%s r:%+v", this.Request().RequestURI, success, err_code, info, r))
		}
	}()
	return res
}

func (this *CustomContext) Success(r interface{}) error {
	vlu, err := ffjson.Marshal(this.toResModel(apiCode.Success, "", r))
	if err != nil {
		return err
	}
	return this.String(http.StatusOK, string(vlu))
}

func (this *CustomContext) Fail(err error) error {
	vlu, err := ffjson.Marshal(this.toResModel(apiCode.FormatApiCode(err), err.Error(), nil))
	if err != nil {
		return err
	}
	return this.String(http.StatusOK, string(vlu))
}

func (this *CustomContext) AutoPageDataRs(i interface{}, total int64, err error) error {
	if err != nil {
		return err
	}
	return this.Success(PageResModel{Data: i, Total: total})
}

func (this *CustomContext) AutoDataRs(i interface{}, err error) error {
	if err != nil {
		return err
	}
	return this.Success(i)
}

func (this *CustomContext) WriteLog(title string, info interface{}) {
	log := new(models.SysLog)
	log.Title = title
	vlu, err := ffjson.Marshal(info)
	if err == nil {
		log.Info = string(vlu)
	} else {
		log.Info = fmt.Sprintf("%+v", info)
	}
	log.IpAddr = this.RealIP()
	log.OpTime = time.Now()
	log.OpUser = this.UserID
	log.ReqUri = this.Request().RequestURI
	log.Insert()
}

func (this *CustomContext) InsertRs(logTag string, req models.InsertInterface) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("CreateBy").CanSet() {
		temp.FieldByName("CreateBy").Set(reflect.ValueOf(this.UserID))
	}
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(logTag, req)
	return req.Insert()
}

func (this *CustomContext) UpdateRs(logTag string, req models.UpdateByIdInterface) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(logTag, req)
	return req.UpdateById()
}

func (this *CustomContext) DeleteRs(logTag string, req models.DeleteByIdInterface) error {
	this.WriteLog(logTag, req)
	return req.DeleteById()
}

func (this *CustomContext) InsertTransRs(logTag string, req models.InsertTransInterface) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("CreateBy").CanSet() {
		temp.FieldByName("CreateBy").Set(reflect.ValueOf(this.UserID))
	}
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(logTag, req)
	return req.InsertTrans()
}

func (this *CustomContext) UpdateTransRs(logTag string, req models.UpdateTransInterface) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName("CreateBy").CanSet() {
		temp.FieldByName("CreateBy").Set(reflect.ValueOf(this.UserID))
	}
	if temp.FieldByName("UpdateBy").CanSet() {
		temp.FieldByName("UpdateBy").Set(reflect.ValueOf(this.UserID))
	}
	this.WriteLog(logTag, req)
	return req.UpdateTrans()
}

func (this *CustomContext) DeleteTransRs(logTag string, req models.DeleteTransInterface) error {
	this.WriteLog(logTag, req)
	return req.DeleteTrans()
}
