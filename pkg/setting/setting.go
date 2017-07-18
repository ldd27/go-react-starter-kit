package setting

import (
	"github.com/go-ini/ini"
	"github.com/jdongdong/go-lib/slog"
)

var (
	Cfg    *ini.File
	Domain string
	Port   int
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		slog.Error(err)
	}
	Domain = Cfg.Section("").Key("domain").MustString("*")
	Port = Cfg.Section("").Key("port").MustInt(8080)
}
