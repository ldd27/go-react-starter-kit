package models

import (
	"time"

	"github.com/jdongdong/go-react-starter-kit/code/comCode"
	"github.com/jdongdong/go-react-starter-kit/code/errCode"
)

func (this *SysToken) Insert() error {
	this.Status = comCode.Status_ON
	return insert(this)
}

func (this *SysToken) UpdateByToken() error {
	this.Status = comCode.Status_OFF
	count, err := db.Where("token=?", this.Token).Update(this)
	if err != nil {
		return errCode.NewErrorDB(err)
	}
	if count == 0 {
		return errCode.NewErrorNoRecord()
	}
	return nil
}

func (this *SysToken) CheckTokenExpireTime(expireMinute int) error {
	user := new(SysToken)
	user.Token = this.Token
	has, err := db.Where("token = ? and status = 'aa'", this.Token).Get(user)
	if err != nil {
		return errCode.NewErrorDB(err)
	}
	if !has {
		return errCode.NewErrorToken()
	}
	if (user.CreateTime.Add(time.Minute * time.Duration(expireMinute))).Unix() < time.Now().Unix() {
		return errCode.NewErrorToken()
	}

	return nil
}
