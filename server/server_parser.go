package server

import (
	"strings"

	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/cmd"
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//解析请求
func ParserRequest(client *Client) {
	token := client.Token
	request := client.Reqest
	log.Debug("开始处理请求，token：", token, "，请求内容为：", request)
	//会话信息校验
	if !client.IsLogin {
		GetSession(client) //尝试从cache中获取客户端信息
	}
	//没有登录
	if !client.IsLogin {
		openSession := client.ServerConfig.ClientOpenSession
		//需要登录，而且也不是免登录的命令
		if openSession && !IsAnonymCommnd(request[0], client.ServerConfig.ServerAnonymCommnd) {
			client.Response = GetCmdResponse(MESSAGE_COMMAND_NO_LOGIN, "", ERROR_COMMAND_NO_LOGIN, client)
			return
		}
		//模拟登录
		if !openSession {
			table := client.ServerConfig.DefaultTable
			cacheTable, _ := cache.Cache(table)
			client = &Client{
				Table:      table,
				CacheTable: cacheTable,
				Token:      token,
				IsLogin:    false,
			}
		}
	}
	//获取请求
	command := cmd.GetCmd(request[0])
	if command == nil {
		log.Debug("请求命令不存在：", request[0])
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_FOUND, "", ERROR_COMMAND_NOT_FOUND, client)
		return
	}
	command.InvokeCount++
	//请求校验
	if len(request)-1 < command.MinParam || len(request)-1 > command.MaxParam {
		log.Debug("命令参数个数错误！最大参数个数：", command.MaxParam, "，最小参数个数：", command.MinParam)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	//处理请求
	command.HandlerFunc(client)
}

//判断是否为免登录命令
func IsAnonymCommnd(commnd, anonymCommnd string) bool {
	list := strings.Split(anonymCommnd, ",")
	for _, c := range list {
		if commnd == c {
			return true
		}
	}
	return false
}

//获取会话
func GetSession(client *Client) {
	token := client.Token
	table, _ := cache.GetSysTable()
	item := table.Get(token)
	if item == nil {
		return
	}
	value, _ := item.Value().(Client)
	//复制
	*client = value
}
