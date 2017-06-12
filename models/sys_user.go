package models

import (
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/tool"
	"github.com/jdongdong/go-react-starter-kit/modules/comCode"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

type SeaSysUser struct {
	SeaModel
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
	Id        string
	UserName  string
	LoginName string
	Phone     string
	Email     string
	Sex       string
	Brief     string
	Menus     []MenuModel
}

func (this *SeaSysUser) WebLogin() (*WebLoginUserModel, error) {
	sea := new(SeaSysUser)
	sea.Password = tool.MD5(this.Password)
	sea.UserType = comCode.UserType_Web
	user, err := sea.login()
	if err != nil {
		return nil, err
	}
	loginUser := new(WebLoginUserModel)
	loginUser.Brief = user.Brief
	loginUser.Email = user.Email
	loginUser.Id = user.Id
	loginUser.LoginName = user.LoginName
	loginUser.Phone = user.Phone
	loginUser.Sex = user.Sex
	loginUser.UserName = user.UserName

	return nil, nil
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
