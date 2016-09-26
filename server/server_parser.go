package server

import (
	"strings"

	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//解析请求
func ParserRequest(client *common.Client) {
	token := client.Token
	request := client.Reqest
	log.Debug("开始处理请求，token：", token, "，请求内容为：", request)
	//会话信息校验
	openSession := common.GetSystemConfig().MustBool("client", "openSession", true)
	isLogin := false
	if !client.IsLogin {
		client, isLogin = GetSession(token) //尝试从cache中获取客户端信息
	}
	//没有登录
	if !isLogin {
		//需要登录，而且也不是免登录的命令
		if openSession && !IsAnonymCommnd(request[0]) {
			log.Debug("======没登录")
			client.Response = common.GetServerRespMsg(common.MESSAGE_COMMAND_NO_LOGIN, "", common.ERROR_COMMAND_NO_LOGIN, nil)
			return
		}
		//模拟登录
		if !openSession {
			table := common.GetSystemConfig().MustValue("table", "default", common.DEFAULT_TABLE_NAME)
			cacheTable, _ := cache.Cache(table)
			client = &common.Client{
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

//创建会话
func CreateSession(token string, c *common.Client) bool {
	//缓存登录信息
	table, _ := cache.GetSysTable()
	table.Set(token, c, 0, common.DATA_TYPE_OBJECT)
	//创建表信息
	cache.Cache(c.Table)
	return true
}

//获取会话
func GetSession(token string) (*common.Client, bool) {
	table, _ := cache.GetSysTable()
	item := table.Get(token)
	if item == nil {
		return &common.Client{}, false
	}
	value, falg := item.Value().(common.Client)
	return &value, falg
}

//销毁会话
func DestroySession(token string) bool {
	table, _ := cache.GetSysTable()
	return table.Delete(token)
}

//判断是否为免登录命令
func IsAnonymCommnd(commnd string) bool {
	anonymCommnd := common.SystemConfigFile.MustValue("server", "anonymCommnd", "ping,connect,exit,help")
	list := strings.Split(strings.ToUpper(anonymCommnd), ",")
	for _, c := range list {
		if commnd == c {
			return true
		}
	}
	return false
}
