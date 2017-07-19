package middleware

import (
	"fmt"

	"io/ioutil"

	"github.com/jdongdong/go-lib/slog"
	"github.com/labstack/echo"
)

func ReqLogHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			request := c.Request()
			body, _ := ioutil.ReadAll(request.Body)

			slog.Trace(fmt.Sprintf("url:%s body:%s", request.RequestURI, string(body)))
			return next(c)
		}
	}
}
