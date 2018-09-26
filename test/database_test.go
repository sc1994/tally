package test

import (
	"testing"
)

func Test_Copy_1(t *testing.T) {
	result := httpRequest("database/copy/5b9621bff35d959bf1b9608a", nil)
	t.Error(result)
}

// // 清洗channel 数据
// func Test_WashChannel_1(t *testing.T) {
// 	var all []models.Channel
// 	data.Find("channel",
// 		bson.M{
// 			"content": bson.M{
// 				"$in": []string{"信用卡", "白条", "花呗"},
// 			},
// 		},
// 		&all)
// 	var ids []bson.ObjectId
// 	linq.From(all).Select(func(x interface{}) interface{} {
// 		return x.(models.Channel).ID
// 	}).ToSlice(&ids)
// 	data.Update("channel",
// 		bson.M{
// 			"$set": bson.M{
// 				"isAdvance": true,
// 			},
// 		},
// 		bson.M{
// 			"_id": bson.M{"$in": ids},
// 		})
// }
