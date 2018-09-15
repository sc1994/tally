package test

import (
	"tally/library"
	"testing"
)

func Test_Job_1(t *testing.T) {
	url := library.GetJobAddress("test7", "get", "http://suncheng.xyz:8888/findusersbyname/s", "result", "*/7 * * * *")
	var s = library.HTTPRequest(url, nil, "")
	t.Error(url)
	t.Error(s)
}
