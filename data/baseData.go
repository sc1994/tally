package data

import (
	"tally/common"
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

// Count 获取数量
func Count(db string, table string, search interface{}) int {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	i, e := c.Find(search).Count()
	if e != nil {
		panic(err)
	}
	return i
}

// Page 分页
func Page(db string, table string, pageIndex int, pageSize int, fields string, search interface{}, result interface{}) int {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	skip := (pageIndex - 1) * pageSize
	find := c.Find(search)
	count, err := find.Count()
	if err != nil {
		panic(err)
	}
	err = find.Sort(fields).Skip(skip).Limit(pageSize).All(result)
	if err != nil {
		panic(err)
	}
	return count
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

// Delete 删除一条数据
func Delete(db string, table string, selector interface{}) bool {
	session, err := mgo.Dial(common.MongoConnect)
	defer session.Close()
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(db).C(table)
	err = c.Remove(selector)
	if err != nil {
		panic(err)
	} else {
		return true
	}
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
