package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type SeaSysLog struct {
	SeaModel
	Id     int64
	Type   string
	Title  string
	OpUser string
}

type SysLogModel struct {
	Id         int64
	Type       string
	Title      string
	Info       string
	OpUser     string
	OpTime     time.Time
	IpAddr     string
	ReqUri     string
	OpUserName string
}

func (this *SeaSysLog) where(session *xorm.Session) {
	if this.Id != 0 {
		session.And("sys_log.id = ?", this.Id)
	}
	if this.Type != "" {
		session.And("sys_log.type like ?", this.Type)
	}
	if this.Title != "" {
		session.And("sys_log.title like ?", toLike(this.Title))
	}
	if this.OpUser != "" {
		session.And("sys_log.op_user = ?", this.OpUser)
	}
	session.
		Table("sys_log").
		Join("LEFT", "sys_user", "sys_user.id = sys_log.op_user").
		Desc("sys_log.op_time")

	session.Statement.ColumnStr = "sys_log.*, sys_user.user_name as op_user_name"
}

func (this *SeaSysLog) GetPaging() (interface{}, int64, error) {
	items := make([]SysLogModel, this.PageSize)
	count, err := this._getPaging(this, new(SysLog), &items)
	return items, count, err
}

func (this *SysLog) Insert() error {
	item := SysLog(*this)
	return _insert(item)
}
