package com

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-react-starter-kit/code/apiCode"
	"github.com/jdongdong/go-react-starter-kit/code/errCode"
	"github.com/json-iterator/go"
	"github.com/labstack/echo"
)

const (
	paramPage = "page"
	paramSize = "size"
)

type Context struct {
	echo.Context
	UserID string

	WriteLogFunc func(info interface{}, logTag ...string)
}

func NewContext(writeLog ...func(info interface{}, logTag ...string)) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &Context{}
			ctx.Context = c
			if len(writeLog) > 0 {
				ctx.WriteLogFunc = writeLog[0]
			}
			if ctx.WriteLogFunc == nil {
				ctx.WriteLogFunc = func(info interface{}, logTag ...string) {}
			}
			return h(ctx)
		}
	}
}

func HandleContext(h func(*Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*Context)
		return h(cc)
	}
}

func NewErrorHandle() echo.HTTPErrorHandler {
	return func(e error, context echo.Context) {
		//ctx := context.(*Context)
		res := ResModel{
			Success: "F",
			Code:    errCode.FormatApiCode(e),
			Info:    e.Error(),
		}
		slog.Trace(fmt.Sprintf("url:%s success:%s code:%d info:%s", context.Request().RequestURI, res.Success, res.Code, e.Error()))
		context.JSON(http.StatusOK, res)
	}
}

func GET(group *echo.Group, url string, han func(c *Context) error) {
	group.GET(url, HandleContext(han))
}

func POST(group *echo.Group, url string, han func(c *Context) error) {
	group.POST(url, HandleContext(han))
}

func PUT(group *echo.Group, url string, han func(c *Context) error) {
	group.PUT(url, HandleContext(han))
}

func DELETE(group *echo.Group, url string, han func(c *Context) error) {
	group.DELETE(url, HandleContext(han))
}

func (c *Context) JSON(code int, i interface{}) error {
	_, pretty := c.QueryParams()["pretty"]
	if c.Echo().Debug || pretty {
		return c.JSONPretty(code, i, "  ")
	}
	b, err := jsoniter.Marshal(i)
	if err != nil {
		return err
	}
	return c.JSONBlob(code, b)
}

func (c *Context) pkgRes(errCode int, info string, r interface{}) ResModel {
	if r == nil {
		r = ""
	}
	success := "T"
	if errCode != apiCode.Success {
		success = "F"
	}

	res := ResModel{Success: success, Code: errCode, Info: info, R: r}
	go func() {
		vlu, err := jsoniter.Marshal(r)
		if err == nil {
			slog.Trace(fmt.Sprintf("url:%s success:%s code:%d info:%s r:%s", c.Request().RequestURI, success, errCode, info, string(vlu)))
		} else {
			slog.Trace(fmt.Sprintf("url:%s success:%s code:%d info:%s r:%+v", c.Request().RequestURI, success, errCode, info, r))
		}
	}()
	return res
}

func (c *Context) Success(i interface{}) error {
	return c.JSON(http.StatusOK, c.pkgRes(apiCode.Success, "", i))
}

func (c *Context) RsPage(i interface{}, total int64, err error) error {
	if err != nil {
		return err
	}
	return c.Success(PageResModel{Data: i, Total: total})
}

func (c *Context) RsData(i interface{}, err error) error {
	if err != nil {
		return err
	}
	return c.Success(i)
}

func (c *Context) Rs(err error) error {
	return c.RsData(nil, err)
}

func (c *Context) ToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errCode.NewErrorParams(err, 1)
	}
	return i, nil
}

func (c *Context) ToIntEx(s string, defaultVlu ...int) int {
	i, err := strconv.Atoi(s)
	if err == nil {
		return i
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (c *Context) ToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, errCode.NewErrorParams(err, 1)
	}
	return i, nil
}

func (c *Context) ToInt64Ex(s string, defaultVlu ...int64) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err == nil {
		return i
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (c *Context) ToFloat64(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 10)
	if err != nil {
		return 0, errCode.NewErrorParams(err, 1)
	}
	return i, nil
}

func (c *Context) ToFloat64Ex(s string, defaultVlu ...float64) float64 {
	i, err := strconv.ParseFloat(s, 10)
	if err == nil {
		return i
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}
}

func (c *Context) BindEx(i interface{}) error {
	err := c.Bind(i)
	if err != nil {
		return errCode.NewErrorJSON(err, 1)
	}
	return nil
}

func (c *Context) GetPageSize() int {
	return c.ToIntEx(c.QueryParam(paramSize), 10)
}

func (c *Context) GetPageIndex() int {
	return c.ToIntEx(c.QueryParam(paramPage), 1)
}

func (c *Context) GetStrParam(name string) (string, error) {
	param := c.Param(name)
	if param == "" {
		return "", errCode.NewErrorParams(1)
	}
	return param, nil
}

func (c *Context) GetStrQueryParam(name string) (string, error) {
	param := c.QueryParam(name)
	if param == "" {
		return "", errCode.NewErrorParams(1)
	}
	return param, nil
}
