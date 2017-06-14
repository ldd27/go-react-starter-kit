package models

import "github.com/jdongdong/go-react-starter-kit/modules/errCode"

func (this *SysRoleUser) GetMenusByUserID() ([]MenuModel, error) {
	menus := make([]SysMenu, 0)
	err := errCode.CheckErrorDB(x.
		Table("sys_menu").
		Where("status='aa' and id in (select distinct menu_id from sys_role_menu where role_id in (select role_id from sys_role_user 	INNER JOIN sys_role on sys_role_user.role_id=sys_role.id and sys_role.status='aa' and sys_role_user.user_id=?))", this.UserId).Find(&menus))
	if err != nil {
		return nil, err
	}
	return this.toMenuModel(menus), nil
}

func (this *SysRoleUser) toMenuModel(menus []SysMenu) []MenuModel {
	rs := make([]MenuModel, len(menus))
	for k, v := range menus {
		item := new(MenuModel)
		item.Id = v.Id
		item.Pid = v.Pid
		item.Name = v.Name
		item.Sort = v.Sort
		item.Router = v.Href
		item.Icon = v.Icon
		rs[k] = *item
	}
	return rs
}
