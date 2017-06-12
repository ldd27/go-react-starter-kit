package models

import (
	"time"
)

type SysLog struct {
	Id     int64     `xorm:"pk autoincr BIGINT(20)"`
	Type   string    `xorm:"VARCHAR(20)"`
	Title  string    `xorm:"VARCHAR(100)"`
	Info   string    `xorm:"VARCHAR(2000)"`
	OpUser string    `xorm:"VARCHAR(32)"`
	OpTime time.Time `xorm:"DATETIME"`
	IpAddr string    `xorm:"VARCHAR(50)"`
	ReqUri string    `xorm:"VARCHAR(255)"`
}

type SysMenu struct {
	Id         string    `xorm:"not null pk VARCHAR(32)"`
	Pid        string    `xorm:"not null VARCHAR(32)"`
	Name       string    `xorm:"not null VARCHAR(50)"`
	Href       string    `xorm:"VARCHAR(100)"`
	Icon       string    `xorm:"VARCHAR(20)"`
	Sort       int       `xorm:"INT(11)"`
	Type       string    `xorm:"not null VARCHAR(10)"`
	Status     string    `xorm:"not null VARCHAR(2)"`
	CreateBy   string    `xorm:"VARCHAR(32)"`
	CreateTime time.Time `xorm:"DATETIME"`
	UpdateBy   string    `xorm:"VARCHAR(32)"`
	UpdateTime time.Time `xorm:"DATETIME"`
}

type SysRole struct {
	Id         string    `xorm:"not null pk VARCHAR(32)"`
	RoleName   string    `xorm:"unique VARCHAR(50)"`
	Status     string    `xorm:"VARCHAR(2)"`
	CreateBy   string    `xorm:"VARCHAR(32)"`
	CreateTime time.Time `xorm:"DATETIME"`
	UpdateBy   string    `xorm:"VARCHAR(32)"`
	UpdateTime time.Time `xorm:"DATETIME"`
}

type SysRoleMenu struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	RoleId string `xorm:"not null unique(role_id) VARCHAR(32)"`
	MenuId string `xorm:"not null unique(role_id) VARCHAR(32)"`
}

type SysRoleUser struct {
	Id     int    `xorm:"not null pk autoincr INT(11)"`
	RoleId string `xorm:"not null unique(role_id) VARCHAR(32)"`
	UserId string `xorm:"not null unique(role_id) VARCHAR(32)"`
}

type SysToken struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId     string    `xorm:"VARCHAR(32)"`
	Token      string    `xorm:"TEXT"`
	Status     string    `xorm:"VARCHAR(2)"`
	CreateTime time.Time `xorm:"DATETIME"`
}

type SysUser struct {
	Id         string    `xorm:"not null pk VARCHAR(32)"`
	UserName   string    `xorm:"not null VARCHAR(20)"`
	Password   string    `xorm:"not null VARCHAR(100)"`
	LoginName  string    `xorm:"not null unique VARCHAR(30)"`
	Phone      string    `xorm:"not null unique VARCHAR(20)"`
	Email      string    `xorm:"VARCHAR(50)"`
	Sex        string    `xorm:"VARCHAR(1)"`
	UserType   string    `xorm:"not null VARCHAR(10)"`
	Status     string    `xorm:"not null VARCHAR(2)"`
	Brief      string    `xorm:"VARCHAR(100)"`
	CreateBy   string    `xorm:"VARCHAR(32)"`
	CreateTime time.Time `xorm:"DATETIME"`
	UpdateBy   string    `xorm:"VARCHAR(32)"`
	UpdateTime time.Time `xorm:"DATETIME"`
}
