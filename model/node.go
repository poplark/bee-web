package model

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(Node))
}

type Node struct {
	Id        uint32    `orm:"pk;auto"`
	Name      string    `orm:"default(node)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}
