package data

import (
	"go-tally/common"
	"time"

	"github.com/go-redis/redis"
	mgo "gopkg.in/mgo.v2"
)

var options = &redis.Options{
	Addr:     "118.24.27.231:6379",
	Password: "sun940622",
	DB:       0,
}

// Find 获取满足条件的全部数据
func Find(db string, table string, search interface{}, result interface{}) {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	err = c.Find(search).All(result)
	if err != nil {
		panic(err)
	}
	return
}

func Page(db string, table string, pageIndex int, pageSize int, fields string, search interface{}, result interface{}) bool {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		return false
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	skip := (pageIndex - 1) * pageSize
	err = c.Find(search).Sort(fields).Skip(skip).Limit(pageSize).All(result)
	if err != nil {
		return false
	}
	return true
}

// Insert 插入一条数据
func Insert(db string, table string, data interface{}) bool {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		return false
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	err = c.Insert(data)
	if err != nil {
		return false
	}
	return true
}

// Update 更新一条数据
func Update(db string, table string, selector interface{}, update interface{}) bool {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	err = c.Update(selector, update)
	if err != nil {
		panic(err)
	}
	return true
}

// RedisGet 读取缓存
func RedisGet(key string) (string, bool) {
	client := redis.NewClient(options)
	defer client.Close()
	r, e := client.Get(key).Result()
	if e != nil {
		return "", false
	}
	return r, true
}

// RedisSet 设置缓存(过期最小粒度:分钟)
func RedisSet(key string, value interface{}, minute int32) bool {
	client := redis.NewClient(options)
	defer client.Close()
	e := client.Set(key, value, time.Duration(minute*60*1000*1000*1000)).Err()
	if e != nil {
		return false
	}
	return true
}

// RedisDel 删除一个key
func RedisDel(key string) bool {
	client := redis.NewClient(options)
	defer client.Close()
	info := client.Del(key)
	if info.Err() != nil {
		return false
	}
	return true
}
