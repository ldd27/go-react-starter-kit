package routers

import (
	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-react-starter-kit/com"
	"github.com/jdongdong/go-react-starter-kit/routers/webapi"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {
	conf := func(option *com.EchoConf) {
		option.Context = com.NewContext(func(info interface{}, logTag ...string) {
			slog.Info(logTag, info)
		})
		option.ErrorHandler = com.NewErrorHandle()
	}

	e := com.NewEcho(conf)

	e.File("/", "static/index.html")
	e.Static("/static", "static")

	api := e.Group("/webApi", com.JWTMiddleware(func(c echo.Context) bool {
		// 返回true表示跳过该路由，不进行检查
		request := c.Request()
		if request.RequestURI == "/webApi/sysLog/page" {
			return true
		}
		return false
	}))
	{
		com.GET(api, "/sysUser/checkIsLogin", webapi.CheckIsLogin)
		com.POST(api, "/sysUser/login", webapi.Login)
		com.POST(api, "/sysUser/logout", webapi.Logout)

		com.GET(api, "/sysLog/page", webapi.GetPagingSysLog)

		com.GET(api, "/dictIndex/tree", webapi.GetDictIndexTree)
		com.GET(api, "/dictItem/valid", webapi.GetValidDictItemsByCode)
		com.GET(api, "/dictItem", webapi.GetDictItems)
		com.POST(api, "/dictItem", webapi.AddDictItem)
		com.PUT(api, "/dictItem", webapi.EditDictItem)
		com.GET(api, "/dictItem", webapi.DelDictItem)
	}

	return e
}
