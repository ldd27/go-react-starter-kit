package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/common/comStruct"
	"github.com/jdongdong/go-react-starter-kit/common/errCode"
	"github.com/jdongdong/go-react-starter-kit/common/logCode"
)

func GetDictIndexTree(c *comStruct.CustomContext) error {
	req := new(models.SeaDictIndex)
	return c.DataRs(req.GetTree())
}

func GetValidDictItemsByCode(c *comStruct.CustomContext) error {
	req := &models.SeaDictItem{
		DictCode: c.QueryParam("DictCode"),
		Valid:    true,
	}
	return c.DataRs(req.GetValidItemsByCode())
}

func GetDictItems(c *comStruct.CustomContext) error {
	req := &models.SeaDictItem{
		DictCode: c.QueryParam("DictCode"),
	}
	if req.DictCode == "" {
		return errCode.ErrorParams
	}
	return c.DataRs(req.GetAll())
}

func AddDictItem(c *comStruct.CustomContext) error {
	req := &models.DictItem{}
	err := c.BindEx(req)
	if err != nil {
		return err
	}
	return c.AutoInsertRs(logCode.AddDictItem, req)
}

func EditDictItem(c *comStruct.CustomContext) error {
	req := &models.DictItem{}
	err := c.BindEx(req)
	if err != nil {
		return err
	}
	return c.AutoUpdateRs(logCode.EditDictItem, req)
}

func DelDictItem(c *comStruct.CustomContext) error {
	id, err := c.ToInt("Id")
	if err != nil {
		return err
	}

	req := &models.DictItem{Id: id}
	return c.AutoDeleteRs(logCode.DelDictItem, req)
}
