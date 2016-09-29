package server

import (
	"strings"

	"github.com/louch2010/dhaiy/cache"
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//解析请求
func ParserRequest(client *Client) {
	token := client.Token
	request := client.Reqest
	log.Debug("开始处理请求，token：", token, "，请求内容为：", request)
	//会话信息校验
	isLogin := false
	if !client.IsLogin {
		client, isLogin = GetSession(token) //尝试从cache中获取客户端信息
	}
	//没有登录
	if !isLogin {
		openSession := GetSystemConfig().MustBool("client", "openSession", true)
		//需要登录，而且也不是免登录的命令
		if openSession && !IsAnonymCommnd(request[0]) {
			client.Response = GetServerRespMsg(MESSAGE_COMMAND_NO_LOGIN, "", ERROR_COMMAND_NO_LOGIN, nil)
			log.Debug("======没登录", client.Response)
			return
		}
		//模拟登录
		if !openSession {
			table := GetSystemConfig().MustValue("table", "default", DEFAULT_TABLE_NAME)
			cacheTable, _ := cache.Cache(table)
			client = &Client{
				Table:      table,
				CacheTable: cacheTable,
				Token:      token,
				IsLogin:    false,
			}
		}
	}
	log.Debug("会话信息：", *client)
	log.Debug("请求信息：", client.Reqest)
}

//判断是否为免登录命令
func IsAnonymCommnd(commnd string) bool {
	anonymCommnd := SystemConfigFile.MustValue("server", "anonymCommnd", "ping,connect,exit,help")
	list := strings.Split(strings.ToUpper(anonymCommnd), ",")
	for _, c := range list {
		if commnd == c {
			return true
		}
	}
	return false
}

//获取会话
func GetSession(token string) (*Client, bool) {
	table, _ := cache.GetSysTable()
	item := table.Get(token)
	if item == nil {
		return &Client{}, false
	}
	value, falg := item.Value().(Client)
	return &value, falg
}
