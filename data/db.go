package data

import (
	"strconv"
	"tally/library"
	"time"

	"gopkg.in/mgo.v2/bson"

	mgo "gopkg.in/mgo.v2"
)

// DBName 库名称
var dbName = "tally"

// MongoConnect 连接字符串
const mongoConnectString string = "118.24.27.231:27017/"

func connectDB(table string) (*mgo.Session, *mgo.Collection) {
	s, err := mgo.Dial(mongoConnectString)
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	c := s.DB(dbName).C(table)
	return s, c
}

// Find 查找
func Find(table string, search interface{}, result interface{}) {
	s, c := connectDB(table)
	defer s.Clone()
	c.Find(search).All(result)
	return
}

// Pipe Pipe
func Pipe(table string, search interface{}, result interface{}) {
	s, c := connectDB(table)
	defer s.Clone()
	c.Pipe(search).All(result)
	return
}

// Page 分页
func Page(table string, search interface{}, index int, size int, result interface{}, sort ...string) int {
	s, c := connectDB(table)
	defer s.Clone()
	skip := (index - 1) * size
	find := c.Find(search)
	count, err := find.Count()
	if err != nil {
		panic(err)
	}
	if library.IsEmpty(sort) {
		sort = append(sort, "_id")
	}
	err = find.Sort(sort...).Skip(skip).Limit(size).All(result)
	if err != nil {
		panic(err)
	}
	return count
}

// GetCount 获取数量
func GetCount(table string, search interface{}) int {
	s, c := connectDB(table)
	defer s.Clone()
	n, e := c.Find(search).Count()
	if e != nil {
		panic(e)
	}
	return n
}

// Exist 是否存在
func Exist(table string, search interface{}) bool {
	return GetCount(table, search) > 0
}

// Update 更新
func Update(table string, update interface{}, selector interface{}) (change *mgo.ChangeInfo) {
	s, c := connectDB(table)
	defer s.Clone()
	change, e := c.UpdateAll(selector, update)
	if e != nil {
		panic(e)
	}
	return
}

// Remove 删除
func Remove(table string, selector interface{}) {
	s, c := connectDB(table)
	defer s.Clone()
	e := c.Remove(selector)
	if e != nil {
		panic(e)
	}
}

// InsertBatch 插入文档支持批量
func InsertBatch(table string, docs []interface{}) {
	s, c := connectDB(table)
	defer s.Clone()
	e := c.Insert(docs...)
	if e != nil {
		panic(e)
	}
}

// Insert 插入
func Insert(table string, doc interface{}) {
	s, c := connectDB(table)
	defer s.Clone()
	e := c.Insert(doc)
	if e != nil {
		panic(e)
	}
}

// CopyTable 数据copy
func CopyTable(table string) {
	// 旧文档
	var docs []interface{}
	Find(table, bson.M{}, &docs)
	// 新文档
	s, err := mgo.Dial(mongoConnectString)
	defer s.Clone()
	if err != nil {
		panic(err)
	}
	s.SetMode(mgo.Monotonic, true)
	// 备份
	newDB := strconv.Itoa(time.Now().YearDay() % 3)
	c := s.DB(dbName + "Backup" + newDB).C(table)
	c.RemoveAll(bson.M{})
	c.Insert(docs...)
}
