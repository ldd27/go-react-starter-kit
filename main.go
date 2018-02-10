package main

import (
	"fmt"

	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/code/setting"
	"github.com/jdongdong/go-react-starter-kit/routers"
)

func main() {
	tool.SetTokenSecretKey(setting.TokenSecret)
	router := routers.Init()
	go func() {
		router.StartAutoTLS(":443")
	}()
	slog.Error(router.Start(fmt.Sprintf(":%d", setting.Port)))
}
