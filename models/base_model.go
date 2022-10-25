package models

import (
	"burlyeducation/lib"
	"burlyeducation/log"
	_ "fmt"

	"github.com/beego/beego/v2/client/cache"
	_ "github.com/beego/beego/v2/client/cache/redis"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/lib/pq"
)

var bm cache.Cache
var isCacheEnable bool
var err error

func init() {

	secretManager, _ := config.String("SECRET_MANAGER")

	cacheEngine, _ := config.String("cache::engine")
	cacheConStr, _ := config.String("cache::con_string")
	isCacheEnable, _ = config.Bool("cache::enable_cache")

	dbDriver, _ := config.String("db::driver")
	dbConString, _ := config.String("db::con_string")

	if secretManager == "aws" {
		dbConString, err = lib.AWSSecretManager{}.GetDBConString()

		if err != nil {
			log.Error(1106, map[string]interface{}{"error_details": err})
		}

		cacheConStr, err = lib.AWSSecretManager{}.GetRedisConString()

		if err != nil {
			log.Error(1105, map[string]interface{}{"error_details": err})
		}
	}

	//init cache if enabled in config
	if isCacheEnable {
		bm, err = cache.NewCache(cacheEngine, cacheConStr)
		if err != nil {
			log.Warning(1051, map[string]interface{}{"error_details": err})
		}
	}

	//init db
	switch dbDriver {
	case "postgres":
		orm.RegisterDriver(dbDriver, orm.DRPostgres)
	case "mysql":
		orm.RegisterDriver(dbDriver, orm.DRMySQL)
	}

	orm.RegisterDataBase("default", dbDriver, dbConString)
	orm.SetMaxIdleConns("default", 10)
	//orm.SetMaxOpenConns("default", 30)

	orm.RegisterModel(new(Article))

	if currentMode, _ := config.String("runmode"); currentMode == "dev" {
		orm.Debug = true
	}
}
