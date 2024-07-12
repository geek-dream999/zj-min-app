package main

import (
	"go.uber.org/zap"
	"meet_directly/core"
	"meet_directly/global"
	"meet_directly/initialize"
)

func main() {
	global.GVA_VP = core.Viper() //初始化Viper
	global.GVA_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	initialize.DBList()               //用于从库配置
	core.RunWindowsServer()           //注册路由&中间件
}
