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
	cmdHandlers["ping"] = newCmd("ping", 0, 0, false, HandlePingCommand)
	cmdHandlers["help"] = newCmd("help", 0, 1, false, HandleHelpCommand)
	cmdHandlers["connect"] = newCmd("connect", 0, 6, false, HandleConnectCommand)
	cmdHandlers["exit"] = newCmd("exit", 0, 0, false, HandleExitCommand)
	cmdHandlers["del"] = newCmd("del", 1, 100, true, HandleDelCommand)
	cmdHandlers["exist"] = newCmd("exist", 1, 1, false, HandleExistCommand)
	cmdHandlers["select"] = newCmd("select", 1, 1, false, HandleSelectCommand)
	cmdHandlers["showt"] = newCmd("showt", 0, 1, false, HandleShowtCommand)
	cmdHandlers["showi"] = newCmd("showi", 0, 1, false, HandleShowiCommand)
	cmdHandlers["info"] = newCmd("info", 0, 0, false, HandleInfoCommand)
	cmdHandlers["bgsave"] = newCmd("bgsave", 0, 0, false, HandleBgSaveCommand)
	cmdHandlers["flushdb"] = newCmd("flushdb", 0, 0, false, HandleFlushDBCommand)
	cmdHandlers["flushall"] = newCmd("flushall", 0, 0, false, HandleFlushAllCommand)
	//string
	cmdHandlers["set"] = newCmd("set", 2, 3, true, HandleSetCommand)
	cmdHandlers["get"] = newCmd("get", 1, 1, false, HandleGetCommand)
	cmdHandlers["append"] = newCmd("append", 2, 2, true, HandleAppendCommand)
	cmdHandlers["strlen"] = newCmd("strlen", 1, 1, false, HandleStrLenCommand)
	cmdHandlers["setnx"] = newCmd("setnx", 2, 2, true, HandleSetNxCommand)
	//num
	cmdHandlers["nset"] = newCmd("nset", 2, 3, true, HandleNSetCommand)
	cmdHandlers["nget"] = newCmd("nget", 1, 1, false, HandleNGetCommand)
	cmdHandlers["incr"] = newCmd("incr", 1, 1, true, HandleIncrCommand)
	cmdHandlers["incrby"] = newCmd("incrby", 2, 2, false, HandleIncrByCommand)
	//map
	cmdHandlers["hset"] = newCmd("hset", 3, 3, true, HandleHSetCommand)
	cmdHandlers["hget"] = newCmd("hget", 2, 2, false, HandleHGetCommand)
	cmdHandlers["hdel"] = newCmd("hdel", 2, 2, true, HandleHDelCommand)
	cmdHandlers["hexists"] = newCmd("hexists", 2, 2, false, HandleHExistsCommand)
	cmdHandlers["hlen"] = newCmd("hlen", 1, 1, false, HandleHLenCommand)
	cmdHandlers["hkeys"] = newCmd("hkeys", 1, 1, false, HandleHKeysCommand)
	cmdHandlers["hvals"] = newCmd("hvals", 1, 1, false, HandleHValsCommand)
	cmdHandlers["hsetnx"] = newCmd("hsetnx", 3, 3, true, HandleHSetNxCommand)
	//set
	cmdHandlers["sadd"] = newCmd("sadd", 2, 2, true, HandleSAddCommand)
	cmdHandlers["scard"] = newCmd("scard", 1, 1, false, HandleSCardCommand)
	cmdHandlers["smembers"] = newCmd("smembers", 1, 1, false, HandleSMembersCommand)
	cmdHandlers["srem"] = newCmd("srem", 2, 2, true, HandleSRemCommand)
	cmdHandlers["sismember"] = newCmd("sismember", 2, 2, false, HandleSisMemberCommand)
	cmdHandlers["spop"] = newCmd("spop", 1, 1, false, HandleSPopCommand)
	cmdHandlers["srandmember"] = newCmd("srandmember", 1, 1, false, HandleSRandMemberCommand)
	//list
	cmdHandlers["lpush"] = newCmd("lpush", 2, 2, true, HandleLPushCommand)
	cmdHandlers["lpushx"] = newCmd("lpushx", 2, 2, true, HandleLPushXCommand)
	cmdHandlers["lrange"] = newCmd("lrange", 3, 3, false, HandleLRangeCommand)
	cmdHandlers["llen"] = newCmd("llen", 1, 1, false, HandleLLenCommand)
	cmdHandlers["lrem"] = newCmd("lrem", 3, 3, true, HandleLRemCommand)
	cmdHandlers["lindex"] = newCmd("lindex", 2, 2, false, HandleLIndexCommand)
	cmdHandlers["lpop"] = newCmd("lpop", 1, 1, true, HandleLPopCommand)
	cmdHandlers["rpush"] = newCmd("rpush", 2, 2, false, HandleRPushCommand)
	cmdHandlers["rpushx"] = newCmd("rpushx", 2, 2, false, HandleRPushXCommand)
	cmdHandlers["rpop"] = newCmd("rpop", 2, 2, false, HandleRPopXCommand)
	//zset
	cmdHandlers["zadd"] = newCmd("zadd", 3, 3, true, HandleZAddCommand)
	cmdHandlers["zcard"] = newCmd("zcard", 1, 1, false, HandleZCardCommand)
	cmdHandlers["zcount"] = newCmd("zcount", 3, 3, false, HandleZCountCommand)
	cmdHandlers["zscore"] = newCmd("zscore", 2, 2, false, HandleZScoreCommand)
	cmdHandlers["zrank"] = newCmd("zrank", 2, 2, false, HandleZRankCommand)
	cmdHandlers["zrem"] = newCmd("zrem", 2, 2, true, HandleZRemkCommand)
	cmdHandlers["zincrby"] = newCmd("zincrby", 3, 3, true, HandleZIncrByCommand)
	//系统命令
	cmdHandlers["slaveof"] = newCmd("slaveof", 1, 1, false, HandleSlaveOfCommand)
	cmdHandlers["syncdata"] = newCmd("syncdata", 4, 4, false, HandleSyncDataCommand)

	log.Info("初始化命令集合完成")
}

func newCmd(name string, min int, max int, write bool, f func(client *Client)) *Cmd {
	cmd := Cmd{
		Name:        name,
		MinParam:    min,
		MaxParam:    max,
		IsWrite:     write,
		InvokeCount: 0,
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
