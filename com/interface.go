package com

import "github.com/go-xorm/xorm"

type SeaInterface interface {
	Where(session *xorm.Session)
}

type SeaDtlInterface interface {
	Where(session *xorm.Session)
	WhereDtl(session *xorm.Session)
}

type PagingInterface interface {
	GetPaging() (interface{}, int64, error)
}

type PagingDtlInterface interface {
	GetDtlPaging() (interface{}, int64, error)
}

type InsertInterface interface {
	Insert() error
}

type UptByIdInterface interface {
	UptByID() error
}

type DeleteByIdInterface interface {
	DelByID() error
}

type InsertTransInterface interface {
	InsertTrans() error
}

type UpdateTransInterface interface {
	UptTrans() error
}

type DeleteTransInterface interface {
	DelTrans() error
}
