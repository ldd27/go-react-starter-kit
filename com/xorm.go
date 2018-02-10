package com

import (
	"fmt"
	"time"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/slog"
)

type XORMConf struct {
	DBType string
	Host   string
	Port   int
	User   string
	Pwd    string
	DB     string
	conn   string

	Debug        bool
	MaxOpenConns int
	MaxIdleConns int
	MaxLifetime  int
	ShowExecTime bool
	ShowSQL      bool
	LogLevel     core.LogLevel
	TZLocation   *time.Location
}

var defaultXORMConf = XORMConf{
	Debug:        false,
	MaxOpenConns: 2000,
	MaxIdleConns: 1000,
	MaxLifetime:  5,
	ShowExecTime: true,
	ShowSQL:      true,
	LogLevel:     core.LOG_DEBUG,
	TZLocation:   time.Local,
}

func NewXORM(opts ...func(*XORMConf)) *xorm.Engine {
	conf := defaultXORMConf
	for _, o := range opts {
		o(&conf)
	}

	initConn(&conf)

	db, err := xorm.NewEngine(conf.DBType, conf.conn)
	if err != nil {
		slog.Critical("初始化数据库引擎失败", err.Error())
		return nil
	}

	db.SetMapper(core.GonicMapper{})
	db.TZLocation = conf.TZLocation
	db.SetMaxOpenConns(conf.MaxOpenConns)
	db.SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))
	db.ShowExecTime(conf.ShowExecTime)

	db.ShowSQL(conf.ShowSQL)
	db.Logger().SetLevel(conf.LogLevel)

	return db
}

func initConn(option *XORMConf) {
	if option.DBType == "mysql" { //mysql
		option.conn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true", option.User, option.Pwd, option.Host, option.Port, option.DB)
		slog.Info("mysql conn:", option.conn)
	} else if option.DBType == "mssql" {
		option.conn = fmt.Sprintf("driver={SQL Server};server=%s;database=%s;user id=%s;password=%s;", option.Host, option.DB, option.User, option.Pwd)
		slog.Info("mssql conn:", option.conn)
	} else {
		slog.Error("unsupport database")
		return
		//panic("unsupport database")
	}
}
