package webapi

import (
	"github.com/jdongdong/go-react-starter-kit/code/errCode"
	"github.com/jdongdong/go-react-starter-kit/com"
	"github.com/jdongdong/go-react-starter-kit/models"
)

func GetDictIndexTree(c *com.Context) error {
	req := new(models.SeaDictIndex)
	return c.RsData(req.GetTree())
}

func GetValidDictItemsByCode(c *com.Context) error {
	req := &models.SeaDictItem{
		DictCode: c.QueryParam("DictCode"),
		Valid:    true,
	}
	return c.RsData(req.GetValidItemsByCode())
}

func GetDictItems(c *com.Context) error {
	req := &models.SeaDictItem{
		DictCode: c.QueryParam("DictCode"),
	}
	if req.DictCode == "" {
		return errCode.NewErrorParams()
	}
	return c.RsData(req.GetAll())
}

func AddDictItem(c *com.Context) error {
	req := &models.DictItem{}
	err := c.BindEx(req)
	if err != nil {
		return err
	}
	return c.AutoRsInsert(req, "添加字典明细")
}

func EditDictItem(c *com.Context) error {
	req := &models.DictItem{}
	err := c.BindEx(req)
	if err != nil {
		return err
	}
	return c.AutoRsUpdate(req, "编辑字典明细")
}

func DelDictItem(c *com.Context) error {
	id, err := c.ToInt("Id")
	if err != nil {
		return err
	}

	req := &models.DictItem{Id: id}
	return c.AutoRsDelete(req, "删除字典明细")
}
