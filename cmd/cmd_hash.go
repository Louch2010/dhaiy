package cmd

import (
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//HSet命令处理
func HandleHSetCommand(client *Client) {
	key := client.Reqest[1]
	field := client.Reqest[2]
	value := client.Reqest[3]
	m := getMap(client)
	if m == nil {
		if client.Response.Code != MESSAGE_ITEM_NOT_EXIST {
			return
		}
		m = make(map[string]string)
	}
	m[field] = value
	client.CacheTable.Set(key, m, 0, DATA_TYPE_MAP)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
}

//HGet命令处理
func HandleHGetCommand(client *Client) {
	field := client.Reqest[2]
	m := getMap(client)
	if m == nil {
		return
	}
	value, ok := m[field]
	if !ok {
		client.Response = GetCmdResponse(MESSAGE_FIELD_NOT_EXIST, "", ERROR_FIELD_NOT_EXIST, client)
		return
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, value, nil, client)
}

//HDel命令处理
func HandleHDelCommand(client *Client) {
	key := client.Reqest[1]
	field := client.Reqest[2]
	m := getMap(client)
	if m == nil {
		return
	}
	_, ok := m[field]
	if !ok {
		client.Response = GetCmdResponse(MESSAGE_FIELD_NOT_EXIST, "", ERROR_FIELD_NOT_EXIST, client)
		return
	}
	delete(m, field)
	client.CacheTable.Set(key, m, 0, DATA_TYPE_MAP)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
}

//HExists命令处理
func HandleHExistsCommand(client *Client) {
	field := client.Reqest[2]
	m := getMap(client)
	if m == nil {
		return
	}
	_, ok := m[field]
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, ok, nil, client)
	client.Response.DataType = DATA_TYPE_BOOL
}

//HLen命令处理
func HandleHLenCommand(client *Client) {
	m := getMap(client)
	if m == nil {
		return
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, len(m), nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//HKeys命令处理
func HandleHKeysCommand(client *Client) {
	m := getMap(client)
	if m == nil {
		return
	}
	l := make([]string, len(m))
	i := 0
	for k, _ := range m {
		l[i] = k
		i++
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, l, nil, client)
	client.Response.DataType = DATA_TYPE_LIST
}

//HSetNx命令处理
func HandleHSetNxCommand(client *Client) {
	key := client.Reqest[1]
	field := client.Reqest[2]
	value := client.Reqest[3]
	m := getMap(client)
	if m == nil {
		if client.Response.Code != MESSAGE_ITEM_NOT_EXIST {
			return
		}
		m = make(map[string]string)
	}
	flag := false
	_, ok := m[field]
	//不存在则设置
	if !ok {
		flag = true
		m[field] = value
		client.CacheTable.Set(key, m, 0, DATA_TYPE_MAP)
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, flag, nil, client)
	client.Response.DataType = DATA_TYPE_BOOL
}

//HVals命令处理
func HandleHValsCommand(client *Client) {
	m := getMap(client)
	if m == nil {
		return
	}
	l := make([]string, len(m))
	i := 0
	for _, v := range m {
		l[i] = v
		i++
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, l, nil, client)
	client.Response.DataType = DATA_TYPE_LIST
}

//获取表信息
func getMap(client *Client) map[string]string {
	key := client.Reqest[1]
	o := client.CacheTable.Get(key)
	//不存在则直接返回
	if o == nil {
		client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
		return nil
	}
	//存在则校验类型并进行数据类型转换
	if o.DataType() != DATA_TYPE_MAP {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		log.Debug("数据类型错误，", o.DataType(), " 不支持使用map方式操作")
		return nil
	}
	m, flag := o.Value().(map[string]string)
	if !flag {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		log.Error("强制类型转换错误！")
		return nil
	}
	return m
}
