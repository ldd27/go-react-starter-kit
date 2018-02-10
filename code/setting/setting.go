package setting

import (
	"github.com/go-ini/ini"
	"github.com/jdongdong/go-lib/slog"
)

var (
	Cfg         *ini.File
	Domain      string
	Port        int
	TokenSecret string

	DBConf struct {
		DBType, Host, User, Pwd, Db string
		Port                        int
	}
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		slog.Error(err)
		return
	}
	Domain = getConf("domain").MustString("*")
	Port = getConf("port").MustInt(8080)
	TokenSecret = getConf("tokenSecret").MustString("MRNdiNmIkOffJYqs")

	getDBConf()
}

func getDBConf() {
	DBConf.DBType = getConf("dbType").MustString("mysql")
	DBConf.Host = getConf("host", DBConf.DBType).MustString("localhost")
	DBConf.Port = getConf("port", DBConf.DBType).MustInt(3306)
	DBConf.User = getConf("user", DBConf.DBType).MustString("root")
	DBConf.Pwd = getConf("pass", DBConf.DBType).String()
	DBConf.Db = getConf("db", DBConf.DBType).String()
}

func getConf(key string, sections ...string) *ini.Key {
	section := ""
	if len(sections) > 0 {
		section = sections[0]
	}
	return Cfg.Section(section).Key(key)
}
