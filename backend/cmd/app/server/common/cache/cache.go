package cache

import (
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	ExpirationRemember    = time.Hour * time.Duration(24*14) // 记住用户,保留14天
	ExpirationNotRemember = time.Hour * time.Duration(1)     // 不记住用户,保留1小时
	ExpirationAccessToken = time.Hour * time.Duration(24*1)  // 记住访问密钥,保留1天
	ExpirationLoginLock   = time.Minute * time.Duration(5)   // 太多失败登录,锁定用户5分钟
)

var TokenManager *cache.Cache       // 用户访问密钥
var LoginFailedManager *cache.Cache // 登录失败后保存登录者特征(ip+用户名),当次数过多时暂时锁定用户

func init() {
	TokenManager = cache.New(5*time.Minute, 10*time.Minute)
	LoginFailedManager = cache.New(5*time.Minute, 10*time.Minute)
}

func SetTokenManagerOnEvicted(f func(string, any)) {
	TokenManager.OnEvicted(f)
}

// GetFailedCountByIpUsername 通过错误尝试次数
func GetFailedCountByIpUsername(ip, username string) int {
	key := ip + "_" + username
	get, ok := LoginFailedManager.Get(key)
	if ok {
		return get.(int)
	}
	return 0
}

// SetFailedCountByIpUsername 通过错误尝试次数
func SetFailedCountByIpUsername(ip, username string, count int) {
	key := ip + "_" + username
	LoginFailedManager.Set(key, count, ExpirationLoginLock)
}

func DeleteTokens(tokens []string) {
	for _, token := range tokens {
		TokenManager.Delete(token)
	}
}
