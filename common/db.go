package common

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB orm.Ormer
)

func InitOrm(dbAlias string, dbDSN string, autoCreateTable bool) {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbDSN)

	// 自动建表
	if autoCreateTable {
		orm.RunSyncdb(dbAlias, false, true)
	}
}

func RunOrm(dbAlias string) {
	DB = orm.NewOrm()
	DB.Using(dbAlias)
}
