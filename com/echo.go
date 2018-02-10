package com

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/code/errCode"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type EchoConf struct {
	Debug               bool
	HideBanner          bool
	EnableGzip          bool
	RemoveTrailingSlash bool //删除 url 地址最后的/
	EnableRecover       bool
	EnableCORS          bool
	EnableLog           bool
	Context             echo.MiddlewareFunc
	ErrorHandler        echo.HTTPErrorHandler
	IgnoreFavicon       bool
	//EnableJWT           bool
	//JWTSkipper          middleware.Skipper
}

var defaultEchoConf = EchoConf{
	Debug:               false,
	HideBanner:          true,
	EnableGzip:          true,
	RemoveTrailingSlash: true,
	EnableRecover:       true,
	EnableCORS:          true,
	EnableLog:           true,
	IgnoreFavicon:       true,
	//EnableJWT:           true,
	//JWTSkipper:          middleware.DefaultSkipper,
}

func NewEcho(opts ...func(*EchoConf)) *echo.Echo {
	conf := defaultEchoConf
	for _, o := range opts {
		o(&conf)
	}

	e := echo.New()

	e.Debug = conf.Debug
	e.HideBanner = conf.HideBanner
	if conf.RemoveTrailingSlash {
		e.Pre(middleware.RemoveTrailingSlash())
	}
	if conf.EnableRecover {
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize: 1 << 10, // 1 KB
		}))
	}
	if conf.EnableGzip {
		e.Use(middleware.Gzip())
	}
	if conf.Context != nil {
		e.Use(conf.Context)
	}
	if conf.EnableLog {
		e.Use(LogMiddleware(conf.IgnoreFavicon))
	}
	if conf.EnableCORS {
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization, "X-Requested-With"},
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}))
	}
	//if conf.EnableJWT {
	//	e.Use(JWTMiddleware(conf.JWTSkipper))
	//}
	if conf.ErrorHandler != nil {
		e.HTTPErrorHandler = conf.ErrorHandler
	}

	if conf.IgnoreFavicon {
		e.GET("/favicon.ico", func(context echo.Context) error {
			return nil
		})
	}

	return e
}

func LogMiddleware(ignoreFavicon bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()

			if ignoreFavicon && request.RequestURI == "/favicon.ico" {
				return next(c)
			}

			// todo: 对文件上传做特殊处理
			body, _ := ioutil.ReadAll(request.Body)
			request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			slog.Trace(fmt.Sprintf("url:%s method:%s body:%s", request.RequestURI, request.Method, string(body)))
			return next(c)
		}
	}
}

func JWTMiddleware(skipper middleware.Skipper) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if skipper(c) {
				return next(c)
			}

			tokenStr := c.Request().Header.Get(echo.HeaderAuthorization)
			if tokenStr == "" {
				return errCode.NewErrorToken()
			}
			data, err := tool.ParseToken(tokenStr)
			if err != nil {
				return errCode.NewErrorToken(err)
			}
			//sysToken := new(models.SysToken)
			//sysToken.Token = tokenStr
			//if err := sysToken.CheckTokenExpireTime(20); err != nil {
			//	slog.Error(fmt.Sprintf("token str is %s", tokenStr), err)
			//	return err
			//}

			cc := c.(*Context)
			if vlu, ok := data.(map[string]interface{}); ok {
				cc.UserID = tool.ToString(vlu["Id"])
			}

			if cc.UserID == "" {
				return errCode.NewErrorToken()
			}

			return next(c)
		}
	}
}
