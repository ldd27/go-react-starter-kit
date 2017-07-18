package routers

import (
	"net/http"

	"github.com/jdongdong/go-react-starter-kit/modules/apiCode"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
	"github.com/jdongdong/go-react-starter-kit/routers/webapi"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	config := middleware.JWTConfig{
		Claims:     &webapi.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	e.Use(middleware.JWTWithConfig(config))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))
	e.HTTPErrorHandler = errHandle

	e.File("/", "static/index.html")

	e.POST("/sysUser/login", webapi.Login)

	api := e.Group("/api", func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			tokenString := c.Request().Header.Get("x-access-token")
			if tokenString == "" {
				return errCode.ErrorInvalidToken
			}
			//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			//	if token.Claims["uid"] != "admin" {
			//		return nil, errors.New("User isvalid")
			//	}
			//	return []byte(SECUREKEY), nil
			//})
			//if err != nil {
			//	return err
			//}
			//c.Set("Uid", token.Claims["uid"])
			return h(c)
		}
	})
	{
		api.GET("/sysLog/page", func(context echo.Context) error {
			//req := new(models.SeaSysLog)
			//panic("dd")
			return context.String(http.StatusOK, "ddd")
			//this.AutoPageDataRs(req.GetPaging())
		})
	}

	return e
}

func errHandle(err error, c echo.Context) {
	res := ResModel{Success: "F", Code: apiCode.FormatApiCode(err), Info: err.Error(), R: ""}
	c.JSON(http.StatusOK, res)
}
