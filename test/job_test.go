package test

import (
	"tally/library"
	"testing"
)

func Test_Job_1(t *testing.T) {
	url := library.GetJobAddress("test1", "get", "http://suncheng.xyz:8888/findusersbyname/s", "result", "*/1 * * * *")
	var s = library.HTTPRequest(url, nil, "")
	t.Error(url)
	t.Error(s)
}
