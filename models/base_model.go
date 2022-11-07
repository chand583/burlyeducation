package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/lib/pq"
)

func init() {
	dbDriver, _ := config.String("db::driver")
	dbConString, _ := config.String("db::con_string")
	switch dbDriver {
	case "postgres":
		err := orm.RegisterDriver(dbDriver, orm.DRPostgres)
		fmt.Println("registre drive", err)
	case "mysql":
		orm.RegisterDriver(dbDriver, orm.DRMySQL)

	}
	err := orm.RegisterDataBase("default", dbDriver, dbConString)
	if err != nil {
		fmt.Println(err)
	}
	orm.SetMaxIdleConns("default", 10)
	orm.Debug = true

}
