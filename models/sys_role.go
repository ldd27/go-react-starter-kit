package models

import (
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

func (this *SysRoleUser) GetMenusByUserID() ([]MenuModel, error) {
	menus := make([]SysMenu, 0)
	err := errCode.CheckErrorDB(x.
		Table("sys_menu").
		Where("status='aa' and id in (select distinct menu_id from sys_role_menu where role_id in (select role_id from sys_role_user 	INNER JOIN sys_role on sys_role_user.role_id=sys_role.id and sys_role.status='aa' and sys_role_user.user_id=?))", this.UserId))
	if err != nil {
		return nil, err
	}
	return this.pkgMenu(menus, "root"), nil
}

func (this *SysRoleUser) pkgMenu(menus []SysMenu, pid string) []MenuModel {
	count := 0
	for _, vlu := range menus {
		if vlu.Pid == pid {
			count++
		}
	}
	if count == 0 {
		return nil
	}

	key := 0
	rs := make([]MenuModel, count)
	for _, vlu := range menus {
		if vlu.Pid == pid {
			node := new(MenuModel)
			node.Key = tool.ToString(vlu.Id)
			node.Href = vlu.Href
			node.Icon = vlu.Icon
			node.Name = vlu.Name
			node.Child = this.pkgMenu(menus, vlu.Id)
			rs[key] = *node
			key++
		}
	}
	return rs
}
