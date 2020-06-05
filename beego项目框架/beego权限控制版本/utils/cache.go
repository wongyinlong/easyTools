package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"os"
	"time"
)

var (
	bm cache.Cache
)

const (
	CacheAdapterRedis  = "redis"
	CacheAdapterMemory = "memory"
	CacheAdapterFile   = "file"
)

func init() {
	rdsAdapterName := beego.AppConfig.DefaultString("rds.cache.adapter", "memory")

	switch rdsAdapterName {
	case CacheAdapterRedis:
		rdsCacheStr := beego.AppConfig.DefaultString("rds.connect.strings", "")
		var err error
		bm, err = cache.NewCache(rdsAdapterName, rdsCacheStr)
		if err != nil {
			logs.Error("缓存初始化失败!", err.Error())
			os.Exit(-1)
		}
	default:
		rdsCacheStr := beego.AppConfig.DefaultString("rds.memory.strings", "")
		bm, _ = cache.NewCache(rdsAdapterName, rdsCacheStr)
	}
}

func Get(key string) interface{} {
	return bm.Get(key)
}

func Put(key string, val interface{}, timeout ...time.Duration) error {
	if len(timeout) > 0 {
		return bm.Put(key, val, timeout[0])
	} else {
		return bm.Put(key, val, time.Duration(0))
	}
}

func Delete(key string) error {
	return bm.Delete(key)
}

func Incr(key string) error {
	return bm.Incr(key)
}

func Decr(key string) error {
	return bm.Decr(key)
}

func IsExist(key string) bool {
	return bm.IsExist(key)
}

func Cache() cache.Cache {
	return bm
}
