package db

import (
	"gorm.io/gorm"
	"msg/cmd/config"
	"sync"
)

var db *gorm.DB
var dblock = &sync.RWMutex{}

func Get() *gorm.DB {
	dblock.RLock()
	defer dblock.RUnlock()
	return db
}

// Load 连接数据库
func Load() {
	dblock.Lock()
	defer dblock.Unlock()
	if config.Global.Database.Type == "sqlite" {
		loadSqlite()
		return
	} else if config.Global.Database.Type == "postgres" {
		loadPostgres()
		return
	}
	return
}

// Create 创建数据库
func Create() (exist bool) {
	dblock.Lock()
	defer dblock.Unlock()
	if config.Global.Database.Type == "sqlite" {
		return createSqlite()
	} else if config.Global.Database.Type == "postgres" {
		return createPostgres()
	}
	return false
}
