package test

import (
	"testing"
)

func Test_Copy_1(t *testing.T) {
	result := httpRequest("database/copy/5b9621bff35d959bf1b9608a", nil)
	t.Error(result)
}
