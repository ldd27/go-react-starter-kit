package comStruct

import (
	"fmt"
	"net/http"

	"github.com/jdongdong/go-lib/slog"
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
	this.JSON(http.StatusOK, this.toResModel(apiCode.Success, "", r))
	return nil
}

func (this *CustomContext) Fail(err error) error {
	this.JSON(http.StatusOK, this.toResModel(apiCode.FormatApiCode(err), err.Error(), nil))
	return nil
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
