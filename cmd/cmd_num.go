package cmd

import (
	"strconv"
	"time"

	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//NSet命令处理
func HandleNSetCommand(client *Client) {
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
	f, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		log.Info("参数类型错误，nset value：", args[0], err)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		return
	}
	//增加缓存项
	item := client.CacheTable.Set(args[0], f, time.Duration(liveTime)*time.Second, DATA_TYPE_NUMBER)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, item, nil, client)
}

//NGet命令处理
func HandleNGetCommand(client *Client) {
	//请求体校验
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	if item == nil {
		client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
		return
	}
	//数据类型校验
	if item.DataType() != DATA_TYPE_NUMBER {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		return
	}
	response := GetCmdResponse(MESSAGE_SUCCESS, item.Value(), nil, client)
	response.DataType = DATA_TYPE_NUMBER
	client.Response = response
}

//Incr命令处理
func HandleIncrCommand(client *Client) {
	//请求体校验
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	var v float64 = 1
	//不存在，则设置为0，存在增加1
	if item == nil {
		client.CacheTable.Set(args[0], v, 0, DATA_TYPE_NUMBER)
	} else {
		//数据类型校验
		if item.DataType() != DATA_TYPE_NUMBER {
			client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
			return
		}
		o, _ := item.Value().(float64)
		v = o + v
		item.SetValue(v)
	}
	response := GetCmdResponse(MESSAGE_SUCCESS, v, nil, client)
	response.DataType = DATA_TYPE_NUMBER
	client.Response = response
}

//IncrBy命令处理
func HandleIncrByCommand(client *Client) {
	//请求体校验
	args := client.Reqest[1:]
	item := client.CacheTable.Get(args[0])
	v, _ := strconv.ParseFloat(args[1], 10)
	//不存在，则设置为0，存在增加1
	if item == nil {
		client.CacheTable.Set(args[0], v, 0, DATA_TYPE_NUMBER)
	} else {
		//数据类型校验
		if item.DataType() != DATA_TYPE_NUMBER {
			client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
			return
		}
		o, _ := item.Value().(float64)
		v = o + v
		item.SetValue(v)
	}
	response := GetCmdResponse(MESSAGE_SUCCESS, v, nil, client)
	response.DataType = DATA_TYPE_NUMBER
	client.Response = response
}
