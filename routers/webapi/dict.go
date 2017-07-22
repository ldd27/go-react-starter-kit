package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/models"
	"github.com/jdongdong/go-react-starter-kit/modules/comStruct"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
	"github.com/jdongdong/go-react-starter-kit/modules/logCode"
)

func GetDictIndexTree(c *comStruct.CustomContext) error {
	req := new(models.SeaDictIndex)
	return c.AutoDataRs(req.GetTree())
}

func GetValidDictItemsByCode(c *comStruct.CustomContext) error {
	req := &models.SeaDictItem{
		DictCode: c.QueryParam("DictCode"),
		Valid:    true,
	}
	return c.AutoDataRs(req.GetValidItemsByCode())
}

func GetDictItems(c *comStruct.CustomContext) error {
	req := &models.SeaDictItem{
		DictCode: c.QueryParam("DictCode"),
	}
	if req.DictCode == "" {
		return errCode.ErrorParams
	}
	return c.AutoDataRs(req.GetAll())
}

func AddDictItem(c *comStruct.CustomContext) error {
	req := &models.DictItem{}
	err := c.BindEx(req)
	if err != nil {
		return err
	}
	return c.InsertRs(logCode.AddDictItem, req)
}

func EditDictItem(c *comStruct.CustomContext) error {
	req := &models.DictItem{}
	err := c.BindEx(req)
	if err != nil {
		return err
	}
	return c.UpdateRs(logCode.EditDictItem, req)
}

func DelDictItem(c *comStruct.CustomContext) error {
	id, err := c.ToInt("Id")
	if err != nil {
		return err
	}

	req := &models.DictItem{Id: id}
	return c.InsertRs(logCode.DelDictItem, req)
}
