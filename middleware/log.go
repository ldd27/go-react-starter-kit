package middleware

import (
	"fmt"
	"github.com/jdongdong/go-lib/slog"
	"github.com/labstack/echo"
	"io/ioutil"
)

func ReqLogHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			bodyer, _ := request.GetBody()
			body, _ := ioutil.ReadAll(bodyer)

			slog.Trace(fmt.Sprintf("url:%s body:%s", request.RequestURI, string(body)))
			return next(c)
		}
	}
}
