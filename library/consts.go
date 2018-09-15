package library

// TallyMode  "收入", "支出", "预支"
var TallyMode = [3]string{"收入", "支出", "预支"}

// AdminID 管理员账号
var AdminID = "5b8aa88e3a7ae00001ff783c"

// JobAddress job地址
var JobAddress = "http://118.24.27.231:8989/addjob?name=%s&method=%s&url=%s&assert=%s&cron=%s"

/*
  s := "123123%s123123"
	u := "https://www.baidu.com/s?ie=UTF-8&wd=123"
	c := url.QueryEscape(u)
	fmt.Printf(c)
	b := fmt.Sprintf(s,c)
*/
