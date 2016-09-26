package main

import (
	"os"

	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/common"
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
	err := common.InitConfig(configPath)
	if err != nil {
		log.Error("初始化配置文件失败！", err)
		return
	}
	//初始化日志
	level := common.GetSystemConfig().MustValue("log", "level", log.LOG_DEFAULT_LEVEL)
	format := common.GetSystemConfig().MustValue("log", "format", log.LOG_DEFAULT_FORMAT)
	path := common.GetSystemConfig().MustValue("log", "path", log.LOG_DEFAULT_PATH)
	roll := common.GetSystemConfig().MustValue("log", "roll", log.LOG_DEFAULT_ROLL)
	consoleOn := common.GetSystemConfig().MustBool("log", "console.on", true)
	err = log.InitLog(level, format, path, roll, consoleOn)
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
	//启动服务
	port := common.GetSystemConfig().MustInt("server", "port", 1334)
	aliveTime := common.GetSystemConfig().MustInt("server", "aliveTime", 30)
	connectType := common.GetSystemConfig().MustValue("server", "connectType", "long")
	server.StartServer(port, aliveTime, connectType)
}