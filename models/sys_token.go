package models

import (
	"time"

	"github.com/jdongdong/go-react-starter-kit/modules/comCode"
	"github.com/jdongdong/go-react-starter-kit/modules/errCode"
)

func (this *SysToken) Insert() error {
	this.Status = comCode.Status_ON
	return _insert(this)
}

func (this *SysToken) UpdateByToken() error {
	this.Status = comCode.Status_OFF
	return errCode.CheckErrorDB(x.Where("token=?", this.Token).Update(this))
}

func (this *SysToken) CheckTokenExpireTime(expireMinute int) error {
	user := new(SysToken)
	user.Token = this.Token
	err := errCode.CheckErrorDataNull(x.Where("token = ? and status = 'aa'", this.Token).Get(user))
	if err != nil {
		return errCode.ErrorInvalidToken
	}

	if (user.CreateTime.Add(time.Minute * time.Duration(expireMinute))).Unix() < time.Now().Unix() {
		return errCode.ErrorInvalidToken
	}

	return nil
}
