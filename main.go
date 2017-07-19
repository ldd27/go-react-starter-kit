package main

import (
	"fmt"

	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-react-starter-kit/pkg/setting"
	"github.com/jdongdong/go-react-starter-kit/routers"
	"github.com/jdongdong/go-lib/tool"
)

func main() {
	tool.SetTokenSecretKey(setting.TokenSecret)
	router := routers.Init()
	slog.Error(router.Start(fmt.Sprintf(":%d", setting.Port)))
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//
	////解决跨域问题
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins:     []string{beego.AppConfig.String("accessdomain")},
	//	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
	//	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "token", "X-Requested-With"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))
	//beego.ErrorController(&controllers.ErrorController{})
	//beego.Run()
}
