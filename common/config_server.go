package common

import (
	"sync"

	"github.com/louch2010/dhaiy/log"
)

//服务端配置
type ServerConfig struct {
	//状态相关
	ServerId   string        //服务唯一标识
	SlaveList  []*ServerInfo //从库列表
	Master     *ServerInfo   //主库
	DumpStatus int           //dump文件状态，1：未执行、2：执行中
	//配置相关
	AppName            string //应用名
	APPVersion         string //版本号
	AppAuthor          string //应用开发作者
	ServerPort         int    //服务端口号
	ServerPassword     string //服务密码
	ServerMaxPoolSize  int    //最大线程数
	ServerCorePoolSize int    //核心线程数
	ServerConnectType  string //连接类型
	ServerAliveTime    int    //连接存活时间
	ServerSysTable     string //系统表表名
	ServerAnonymCommnd string //匿名命令
	DefaultTable       string //默认表名
	ClientOpenSession  bool   //客户端是否启用会话
	DumpOn             bool   //是否开启DUMP功能
	DumpFilePath       string //dump文件路径
	DumpTrigger        string //dump触发器
	LogLevel           string //日志级别
	LogConsoleOn       bool   //是否开户控制台日志
	LogFormat          string //日志格式
	LogPath            string //日志路径
	LogRoll            string //日志文件切换策略
}

var (
	config *ServerConfig
	lock   sync.RWMutex
)

func LoadServerConfig() {
	config = initServerConfig()
}

//初始化服务端配置
func initServerConfig() *ServerConfig {
	lock.Lock()
	defer lock.Unlock()
	config := &ServerConfig{
		ServerId:           GetGuid(),
		ServerPort:         GetSystemConfig().MustInt("server", "port", 1334),                              //服务端口号
		ServerPassword:     GetSystemConfig().MustValue("server", "password", ""),                          //服务密码
		ServerMaxPoolSize:  10,                                                                             //最大线程数
		ServerCorePoolSize: 5,                                                                              //核心线程数
		ServerConnectType:  GetSystemConfig().MustValue("server", "connectType", "long"),                   //连接类型
		ServerAliveTime:    GetSystemConfig().MustInt("server", "aliveTime", 30),                           //连接存活时间                                                            //服务唯一标识
		ServerSysTable:     GetSystemConfig().MustValue("server", "sysTable", "sys"),                       //系统表表名
		ServerAnonymCommnd: SystemConfigFile.MustValue("server", "anonymCommnd", "ping,connect,exit,help"), //匿名命令
		DefaultTable:       GetSystemConfig().MustValue("table", "default", "default"),                     //默认表名
		ClientOpenSession:  GetSystemConfig().MustBool("client", "openSession", true),                      //客户端是否启用会话
		DumpOn:             GetSystemConfig().MustBool("dump", "dump.on", true),                            //是否开启DUMP功能
		DumpFilePath:       GetSystemConfig().MustValue("dump", "filePath", "./data/dump.gdb"),             //dump文件路径
		DumpTrigger:        GetSystemConfig().MustValue("dump", "trigger", ""),                             //dump触发器
		LogLevel:           GetSystemConfig().MustValue("log", "level", log.LOG_DEFAULT_LEVEL),             //日志级别
		LogConsoleOn:       GetSystemConfig().MustBool("log", "console.on", true),                          //是否开户控制台日志
		LogFormat:          GetSystemConfig().MustValue("log", "format", log.LOG_DEFAULT_FORMAT),           //日志格式
		LogPath:            GetSystemConfig().MustValue("log", "path", log.LOG_DEFAULT_PATH),               //日志路径
		LogRoll:            GetSystemConfig().MustValue("log", "roll", log.LOG_DEFAULT_ROLL),               //日志文件切换策略
	}
	return config
}

//获取服务端配置
func GetServerConfig() *ServerConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}
