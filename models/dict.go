package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
)

type SeaDictIndex struct {
	SeaModel
	//Status string
}

func (this *SeaDictIndex) where(session *xorm.Session) {
	//if this.Status != "" {
	//	session.And("dict_index.status = ?", this.Status)
	//}
	session.Table("dict_index").Asc("dict_index.dict_code")
}

func (this *SeaDictIndex) GetTree() ([]TreeModel, error) {
	items := make([]DictIndex, 0)
	//this.Status = comCode.Status_ON
	err := this._getAll(this, &items)
	if err != nil {
		return nil, err
	}
	types := make([]string, 0)
	for _, v := range items {
		isExist := false
		for _, vv := range types {
			if vv == v.DictType {
				isExist = true
			}
		}
		if !isExist {
			types = append(types, v.DictType)
		}
	}
	rs := make([]TreeModel, len(types))
	for k, v := range types {
		node := new(TreeModel)
		node.Title = v
		node.Key = v
		node.Type = "type"
		node.Children = make([]TreeModel, 0)
		for _, vv := range items {
			if vv.DictType == v {
				child := new(TreeModel)
				child.Key = fmt.Sprintf("%s-%s", vv.DictCode, vv.IsSys)
				child.Title = vv.DictName
				node.Children = append(node.Children, *child)
			}
		}
		rs[k] = *node
	}
	return rs, nil
}

type SeaDictItem struct {
	SeaModel
	DictCode string
	Valid    bool `json:"-"`
}

func (this *SeaDictItem) where(session *xorm.Session) {
	if this.DictCode != "" {
		session.And("a.dict_code = ?", this.DictCode)
	}
	if this.Valid {
		session.And("a.status = 'aa' and a.dict_code in (select dict_code from dict_index and dict_index.status='aa' and dict_index.dict_code=?)", this.DictCode)
	}
	session.Table("dict_item").Alias("a").Asc("dict_index.item_code")
}

func (this *SeaDictItem) GetValidItemsByCode() ([]DictItem, error) {
	items := make([]DictItem, 0)
	this.Valid = true
	return items, this._getAll(this, &items)
}

func (this *SeaDictItem) GetAll() ([]DictItem, error) {
	items := make([]DictItem, 0)
	return items, this._getAll(this, &items)
}
