package cmd

import (
	"strconv"
	"strings"

	"github.com/louch2010/dhaiy/cache"
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/gdb"
	"github.com/louch2010/dhaiy/log"
	"github.com/louch2010/goutil"
)

//ping命令处理
func HandlePingCommand(client *Client) {
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, MESSAGE_PONG, nil, client)
}

//帮助命令处理
func HandleHelpCommand(client *Client) {
	response := ""
	args := client.Reqest[1:]
	help := GetHelpConfig()
	if len(args) == 0 { //没有请求体，则显示所有命令名称
		for index, sec := range help.GetSectionList() {
			response += "[" + strconv.Itoa(index+1) + "] " + sec + "\r\n"
		}
		response += "use 'help commnd' to see detail info"
	} else {
		sec, err := help.GetSection(args[0])
		if err != nil {
			response = "no help for the commnd"
		} else {
			response += "[" + args[0] + "]\r\n"
			for k, v := range sec {
				response += k + ": " + v + "\r\n"
			}
			//显示调用信息
			cmd := GetCmd(args[0])
			response += "MinParam: " + strconv.Itoa(cmd.MinParam) + "\r\n"
			response += "MaxParam: " + strconv.Itoa(cmd.MaxParam) + "\r\n"
			response += "IsWrite: " + strconv.FormatBool(cmd.IsWrite) + "\r\n"
			response += "InvokeCount: " + strconv.FormatInt(cmd.InvokeCount, 10)
		}
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, response, nil, client)
}

//连接命令处理connect [-t'table'] [-a'pwd'] [-i'ip'] [-p'port'] [-e'e1,e2...']
func HandleConnectCommand(client *Client) {
	log.Debug("处理connect请求")
	table := client.ServerConfig.DefaultTable
	token := client.Token
	var pwd, ip, port, event, protocol string
	args := client.Reqest[1:]
	for _, arg := range args {
		//参数长度小于3或不是以-开头，说明参数不对，直接跳过
		if len(arg) < 3 || !strings.HasPrefix(arg, "-") {
			continue
		}
		paramType := arg[1]
		param := arg[2:len(arg)]
		switch paramType {
		case 't':
			table = param
			break
		case 'a':
			pwd = param
			break
			break
		case 'i':
			ip = param
			break
		case 'p':
			port = param
			break
		case 'e':
			event = param
			break
		case 'm':
			protocol = param
			break
		default:
		}
	}
	//密码校验
	syspwd := client.ServerConfig.ServerPassword
	if len(syspwd) > 0 {
		if len(pwd) == 0 {
			client.Response = GetCmdResponse(MESSAGE_NO_PWD, "", ERROR_AUTHORITY_NO_PWD, nil)
			return
		}
		if syspwd != pwd {
			client.Response = GetCmdResponse(MESSAGE_PWD_ERROR, "", ERROR_AUTHORITY_PWD_ERROR, nil)
			return
		}
	}
	//端口校验
	portInt := 0
	if len(port) > 0 {
		p, err := strconv.Atoi(port)
		if err != nil {
			log.Info("端口转换错误，", err)
			client.Response = GetCmdResponse(MESSAGE_PORT_ERROR, "", ERROR_PORT_ERROR, nil)
			return
		}
		portInt = p
	}
	//协议校验
	if len(protocol) > 0 && protocol != PROTOCOL_RESPONSE_JSON && protocol != PROTOCOL_RESPONSE_TERMINAL {
		log.Info("协议错误：", protocol)
		client.Response = GetCmdResponse(MESSAGE_PROTOCOL_ERROR, "", ERROR_PROTOCOL_ERROR, nil)
		return
	}
	//获取表
	cacheTable, err := cache.Cache(table)
	if err != nil {
		log.Error("连接时获取表失败！", err)
		client.Response = GetCmdResponse(MESSAGE_ERROR, "", ERROR_SYSTEM, nil)
		return
	}
	//存储连接信息
	*client = Client{
		Host:        ip,
		Port:        portInt,
		Table:       table,
		CacheTable:  cacheTable,
		ListenEvent: strings.Split(event, ","),
		Protocol:    protocol,
		Token:       token,
		IsLogin:     true,
	}
	CreateSession(token, client)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, token, nil, client)
}

//Exit命令处理
func HandleExitCommand(client *Client) {
	DestroySession(client.Token)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
	client.Response.Clo = true
}

//Delete命令处理
func HandleDelCommand(client *Client) {
	request := client.Reqest
	if client.CacheTable.Delete(request[1]) {
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
		return
	}
	client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
}

