package models

import (
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/modules/comCode"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

type SeaSysUser struct {
	SeaModel
	Id       string
	LoginKey string
	Password string
	UserType string
	Status   string
}

type SysUserModel struct {
	Id        string
	UserName  string
	LoginName string
	Phone     string
	Email     string
	Sex       string
	UserType  string
	Status    string
	Brief     string
}

type WebLoginUserModel struct {
	Id       string
	UserName string
	Token    string
	Menus    []LeftMenuModel
}

func (this *SeaSysUser) GetLoginByUserID() (*WebLoginUserModel, error) {
	sea := new(SeaSysUser)
	sea.Id = this.Id
	sea.Status = comCode.Status_ON
	user := new(SysUser)
	err := sea._getOne(sea, user)
	if err != nil {
		return nil, err
	}
	loginUser := new(WebLoginUserModel)
	loginUser.UserName = user.UserName
	loginUser.Id = user.Id

	seaRole := new(SysRoleUser)
	seaRole.UserId = user.Id
	menus, err := seaRole.GetMenusByUserID()
	if err != nil {
		return nil, err
	}
	loginUser.Menus = menus

	return loginUser, nil
}

func (this *SeaSysUser) WebLogin() (*WebLoginUserModel, error) {
	sea := new(SeaSysUser)
	sea.Password = tool.MD5(this.Password)
	sea.UserType = comCode.UserType_Web
	sea.LoginKey = this.LoginKey
	user, err := sea.login()
	if err != nil {
		return nil, err
	}
	loginUser := new(WebLoginUserModel)
	loginUser.UserName = user.UserName
	loginUser.Id = user.Id

	//token, err := tool.GenToken(loginUser)
	//if err != nil {
	//	return nil, err
	//}

	seaRole := new(SysRoleUser)
	seaRole.UserId = user.Id
	menus, err := seaRole.GetMenusByUserID()
	if err != nil {
		return nil, err
	}
	loginUser.Menus = menus

	//sysToken := new(SysToken)
	//sysToken.UserId = user.Id
	//sysToken.Token = token
	//err = sysToken.Insert()
	//if err != nil {
	//	return nil, err
	//}
	//loginUser.Token = token

	return loginUser, nil
}

func (this *SeaSysUser) login() (*SysUser, error) {
	user := new(SysUser)
	this.Status = comCode.Status_ON
	err := errCode.CheckErrorInvalidUser(this._getOne(this, user))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (this *SeaSysUser) where(session *xorm.Session) {
	if this.LoginKey != "" {
		session.And("(a.user_name like ? or a.login_name like ? or a.phone like ?)", toLike(this.LoginKey), toLike(this.LoginKey), toLike(this.LoginKey))
	}
	if this.Password != "" {
		session.And("a.password = ?", this.Password)
	}
	if this.Status != "" {
		session.And("a.status = ?", this.Status)
	}
	if this.Id != "" {
		session.And("a.id = ?", this.Id)
	}
	session.Table("sys_user").Alias("a").Desc("a.id")
}

func (this *SeaSysUser) GetPaging() (interface{}, int64, error) {
	items := make([]SysUserModel, 0, this.PageSize)
	count, err := this._getPaging(this, new(SysUser), &items)
	return items, count, err
}

func (this *SysUser) Insert() error {
	this.Id = tool.NewStrID()
	this.Password = tool.MD5("000000")
	return _insert(this)
}

func (this *SysUser) UpdateById() error {
	return _uptByID(this.Id, &this)
}

func (this *SysUser) DeleteById() error {
	return _delByID(this.Id, new(SysUser))
}
