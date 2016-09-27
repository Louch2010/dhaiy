package cmd

import (
	"container/list"
	"strconv"
	"strings"

	"github.com/louch2010/dhaiy/cache"
	"github.com/louch2010/dhaiy/common"
	//"github.com/louch2010/dhaiy/gdb"
	"github.com/louch2010/dhaiy/log"
	"github.com/louch2010/goutil"
)

//ping命令处理
func HandlePingCommnd(client *common.Client) {
	client.Response = common.GetServerRespMsg(common.MESSAGE_SUCCESS, common.MESSAGE_PONG, nil, client)
}

//帮助命令处理
func HandleHelpCommnd(client *common.Client) {
	response := ""
	help := common.GetHelpConfig()
	if len(client.Reqest) == 0 { //没有请求体，则显示所有命令名称
		for index, sec := range help.GetSectionList() {
			response += "[" + strconv.Itoa(index+1) + "] " + sec + "\r\n"
		}
		response += "use 'help commnd' to see detail info"
	} else {
		cmd := strings.ToLower(client.Reqest[1])
		sec, err := help.GetSection(cmd)
		if err != nil {
			response = "no help for the commnd"
		} else {
			response += "[" + cmd + "]\r\n"
			for k, v := range sec {
				response += k + ": " + v + "\r\n"
			}
		}
	}
	client.Response = common.GetServerRespMsg(common.MESSAGE_SUCCESS, response, nil, client)
}

//连接命令处理connect [-t'table'] [-a'pwd'] [-i'ip'] [-p'port'] [-e'e1,e2...']
func HandleConnectCommnd(client *common.Client) {
	table := common.GetSystemConfig().MustValue("table", "default", common.DEFAULT_TABLE_NAME)
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
	syspwd := common.GetSystemConfig().MustValue("server", "password", "")
	if len(syspwd) > 0 {
		if len(pwd) == 0 {
			client.Response = common.GetServerRespMsg(common.MESSAGE_NO_PWD, "", common.ERROR_AUTHORITY_NO_PWD, nil)
			return
		}
		if syspwd != pwd {
			client.Response = common.GetServerRespMsg(common.MESSAGE_PWD_ERROR, "", common.ERROR_AUTHORITY_PWD_ERROR, nil)
			return
		}
	}
	//端口校验
	portInt := 0
	if len(port) > 0 {
		p, err := strconv.Atoi(port)
		if err != nil {
			log.Info("端口转换错误，", err)
			client.Response = common.GetServerRespMsg(common.MESSAGE_PORT_ERROR, "", common.ERROR_PORT_ERROR, nil)
			return
		}
		portInt = p
	}
	//协议校验
	if len(protocol) > 0 && protocol != common.PROTOCOL_RESPONSE_JSON && protocol != common.PROTOCOL_RESPONSE_TERMINAL {
		log.Info("协议错误：", protocol)
		client.Response = common.GetServerRespMsg(common.MESSAGE_PROTOCOL_ERROR, "", common.ERROR_PROTOCOL_ERROR, nil)
		return
	}
	//获取表
	cacheTable, err := cache.Cache(table)
	if err != nil {
		log.Error("连接时获取表失败！", err)
		client.Response = common.GetServerRespMsg(common.MESSAGE_ERROR, "", common.ERROR_SYSTEM, nil)
		return
	}
	//存储连接信息
	client = &common.Client{
		Host:        ip,
		Port:        portInt,
		Table:       table,
		CacheTable:  cacheTable,
		ListenEvent: strings.Split(event, ","),
		Protocol:    protocol,
		Token:       token,
	}
	CreateSession(token, client)
	client.Response = common.GetServerRespMsg(common.MESSAGE_SUCCESS, token, nil, client)
}

//Delete命令处理
func HandleDeleteCommnd(client *common.Client) {
	request := client.Reqest
	if client.CacheTable.Delete(request[1]) {
		client.Response = common.GetServerRespMsg(common.MESSAGE_SUCCESS, "", nil, client)
		return
	}
	client.Response = common.GetServerRespMsg(common.MESSAGE_ITEM_NOT_EXIST, "", common.ERROR_ITEM_NOT_EXIST, client)
}

//Exist命令处理
func HandleExistCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 1, 1)
	if !check {
		return resp
	}
	response := common.GetServerRespMsg(common.MESSAGE_SUCCESS, client.CacheTable.IsExist(args[0]), nil, client)
	response.DataType = common.DATA_TYPE_BOOL
	return response
}

//切换表
func HandleUseCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 1, 1)
	if !check {
		return resp
	}
	cacheTable, err := cache.Cache(args[0])
	if err != nil {
		log.Error("切换表时获取表失败！", err)
		return common.GetServerRespMsg(common.MESSAGE_ERROR, "", common.ERROR_SYSTEM, nil)
	}
	client.Table = args[0]
	client.CacheTable = cacheTable
	if CreateSession(client.Token, client) {
		return common.GetServerRespMsg(common.MESSAGE_SUCCESS, "", nil, client)
	}
	return common.GetServerRespMsg(common.MESSAGE_ERROR, "", common.ERROR_SYSTEM, client)
}

