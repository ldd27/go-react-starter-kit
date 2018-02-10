package com

import "reflect"

const (
	fieldPage     = "Page"
	fieldSize     = "Size"
	fieldCreateBy = "CreateBy"
	fieldChangeBy = "UpdateBy"
)

func (c *Context) AutoRsPage(req PagingInterface, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldPage).CanSet() {
		temp.FieldByName(fieldPage).Set(reflect.ValueOf(c.GetPageIndex()))
	}
	if temp.FieldByName(fieldSize).CanSet() {
		temp.FieldByName(fieldSize).Set(reflect.ValueOf(c.GetPageSize()))
	}
	return c.RsPage(req.GetPaging())
}

func (c *Context) AutoRsPageDtl(req PagingDtlInterface, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldPage).CanSet() {
		temp.FieldByName(fieldPage).Set(reflect.ValueOf(c.GetPageIndex()))
	}
	if temp.FieldByName(fieldSize).CanSet() {
		temp.FieldByName(fieldSize).Set(reflect.ValueOf(c.GetPageSize()))
	}
	return c.RsPage(req.GetDtlPaging())
}

func (c *Context) AutoRsInsert(req InsertInterface, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldCreateBy).CanSet() {
		temp.FieldByName(fieldCreateBy).Set(reflect.ValueOf(c.UserID))
	}
	if temp.FieldByName(fieldChangeBy).CanSet() {
		temp.FieldByName(fieldChangeBy).Set(reflect.ValueOf(c.UserID))
	}
	c.WriteLogFunc(req, logTag...)
	return c.Rs(req.Insert())
}

func (c *Context) AutoRsUpdate(req UptByIdInterface, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldChangeBy).CanSet() {
		temp.FieldByName(fieldChangeBy).Set(reflect.ValueOf(c.UserID))
	}
	c.WriteLogFunc(req, logTag...)
	return c.Rs(req.UptByID())
}

func (c *Context) AutoRsDelete(req DeleteByIdInterface, logTag ...string) error {
	c.WriteLogFunc(req, logTag...)
	return c.Rs(req.DelByID())
}

func (c *Context) AutoRsInsertTrans(req InsertTransInterface, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldCreateBy).CanSet() {
		temp.FieldByName(fieldCreateBy).Set(reflect.ValueOf(c.UserID))
	}
	if temp.FieldByName(fieldChangeBy).CanSet() {
		temp.FieldByName(fieldChangeBy).Set(reflect.ValueOf(c.UserID))
	}
	c.WriteLogFunc(req, logTag...)
	return c.Rs(req.InsertTrans())
}

func (c *Context) AutoRsUpdateTrans(req UpdateTransInterface, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldChangeBy).CanSet() {
		temp.FieldByName(fieldChangeBy).Set(reflect.ValueOf(c.UserID))
	}
	c.WriteLogFunc(req, logTag...)
	return c.Rs(req.UptTrans())
}

func (c *Context) AutoRsDeleteTrans(req DeleteTransInterface, logTag ...string) error {
	c.WriteLogFunc(req, logTag...)
	return c.Rs(req.DelTrans())
}

func (c *Context) AutoRs(req interface{}, err error, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldCreateBy).CanSet() {
		temp.FieldByName(fieldCreateBy).Set(reflect.ValueOf(c.UserID))
	}
	if temp.FieldByName(fieldChangeBy).CanSet() {
		temp.FieldByName(fieldChangeBy).Set(reflect.ValueOf(c.UserID))
	}
	c.WriteLogFunc(req, logTag...)
	return c.Rs(err)
}

func (c *Context) AutoRsData(req interface{}, data interface{}, err error, logTag ...string) error {
	temp := reflect.ValueOf(req).Elem()
	if temp.FieldByName(fieldCreateBy).CanSet() {
		temp.FieldByName(fieldCreateBy).Set(reflect.ValueOf(c.UserID))
	}
	if temp.FieldByName(fieldChangeBy).CanSet() {
		temp.FieldByName(fieldChangeBy).Set(reflect.ValueOf(c.UserID))
	}
	c.WriteLogFunc(req, logTag...)
	return c.RsData(data, err)
}
