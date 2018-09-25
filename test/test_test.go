package test

import (
	"tally/models"
	"testing"
)

func Test_Test_1(t *testing.T) {
	models.AddTestStruct()
	r := models.GetTestStruct()
	t.Error(r)
}
