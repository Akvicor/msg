package db

import (
	"database/sql"
	"fmt"
	"github.com/Akvicor/glog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"msg/cmd/config"

	_ "github.com/lib/pq"
)

func loadPostgres() {
	var err error
	var d *gorm.DB
	var dsn string

	host := config.Global.Database.Host
	port := config.Global.Database.Port
	dbname := config.Global.Database.Database
	username := config.Global.Database.Username
	password := config.Global.Database.Password

	var logMode logger.Interface
	if utils.Contains(config.Global.Log.Debug, "database") {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, username, password, dbname)
	d, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logMode,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		glog.Fatal("连接数据库异常: %v", err)
	}

	db = d
}

func createPostgres() (exist bool) {
	exist = false
	host := config.Global.Database.Host
	port := config.Global.Database.Port
	dbname := config.Global.Database.Database
	username := config.Global.Database.Username
	password := config.Global.Database.Password

	var err error
	var td *sql.DB
	var d *gorm.DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, username, password)

	td, err = sql.Open("postgres", dsn)
	if err != nil {
		glog.Fatal("连接数据库异常: %v", err)
	}
	defer td.Close()

	err = td.QueryRow("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)", dbname).Scan(&exist)
	if err != nil {
		glog.Fatal("查询数据库异常: %v", err)
	}

	if exist {
		glog.Info("数据库 '%s' 已存在", dbname)
	} else {
		_, err = td.Exec(fmt.Sprintf("CREATE DATABASE %s", dbname))
		if err != nil {
			glog.Fatal("创建数据库异常: %v", err)
		}
		glog.Info("数据库 '%s' 已创建", dbname)
	}

	var logMode logger.Interface
	if utils.Contains(config.Global.Log.Debug, "database") {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	dsn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai", host, port, username, password, dbname)
	d, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:                 logMode,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		glog.Fatal("连接数据库异常: %v", err)
	}

	db = d
	return exist
}
