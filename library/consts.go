package library

import (
	"fmt"
	"net/url"
)

// TallyMode  "收入", "支出", "预支"
var TallyMode = [3]string{"收入", "支出", "预支"}

// AdminID 管理员账号
var AdminID = "5b8aa88e3a7ae00001ff783c"

// JobAddress job地址
var jobAddress = "http://118.24.27.231:8989/addjob?name=%s&method=%s&url=%s&assert=%s&cron=%s"

// GetJobAddress job地址
func GetJobAddress(name string, method string, urlSring string, assert string, cron string) string {
	u := url.QueryEscape(urlSring)
	return fmt.Sprintf(jobAddress, name, method, u, assert, url.QueryEscape(cron))
}