//Exist命令处理
func HandleExistCommand(client *Client) {
	arg := client.Reqest[1]
	response := GetCmdResponse(MESSAGE_SUCCESS, client.CacheTable.IsExist(arg), nil, client)
	response.DataType = DATA_TYPE_BOOL
	client.Response = response
}

//切换表
func HandleSelectCommand(client *Client) {
	arg := client.Reqest[1]
	cacheTable, err := cache.Cache(arg)
	if err != nil {
		log.Error("切换表时获取表失败！", err)
		client.Response = GetCmdResponse(MESSAGE_ERROR, "", ERROR_SYSTEM, nil)
		return
	}
	client.Table = arg
	client.CacheTable = cacheTable
	if CreateSession(client.Token, client) {
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
		return
	}
	client.Response = GetCmdResponse(MESSAGE_ERROR, "", ERROR_SYSTEM, client)
}

//显示表信息
func HandleShowtCommand(client *Client) {
	response := ""
	args := client.Reqest[1:]
	if len(args) == 0 { //没有请求体，则显示所有表名
		list := cache.GetCacheTables()
		index := 1
		for k, _ := range list {
			if k == client.Table {
				response += "[* " + strconv.Itoa(index) + "] "
			} else {
				response += "[" + strconv.Itoa(index) + "] "
			}
			response += k + "\r\n"
			index += 1
		}
		response += "use 'showt tableName' to see detail info"
	} else {
		table, ok := cache.GetCacheTable(args[0])
		if !ok {
			client.Response = GetCmdResponse(MESSAGE_TABLE_NOT_EXIST, response, ERROR_TABLE_NOT_EXIST, client)
			return
		}
		response += "name:" + table.Name() + "\r\n"
		response += "itemCount: " + strconv.Itoa(table.ItemCount()) + "\r\n"
		response += "createTime: " + goutil.DateUtil().TimeFullFormat(table.CreateTime()) + "\r\n"
		response += "lastAccessTime: " + goutil.DateUtil().TimeFullFormat(table.LastAccessTime()) + "\r\n"
		response += "lastModifyTime: " + goutil.DateUtil().TimeFullFormat(table.LastModifyTime()) + "\r\n"
		response += "accessCount: " + strconv.FormatInt(table.AccessCount(), 10)
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, response, nil, client)
}

//显示项信息
func HandleShowiCommand(client *Client) {
	response := ""
	table, _ := cache.Cache(client.Table)
	args := client.Reqest[1:]
	if len(args) == 0 { //没有请求体，则显示所有项
		index := 1
		for k, _ := range table.GetItems() {
			response += "[" + strconv.Itoa(index) + "] " + k + "\r\n"
			index += 1
		}
		response += "use 'showi key' to see detail info"
	} else {
		item := table.Get(args[0])
		if item == nil {
			client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
			return
		}
		response += "key: " + item.Key() + "\r\n"
		//response += "value: " + toString(item) + "\r\n"
		response += "liveTime: " + item.LiveTime().String() + "\r\n"
		response += "createTime: " + goutil.DateUtil().TimeFullFormat(item.CreateTime()) + "\r\n"
		response += "lastAccessTime: " + goutil.DateUtil().TimeFullFormat(item.LastAccessTime()) + "\r\n"
		response += "accessCount: " + strconv.FormatInt(item.AccessCount(), 10) + "\r\n"
		response += "dataType: " + item.DataType() + "\r\n"
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, response, nil, client)
}

//服务器信息
func HandleInfoCommand(client *Client) {
	//获取默认section中的信息
	info, _ := GetSystemConfig().GetSection("")
	response := ""
	for k, v := range info {
		response += k + ": " + v + "\r\n"
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, response, nil, client)
}

//flushdb处理（当前数据库）
func HandleFlushDBCommand(client *Client) {
	m := make(map[string]*CacheItem)
	client.CacheTable.Items(m)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
}

//flushall处理（所有库）
func HandleFlushAllCommand(client *Client) {
	db := cache.GetCacheTables()
	sysTable := client.ServerConfig.ServerSysTable
	for k, v := range db {
		if k != sysTable {
			m := make(map[string]*CacheItem)
			v.Items(m)
		}
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
}

//后台保存gdb文件
func HandleBgSaveCommand(client *Client) {
	go gdb.SaveDB()
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
}

//创建会话
func CreateSession(token string, c *Client) bool {
	//缓存登录信息
	table, _ := cache.GetSysTable()
	table.Set(token, c, 0, DATA_TYPE_OBJECT)
	//创建表信息
	cache.Cache(c.Table)
	return true
}

//销毁会话
func DestroySession(token string) bool {
	table, _ := cache.GetSysTable()
	return table.Delete(token)
}
