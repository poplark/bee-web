package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	UserId    uint32    `orm:"pk;auto;default(1)" description:"用户ID"`
	UserName  string    `orm:"default(username)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
