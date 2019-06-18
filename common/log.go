package common

import (
	"github.com/astaxie/beego/logs"
)

func InitLog() {
	logs.SetLevel(logs.LevelDebug)
	logs.SetLogger(logs.AdapterFile, `{"filename": "logs/test.log", "level": 7}`)
}
