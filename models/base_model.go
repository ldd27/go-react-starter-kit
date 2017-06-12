package models

import (
	"fmt"
	"time"

	"reflect"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

var (
	x      *xorm.Engine
	tables []interface{}
	DbCfg  struct {
		Host, Port, User, Pwd, Db string
	}
)

type SeaModel struct {
	PageIndex    int `json:"pageIndex"`
	PageSize     int `json:"pageSize"`
	seaInterface SeaInterface
}

type TreeModel struct {
	Key      string      `json:"key"`
	Title    string      `json:"title"`
	Type     string      `json:"type"`
	Checked  bool        `json:"checked"`
	Children []TreeModel `json:"children"`
}

type MenuModel struct {
	Key   string      `json:"key"`
	Href  string      `json:"href"`
	Name  string      `json:"name"`
	Icon  string      `json:"icon"`
	Child []MenuModel `json:"child"`
}

type SeaInterface interface {
	where(session *xorm.Session)
}

type PagingInterface interface {
	GetPaging() (interface{}, int64, error)
}

type InsertInterface interface {
	Insert() error
}

type UpdateByIdInterface interface {
	UpdateById() error
}

type DeleteByIdInterface interface {
	DeleteById() error
}

type InsertTransInterface interface {
	InsertTrans() error
}

type UpdateTransInterface interface {
	UpdateTrans() error
}

type DeleteTransInterface interface {
	DeleteTrans() error
}

func init() {
	tables = append(tables, new(SysUser))
	LoadConfig()
	err := NewEngine()
	slog.Error(err)
	err = x.Ping()
	slog.Error(err)
}

func LoadConfig() {
	DbCfg.Host = beego.AppConfig.DefaultString("mysqlhost", "localhost")
	DbCfg.Port = beego.AppConfig.DefaultString("mysqlport", "3306")
	DbCfg.User = beego.AppConfig.DefaultString("mysqluser", "root")
	DbCfg.Pwd = beego.AppConfig.String("mysqlpass")
	DbCfg.Db = beego.AppConfig.String("mysqldb")
}

func getEngine() (*xorm.Engine, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", DbCfg.User, DbCfg.Pwd, DbCfg.Host, DbCfg.Port, DbCfg.Db)
	slog.Info("conn:" + connStr)
	return xorm.NewEngine("mysql", connStr)
}

func SetEngine() (err error) {
	x, err = getEngine()
	if err != nil {
		return fmt.Errorf("fail to connect to database: %v", err)
	}
	x.SetMapper(core.GonicMapper{})
	x.TZLocation = time.Local
	x.SetMaxOpenConns(2000)
	x.SetMaxIdleConns(1000)
	x.DB().SetConnMaxLifetime(time.Second * 5)
	x.ShowExecTime(true)

	x.ShowSQL(true)
	x.Logger().SetLevel(core.LOG_ERR)
	return nil
}

func NewEngine() (err error) {
	if err = SetEngine(); err != nil {
		return err
	}
	//同步数据库结构
	//if err = x.StoreEngine("InnoDB").Sync2(tables...); err != nil {
	//	return fmt.Errorf("sync database struct error: %v", err)
	//}

	return nil
}

func toLike(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

func toPaging(pageIndex, pageSize int) (limit, start int) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	limit = pageSize
	start = (pageIndex - 1) * pageSize
	return limit, start
}

func (this *SeaModel) _getPaging(i SeaInterface, bean interface{}, item interface{}) (int64, error) {
	session := x.NewSession()
	i.where(session)

	total, err := session.Count(bean)
	if err := errCode.CheckErrorDB(err); err != nil {
		return 0, err
	}
	session2 := x.NewSession()
	i.where(session2)
	err = session2.
		Limit(toPaging(this.PageIndex, this.PageSize)).
		Find(item)
	return total, errCode.CheckErrorDB(err)
}

func (this *SeaModel) _getAll(i SeaInterface, item interface{}) error {
	session := x.NewSession()
	i.where(session)
	return errCode.CheckErrorDB(session.Find(item))
}

func (this *SeaModel) _getOne(i SeaInterface, item interface{}) error {
	session := x.NewSession()
	i.where(session)
	return errCode.CheckErrorDataNull(session.Get(item))
}

func _insert(item interface{}) error {
	temp := reflect.ValueOf(item).Elem()
	if temp.FieldByName("CreateTime").CanSet() {
		temp.FieldByName("CreateTime").Set(reflect.ValueOf(time.Now()))
	}
	if temp.FieldByName("UpdateTime").CanSet() {
		temp.FieldByName("UpdateTime").Set(reflect.ValueOf(time.Now()))
	}
	_, err := x.Insert(item)
	return errCode.CheckErrorDB(err)
}

func _uptByID(id interface{}, item interface{}) error {
	temp := reflect.ValueOf(item).Elem()
	if temp.FieldByName("UpdateTime").CanSet() {
		temp.FieldByName("UpdateTime").Set(reflect.ValueOf(time.Now()))
	}
	_, err := x.Omit("create_time").ID(id).Update(item)
	return errCode.CheckErrorDB(err)
}

func _delByID(id interface{}, item interface{}) error {
	_, err := x.Id(id).Delete(item)
	return errCode.CheckErrorDB(err)
}

func _trans(fun func(session *xorm.Session) error) error {
	session, err := _startTrans()
	defer session.Close()
	if err != nil {
		return err
	}

	err = fun(session)
	if err != nil {
		return _checkRollback(session, err)
	}

	return _commitTrans(session)
}

func _startTrans() (*xorm.Session, error) {
	session := x.NewSession()
	err := errCode.CheckErrorDB(session.Begin())
	return session, err
}

func _commitTrans(session *xorm.Session) error {
	return errCode.CheckErrorDB(session.Commit())
}

func _checkRollback(session *xorm.Session, err error) error {
	if err = errCode.CheckErrorDB(err); err != nil {
		session.Rollback()
		return err
	} else {
		return nil
	}
}
