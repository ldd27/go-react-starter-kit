package routers

import (
	"net/http"

	"fmt"

	"github.com/jdongdong/go-lib/slog"
	cus "github.com/jdongdong/go-react-starter-kit/middleware"
	"github.com/jdongdong/go-react-starter-kit/modules/apiCode"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
	"github.com/jdongdong/go-react-starter-kit/routers/webapi"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//func init() {
//	ns := beego.NewNamespace("/v1",
//		beego.NSNamespace("/sysUser", beego.NSInclude(&controllers.SysUserController{})),
//		beego.NSNamespace("/sysLog", beego.NSInclude(&controllers.SysLogController{})),
//	)
//	beego.AddNamespace(ns)
//}

func Init() *echo.Echo {
	e := echo.New()

	e.Debug = true

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	//e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(cusContext)
	e.Use(cus.ReqLogHandler())

	//config := middleware.JWTConfig{
	//	Claims:     &webapi.JwtCustomClaims{},
	//	SigningKey: []byte("secret"),
	//}
	//e.Use(middleware.JWTWithConfig(config))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	e.HTTPErrorHandler = errHandle

	e.File("/", "static/index.html")
	e.Static("/static", "static")

	//e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	//	return func(c echo.Context) (err error) {
	//		if c.Request().Method == "Options" {
	//			c.NoContent(200)
	//			return
	//		}
	//		return next(c)
	//	}
	//})

	api := e.Group("/webApi", cus.JwtHandler())
	{
		api.GET("/sysUser/checkIsLogin", handleCusContext(webapi.CheckIsLogin))
		api.POST("/sysUser/login", handleCusContext(webapi.Login))
		api.POST("/sysUser/logout", handleCusContext(webapi.Logout))

		api.GET("/sysLog/page", handleCusContext(webapi.GetPagingSysLog))

		api.GET("/dictIndex/tree", handleCusContext(webapi.GetDictIndexTree))
		api.GET("/dictItem/valid", handleCusContext(webapi.GetValidDictItemsByCode))
		api.GET("/dictItem", handleCusContext(webapi.GetDictItems))
		api.POST("/dictItem", handleCusContext(webapi.AddDictItem))
		api.PUT("/dictItem", handleCusContext(webapi.EditDictItem))
		api.DELETE("/dictItem", handleCusContext(webapi.DelDictItem))
	}

	return e
}

func cusContext(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := &comStruct.CustomContext{Context: c}
		return h(cc)
	}
}

func handleCusContext(h comStruct.CusHandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*comStruct.CustomContext)
		return h(cc)
	}
}

func errHandle(err error, c echo.Context) {
	res := comStruct.ResModel{Success: "F", Code: apiCode.FormatApiCode(err), Info: err.Error(), R: ""}
	slog.Debug(fmt.Sprintf("url:%s success:%s code:%d info:%s", c.Request().RequestURI, res.Success, res.Code, err.Error()))
	c.JSON(http.StatusOK, res)
}