//显示表信息
func HandleShowtCommnd(body string, client *common.Client) common.ServerRespMsg {
	response := ""
	args, resp, check := initParam(body, 0, 1)
	if !check {
		return resp
	}
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
			return common.GetServerRespMsg(common.MESSAGE_TABLE_NOT_EXIST, response, common.ERROR_TABLE_NOT_EXIST, client)
		}
		response += "name:" + table.Name() + "\r\n"
		response += "itemCount: " + strconv.Itoa(table.ItemCount()) + "\r\n"
		response += "createTime: " + goutil.DateUtil().TimeFullFormat(table.CreateTime()) + "\r\n"
		response += "lastAccessTime: " + goutil.DateUtil().TimeFullFormat(table.LastAccessTime()) + "\r\n"
		response += "lastModifyTime: " + goutil.DateUtil().TimeFullFormat(table.LastModifyTime()) + "\r\n"
		response += "accessCount: " + strconv.FormatInt(table.AccessCount(), 10)
	}
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, response, nil, client)
}

//显示项信息
func HandleShowiCommnd(body string, client *common.Client) common.ServerRespMsg {
	response := ""
	table, _ := cache.Cache(client.Table)
	args, resp, check := initParam(body, 0, 1)
	if !check {
		return resp
	}
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
			return common.GetServerRespMsg(common.MESSAGE_ITEM_NOT_EXIST, "", common.ERROR_ITEM_NOT_EXIST, client)
		}
		response += "key: " + item.Key() + "\r\n"
		//response += "value: " + toString(item) + "\r\n"
		response += "liveTime: " + item.LiveTime().String() + "\r\n"
		response += "createTime: " + goutil.DateUtil().TimeFullFormat(item.CreateTime()) + "\r\n"
		response += "lastAccessTime: " + goutil.DateUtil().TimeFullFormat(item.LastAccessTime()) + "\r\n"
		response += "accessCount: " + strconv.FormatInt(item.AccessCount(), 10) + "\r\n"
		response += "dataType: " + item.DataType() + "\r\n"
	}
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, response, nil, client)
}

//服务器信息
func HandleInfoCommnd(body string, client *common.Client) common.ServerRespMsg {
	_, resp, check := initParam(body, 0, 0)
	if !check {
		return resp
	}
	//获取默认section中的信息
	info, _ := common.GetSystemConfig().GetSection("")
	response := ""
	for k, v := range info {
		response += k + ": " + v + "\r\n"
	}
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, response, nil, client)
}

//后台保存gdb文件
func HandleBgSaveCommnd(body string, client *common.Client) common.ServerRespMsg {
	_, resp, check := initParam(body, 0, 0)
	if !check {
		return resp
	}
	//go gdb.SaveDB()
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, "", nil, client)
}

//初始化参数
func initParam(body string, minBodyLen int, maxBodyLen int) ([]string, common.ServerRespMsg, bool) {
	result := make([]string, 0)
	if len(body) == 0 {
		if minBodyLen != 0 {
			return nil, common.GetServerRespMsg(common.MESSAGE_COMMAND_PARAM_ERROR, "", common.ERROR_COMMAND_PARAM_ERROR, nil), false
		}
		return result, common.GetServerRespMsg(common.MESSAGE_SUCCESS, "", nil, nil), true
	}
	//如果包含引号，则需要特殊处理
	if strings.Contains(body, "\"") {
		l := list.New()
		open := false
		buffer := ""
		for _, c := range body {
			if '"' == c {
				if open {
					l.PushBack(buffer)
					buffer = ""
				}
				open = !open
				continue
			}
			if ' ' == c && !open {
				if len(buffer) > 0 {
					l.PushBack(buffer)
					buffer = ""
				}
				continue
			}
			buffer += string(c)
		}
		result = make([]string, l.Len())
		var i = 0
		for e := l.Front(); e != nil; e = e.Next() {
			result[i] = e.Value.(string)
			i = i + 1
		}
	} else {
		body = strings.Replace(body, "  ", " ", 99)
		result = strings.Split(body, " ")
	}
	log.Debug("初始化请求参数完成，请求参数为：", result, "，长度为：", len(result))
	if len(result) < minBodyLen || len(result) > maxBodyLen {
		return nil, common.GetServerRespMsg(common.MESSAGE_COMMAND_PARAM_ERROR, "", common.ERROR_COMMAND_PARAM_ERROR, nil), false
	}
	return result, common.GetServerRespMsg(common.MESSAGE_SUCCESS, "", nil, nil), true
}
