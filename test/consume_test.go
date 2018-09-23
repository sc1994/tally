package test

import (
	"testing"

	"tally/models"

	"gopkg.in/mgo.v2/bson"
)

func Test_GetPipe_1(t *testing.T) {
	consumes := new(models.TallyRequest).Pipe(
		bson.M{
			"$match": bson.M{
			// "money": bson.M{"$gt": 1},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   "$type",
				"count": bson.M{"$sum": 1},
				"utime": bson.M{"$max": "$utime"},
				"ctime": bson.M{"$max": "$ctime"},
			},
		})
	// keys := make([]string, 0, len(consumes))
	// result := make([]interface{}, 0, 1)
	// // keys = append(keys, k)
	// for k := range consumes[0] {
	// 	result = append(result, consumes[0][k])
	// }
	t.Error(consumes)
}
