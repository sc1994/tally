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

func Test_Backup_1(t *testing.T) {
	url := library.GetJobAddress("tally_backup", "get", "http://118.24.27.231/:8888/database/copy/5b9621bff35d959bf1b9608a", "success", "0 0 * * *")
	var s = library.HTTPRequest(url, nil, "")
	if s != "ok" {
		t.Error(s)
	}
}
