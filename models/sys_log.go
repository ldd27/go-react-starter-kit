package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type SeaSysLog struct {
	seaModel
	Id     int64
	Type   string
	Title  string
	OpUser string
}

type SysLogModel struct {
	Id     int64
	Type   string
	Title  string
	Info   string
	OpUser string
	OpTime time.Time
	IpAddr string
	ReqUri string
}

type SysLogDtlModel struct {
	SysLogModel `xorm:"extends"`
	OpUserName  string
}

func (this *SeaSysLog) Where(session *xorm.Session) {
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
	session.Table("sys_log").Desc("sys_log.op_time")
}

func (this *SeaSysLog) WhereDtl(session *xorm.Session) {
	session.Join("LEFT", "sys_user", "sys_user.id = sys_log.op_user")
	session.Statement.ColumnStr = "sys_log.*, sys_user.user_name as op_user_name"
}

func (this *SeaSysLog) GetPaging() (interface{}, int64, error) {
	items := make([]SysLogModel, 0, this.Size)
	count, err := this.getPaging(this, new(SysLog), &items)
	return items, count, err
}

func (this *SeaSysLog) GetDtlPaging() (interface{}, int64, error) {
	items := make([]SysLogDtlModel, 0, this.Size)
	count, err := this.getDtlPaging(this, new(SysLog), &items)
	return items, count, err
}

func (this *SysLog) Insert() error {
	item := SysLog(*this)
	return insert(item)
}
