package common

// "github.com/gorilla/websocket"

// MongoConnect 连接字符串
const MongoConnect string = "118.24.27.231:27017/"

// TallyMode  "收入", "支出", "预支"
var TallyMode = [3]string{"收入", "支出", "预支"}

// AdminID 管理员账号
var AdminID = "5b83f2461cdea200010ea5f1"

// websocker 跨域配置
// var Upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }
