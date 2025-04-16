package db

import (
	"fmt"
	"github.com/Akvicor/glog"
	"github.com/Akvicor/util"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"msg/cmd/config"
)

func loadSqlite() {
	if util.FileStat(config.Global.Database.File).NotFile() {
		glog.Fatal("database file %s not exit", config.Global.Database.File)
	}
	var logMode logger.Interface
	if utils.Contains(config.Global.Log.Debug, "database") {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	var err error
	var d *gorm.DB
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.Global.Database.File)
	d, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger:                                   logMode,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		glog.Fatal("连接数据库异常: %v", err)
	}

	db = d
}

func createSqlite() (exist bool) {
	exist = false
	if util.FileStat(config.Global.Database.File).IsExist() {
		exist = true
	}

	var logMode logger.Interface
	if utils.Contains(config.Global.Log.Debug, "database") {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	var err error
	var d *gorm.DB
	dsn := fmt.Sprintf("file:%s?cache=shared&mode=rwc", config.Global.Database.File)
	d, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger:                                   logMode,
		SkipDefaultTransaction:                   true,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		glog.Fatal("连接数据库异常: %v", err)
	}

	db = d
	return exist
}
