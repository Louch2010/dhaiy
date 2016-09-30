package cmd

import (
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//命令集合
var cmdHandlers = make(map[string]*Cmd)

//初始化命令集合
func InitCmd() {
	log.Info("初始化命令集合...")
	//通用
	cmdHandlers["ping"] = newCmd("ping", 0, 0, false, HandlePingCommnd)
	cmdHandlers["help"] = newCmd("help", 0, 1, false, HandleHelpCommnd)
	cmdHandlers["connect"] = newCmd("connect", 0, 0, false, HandleConnectCommnd)
	cmdHandlers["delete"] = newCmd("delete", 1, 100, true, HandleDeleteCommnd)
	cmdHandlers["exist"] = newCmd("exist", 1, 1, false, HandleExistCommnd)
	cmdHandlers["use"] = newCmd("use", 1, 1, false, HandleUseCommnd)
	cmdHandlers["showt"] = newCmd("showt", 0, 1, false, HandleShowtCommnd)
	cmdHandlers["showi"] = newCmd("showi", 0, 1, false, HandleShowiCommnd)
	cmdHandlers["info"] = newCmd("info", 0, 0, false, HandleInfoCommnd)
	cmdHandlers["bgsave"] = newCmd("bgsave", 0, 0, false, HandleBgSaveCommnd)
	//string
	cmdHandlers["set"] = newCmd("set", 2, 3, true, HandleSetCommnd)
	cmdHandlers["get"] = newCmd("get", 1, 1, false, HandleGetCommnd)
	cmdHandlers["append"] = newCmd("append", 2, 2, true, HandleAppendCommnd)
	cmdHandlers["strlen"] = newCmd("strlen", 1, 1, false, HandleStrLenCommnd)
	cmdHandlers["setnx"] = newCmd("setnx", 2, 2, true, HandleSetNxCommnd)
	//num
	cmdHandlers["nset"] = newCmd("nset", 2, 3, true, HandleNSetCommnd)
	cmdHandlers["nget"] = newCmd("nget", 1, 1, false, HandleNGetCommnd)
	cmdHandlers["incr"] = newCmd("incr", 1, 1, true, HandleIncrCommnd)
	cmdHandlers["incrby"] = newCmd("incrby", 2, 2, false, HandleIncrByCommnd)

	log.Info("初始化命令集合完成")
}

func newCmd(name string, min int, max int, write bool, f func(client *Client)) *Cmd {
	cmd := Cmd{
		Name:        name,
		MinParam:    min,
		MaxParam:    max,
		IsWrite:     write,
		HandlerFunc: f,
	}
	return &cmd
}

func GetCmd(name string) *Cmd {
	cmd, ok := cmdHandlers[name]
	if !ok {
		return nil
	}
	return cmd
}
