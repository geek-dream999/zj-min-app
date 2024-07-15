package core

import (
	"fmt"
	"go.uber.org/zap"
	"meet_directly/global"
	"meet_directly/initialize"
	"meet_directly/service/system"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	/**
	* @Author: lyh
	如果使用mongo打开注释即可
	*/
	//if global.GVA_CONFIG.System.UseMongo {
	//	err := initialize.Mongo.Initialization()
	//	if err != nil {
	//		zap.L().Error(fmt.Sprintf("%+v", err))
	//	}
	//}
	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
