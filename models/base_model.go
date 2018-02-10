package models

import (
	"fmt"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-react-starter-kit/code/errCode"
	"github.com/jdongdong/go-react-starter-kit/code/setting"
	"github.com/jdongdong/go-react-starter-kit/com"
)

var (
	db     *xorm.Engine
	tables []interface{}
)

type seaModel struct {
	Page         int `json:"page"`
	Size         int `json:"size"`
	seaInterface com.SeaInterface
}

type TreeModel struct {
	Key      string      `json:"key"`
	Title    string      `json:"title"`
	Type     string      `json:"type"`
	Checked  bool        `json:"checked"`
	Children []TreeModel `json:"children"`
}

type LeftMenuModel struct {
	Id     int64  `json:"id"`
	Pid    int64  `json:"pid"`
	MPid   int64  `json:"breadPid"`
	Sort   int    `json:"sort"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
	Router string `json:"router"`
}

func init() {
	conf := func(option *com.XORMConf) {
		option.DBType = setting.DBConf.DBType
		option.Host = setting.DBConf.Host
		option.Port = setting.DBConf.Port
		option.DB = setting.DBConf.Db
		option.User = setting.DBConf.User
		option.Pwd = setting.DBConf.Pwd
		option.LogLevel = core.LOG_INFO
	}
	db = com.NewXORM(conf)
}

func toLike(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

func toPaging(page, size int) (limit, start int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	limit = size
	start = (page - 1) * size
	return limit, start
}

func (s *seaModel) getPaging(i com.SeaInterface, bean interface{}, item interface{}) (int64, error) {
	session := db.NewSession()
	i.Where(session)

	total, err := db.Alias(session.Statement.TableAlias).Where(session.Conds()).Count(bean)
	if err != nil {
		return 0, errCode.NewErrorDB(err, 1)
	}
	err = session.
		Limit(toPaging(s.Page, s.Size)).
		Find(item)
	if err != nil {
		return 0, errCode.NewErrorDB(err, 1)
	}
	return total, nil
}

func (s *seaModel) getDtlPaging(i com.SeaDtlInterface, bean interface{}, item interface{}) (int64, error) {
	session := db.NewSession()
	i.Where(session)
	i.WhereDtl(session)

	total, err := db.Alias(session.Statement.TableAlias).Where(session.Conds()).Count(bean)
	if err != nil {
		return 0, errCode.NewErrorDB(err, 1)
	}
	err = session.
		Limit(toPaging(s.Page, s.Size)).
		Find(item)
	if err != nil {
		return 0, errCode.NewErrorDB(err, 1)
	}
	return total, nil
}

func (s *seaModel) getAll(i com.SeaInterface, item interface{}) error {
	session := db.NewSession()
	i.Where(session)
	err := session.Find(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	return nil
}

func (s *seaModel) getDtlAll(i com.SeaDtlInterface, item interface{}) error {
	session := db.NewSession()
	i.Where(session)
	i.WhereDtl(session)
	err := session.Find(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	return nil
}

func (s *seaModel) getOne(i com.SeaInterface, item interface{}) error {
	session := db.NewSession()
	i.Where(session)
	has, err := session.Get(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	if !has {
		return errCode.NewErrorNoRecord(1)
	}
	return nil
}

func (s *seaModel) getDtlOne(i com.SeaDtlInterface, item interface{}) error {
	session := db.NewSession()
	i.Where(session)
	i.WhereDtl(session)
	has, err := session.Get(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	if !has {
		return errCode.NewErrorNoRecord(1)
	}
	return nil
}

func insert(item interface{}) error {
	temp := reflect.ValueOf(item).Elem()
	if temp.FieldByName("CreateTime").CanSet() {
		temp.FieldByName("CreateTime").Set(reflect.ValueOf(time.Now()))
	}
	if temp.FieldByName("UpdateTime").CanSet() {
		temp.FieldByName("UpdateTime").Set(reflect.ValueOf(time.Now()))
	}
	count, err := db.Insert(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	if count == 0 {

	}
	return nil
}

func uptByID(id interface{}, item interface{}) error {
	temp := reflect.ValueOf(item).Elem()
	if temp.FieldByName("UpdateTime").CanSet() {
		temp.FieldByName("UpdateTime").Set(reflect.ValueOf(time.Now()))
	}
	count, err := db.ID(id).Update(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	if count == 0 {
		return errCode.NewErrorNoRecord(1)
	}
	return nil
}

func uptIsNullByID(id interface{}, item interface{}, columns ...string) error {
	temp := reflect.ValueOf(item).Elem()
	if temp.FieldByName("UpdateTime").CanSet() {
		temp.FieldByName("UpdateTime").Set(reflect.ValueOf(time.Now()))
	}
	count, err := db.Cols(columns...).ID(id).Update(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	if count == 0 {
		return errCode.NewErrorNoRecord(1)
	}
	return nil
}

func delByID(id interface{}, item interface{}) error {
	count, err := db.Id(id).Delete(item)
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	if count == 0 {
		return errCode.NewErrorNoRecord(1)
	}
	return nil
}

func trans(fun func(session *xorm.Session) error) error {
	session, err := startTrans()
	defer session.Close()
	if err != nil {
		return err
	}

	err = fun(session)
	if err != nil {
		return checkRollback(session, err)
	}

	return commitTrans(session)
}

func startTrans() (*xorm.Session, error) {
	session := db.NewSession()
	err := session.Begin()
	if err != nil {
		return nil, errCode.NewErrorDB(err, 1)
	}
	return session, err
}

func commitTrans(session *xorm.Session) error {
	err := session.Commit()
	if err != nil {
		return errCode.NewErrorDB(err, 1)
	}
	return nil
}

func checkRollback(session *xorm.Session, err error) error {
	if err != nil {
		session.Rollback()
	}
	return err
}
