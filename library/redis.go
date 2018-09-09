package library

import (
	"time"

	"github.com/go-redis/redis"
)

var options = &redis.Options{
	Addr:     "118.24.27.231:6379",
	Password: "sun940622",
	DB:       0,
}

// GetRedis 读取缓存
func GetRedis(key string) (string, bool) {
	client := redis.NewClient(options)
	defer client.Close()
	r, e := client.Get(key).Result()
	if e != nil {
		return "", false
	}
	return r, true
}

// GetRedisKeys 获取Keys
func GetRedisKeys(pattern string) []string {
	client := redis.NewClient(options)
	defer client.Close()
	return client.Keys(pattern).Val()
}

// SetRedis 设置缓存(过期最小粒度:分钟)
func SetRedis(key string, value interface{}, minute int32) bool {
	client := redis.NewClient(options)
	defer client.Close()
	time := time.Duration(minute * 60 * 1000 * 1000 * 1000)
	e := client.Set(key, value, time).Err()
	if e != nil {
		return false
	}
	return true
}

// DelRedis 删除一个key
func DelRedis(key string) bool {
	client := redis.NewClient(options)
	defer client.Close()
	info := client.Del(key)
	if info.Err() != nil {
		return false
	}
	return true
}
