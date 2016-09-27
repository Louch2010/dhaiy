package cmd

import (
	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//命令集合
var cmdHandlers = make(map[string]*common.Cmd)

//初始化命令集合
func InitCmd() {
	log.Info("初始化命令集合...")
	cmdHandlers["ping"] = newCmd("ping", 0, 0, false, HandlePingCommnd)
	cmdHandlers["help"] = newCmd("help", 0, 1, false, HandleHelpCommnd)
	cmdHandlers["connect"] = newCmd("connect", 0, 0, false, HandleConnectCommnd)
	//	cmdHandlers["delete"] = newCmd("delete", 1, 100, true, HandleDeleteCommnd)
	//	cmdHandlers["exist"] = newCmd("exist", 1, 1, false, HandleExistCommnd)
	//	cmdHandlers["use"] = newCmd("use", 1, 1, false, HandleUseCommnd)
	//	cmdHandlers["showt"] = newCmd("showt", 0, 1, false, HandleShowtCommnd)
	//	cmdHandlers["showi"] = newCmd("showi", 0, 1, false, HandleShowiCommnd)
	//	cmdHandlers["info"] = newCmd("info", 0, 0, false, HandleInfoCommnd)
	//	cmdHandlers["bgsave"] = newCmd("bgsave", 0, 0, false, HandleBgSaveCommnd)
}

func newCmd(name string, min int, max int, write bool, f func(client *common.Client)) *common.Cmd {
	cmd := common.Cmd{
		Name:        name,
		MinParam:    min,
		MaxParam:    max,
		IsWrite:     write,
		HandlerFunc: f,
	}
	return &cmd
}

func GetCmd(name string) *common.Cmd {
	cmd, ok := cmdHandlers[name]
	if !ok {
		return nil
	}
	return cmd
}

//创建会话
func CreateSession(token string, c *common.Client) bool {
	//缓存登录信息
	table, _ := cache.GetSysTable()
	table.Set(token, c, 0, common.DATA_TYPE_OBJECT)
	//创建表信息
	cache.Cache(c.Table)
	return true
}

//销毁会话
func DestroySession(token string) bool {
	table, _ := cache.GetSysTable()
	return table.Delete(token)
}
