package setting

import (
	"fmt"

	"github.com/go-ini/ini"
	"github.com/jdongdong/go-lib/slog"
)

var (
	Cfg         *ini.File
	Domain      string
	Port        int
	TokenSecret string

	MysqlCfg struct {
		Host, User, Pwd, Db, Conn string
		Port                      int
	}
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		slog.Error(err)
	}
	Domain = Cfg.Section("").Key("domain").MustString("*")
	Port = Cfg.Section("").Key("port").MustInt(8080)
	TokenSecret = Cfg.Section("").Key("tokenSecret").MustString("sadjaksdhajks")

	//mysql
	MysqlCfg.Host = Cfg.Section("mysql").Key("host").MustString("localhost")
	MysqlCfg.Port = Cfg.Section("mysql").Key("port").MustInt(3306)
	MysqlCfg.User = Cfg.Section("mysql").Key("user").MustString("root")
	MysqlCfg.Pwd = Cfg.Section("mysql").Key("pass").String()
	MysqlCfg.Db = Cfg.Section("mysql").Key("db").String()
	MysqlCfg.Conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", MysqlCfg.User, MysqlCfg.Pwd, MysqlCfg.Host, MysqlCfg.Port, MysqlCfg.Db)
	slog.Info("mysql conn:", MysqlCfg.Conn)
}
