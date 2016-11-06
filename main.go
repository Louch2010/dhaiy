package main

import (
	"os"

	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/cmd"
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/gdb"
	"github.com/louch2010/dhaiy/log"
	"github.com/louch2010/dhaiy/server"
)

func main() {
	//初始化日志
	args := os.Args
	log.Info("启动参数：", args)
	//初始化配置文件
	configPath := ""
	if len(args) > 1 {
		configPath = args[1]
	}
	err := InitConfig(configPath)
	if err != nil {
		log.Error("初始化配置文件失败！", err)
		return
	}
	//实例化配置
	LoadServerConfig()
	config := GetServerConfig()
	//初始化日志
	err = log.InitLog(config.LogLevel, config.LogFormat, config.LogPath, config.LogRoll, config.LogConsoleOn)
	if err != nil {
		log.Error("初始化日志失败！", err)
		return
	}
	//初始化缓存
	log.Info("初始化缓存...")
	err = cache.InitCache()
	if err != nil {
		log.Error("初始化缓存失败！", err)
		return
	}
	//初始化命令
	cmd.InitCmd()
	//载入gdb文件
	log.Info("初始化、加载持久化文件...")
	err = gdb.InitGDB()
	if err != nil {
		log.Error("初始化、加载持久化文件失败！", err)
	}
	//启动服务
	server.StartServer(config)
	//服务停止 - 持久化
	log.Info("服务停止前进行执行久...")
	gdb.SaveDB()
	log.Info("服务停止完成!")
}
