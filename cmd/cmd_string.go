package cmd

import (
	"strconv"
	"time"

	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//Set命令处理
func HandleSetCommand(client *Client) {
	args := client.Reqest[1:]
	var liveTime int = 0
	if len(args) == 3 {
		t, err := strconv.Atoi(args[2])
		if err != nil {
			log.Info("参数转换错误，liveTime：", args[2], err)
			client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
			return
		}
		liveTime = t
	}
	//增加缓存项
	item := client.CacheTable.Set(args[0], args[1], time.Duration(liveTime)*time.Second, DATA_TYPE_STRING)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, item, nil, client)
}

//Get命令处理
func HandleGetCommand(client *Client) {
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	if item == nil {
		client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
		return
	}
	//数据类型校验
	if item.DataType() != DATA_TYPE_STRING {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		return
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, item.Value(), nil, client)
}

//Append命令处理
func HandleAppendCommand(client *Client) {
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	//不存在，则设置
	if item == nil {
		client.CacheTable.Set(args[0], args[1], 0, DATA_TYPE_STRING)
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, args[1], nil, client)
		return
	}
	//数据类型校验
	if item.DataType() != DATA_TYPE_STRING {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		return
	}
	v := item.Value().(string) + args[1]
	item.SetValue(v)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, v, nil, client)
}

//StrLen命令处理
func HandleStrLenCommand(client *Client) {
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	length := 0
	if item != nil {
		//数据类型校验
		if item.DataType() != DATA_TYPE_STRING {
			client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
			return
		}
		length = len(item.Value().(string))
	}
	response := GetCmdResponse(MESSAGE_SUCCESS, length, nil, client)
	response.DataType = DATA_TYPE_NUMBER
	client.Response = response
}

//SetNx命令处理
func HandleSetNxCommand(client *Client) {
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	flag := false
	//不存在，则设置
	if item == nil {
		flag = true
		client.CacheTable.Set(args[0], args[1], 0, DATA_TYPE_STRING)
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, flag, nil, client)
	client.Response.DataType = DATA_TYPE_BOOL
}
