package main

import (
	"os"

	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/cmd"
	. "github.com/louch2010/dhaiy/common"
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
	//初始化日志
	level := GetSystemConfig().MustValue("log", "level", log.LOG_DEFAULT_LEVEL)
	format := GetSystemConfig().MustValue("log", "format", log.LOG_DEFAULT_FORMAT)
	path := GetSystemConfig().MustValue("log", "path", log.LOG_DEFAULT_PATH)
	roll := GetSystemConfig().MustValue("log", "roll", log.LOG_DEFAULT_ROLL)
	consoleOn := GetSystemConfig().MustBool("log", "console.on", true)
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
	//初始化命令
	cmd.InitCmd()
	//启动服务
	port := GetSystemConfig().MustInt("server", "port", 1334)
	aliveTime := GetSystemConfig().MustInt("server", "aliveTime", 30)
	connectType := GetSystemConfig().MustValue("server", "connectType", "long")
	server.StartServer(port, aliveTime, connectType)
}
