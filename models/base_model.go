package models

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
)

func init() {

	dbDriver, _ := config.String("db::driver")
	dbConString, _ := config.String("db::con_string")
	switch dbDriver {
	case "postgres":
		orm.RegisterDriver(dbDriver, orm.DRPostgres)
	case "mysql":
		orm.RegisterDriver(dbDriver, orm.DRMySQL)

	}
	orm.RegisterDataBase("default", dbDriver, dbConString)
	orm.SetMaxIdleConns("default", 10)
	//orm.RegisterModel(new())
	orm.Debug = true

}
