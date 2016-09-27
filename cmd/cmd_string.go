package cmd

import (
	"strconv"
	"time"

	"github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//Set命令处理
func HandleSetCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 2, 3)
	if !check {
		return resp
	}
	var liveTime int = 0
	if len(args) == 3 {
		t, err := strconv.Atoi(args[2])
		if err != nil {
			log.Info("参数转换错误，liveTime：", args[2], err)
			return common.GetServerRespMsg(common.MESSAGE_COMMAND_PARAM_ERROR, "", common.ERROR_COMMAND_PARAM_ERROR, client)
		}
		liveTime = t
	}
	//增加缓存项
	item := client.CacheTable.Set(args[0], args[1], time.Duration(liveTime)*time.Second, common.DATA_TYPE_STRING)
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, item, nil, client)
}

//Get命令处理
func HandleGetCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 1, 1)
	if !check {
		return resp
	}
	item := client.CacheTable.Get(args[0])
	if item == nil {
		return common.GetServerRespMsg(common.MESSAGE_ITEM_NOT_EXIST, "", common.ERROR_ITEM_NOT_EXIST, client)
	}
	//数据类型校验
	if item.DataType() != common.DATA_TYPE_STRING {
		return common.GetServerRespMsg(common.MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", common.ERROR_COMMAND_NOT_SUPPORT_DATA, client)
	}
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, item.Value(), nil, client)
}

//Append命令处理
func HandleAppendCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 2, 2)
	if !check {
		return resp
	}
	item := client.CacheTable.Get(args[0])
	//不存在，则设置
	if item == nil {
		client.CacheTable.Set(args[0], args[1], 0, common.DATA_TYPE_STRING)
		return common.GetServerRespMsg(common.MESSAGE_SUCCESS, args[1], nil, client)
	}
	//数据类型校验
	if item.DataType() != common.DATA_TYPE_STRING {
		return common.GetServerRespMsg(common.MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", common.ERROR_COMMAND_NOT_SUPPORT_DATA, client)
	}
	v := item.Value().(string) + args[1]
	item.SetValue(v)
	return common.GetServerRespMsg(common.MESSAGE_SUCCESS, v, nil, client)
}

//StrLen命令处理
func HandleStrLenCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 1, 1)
	if !check {
		return resp
	}
	item := client.CacheTable.Get(args[0])
	length := 0
	if item != nil {
		//数据类型校验
		if item.DataType() != common.DATA_TYPE_STRING {
			return common.GetServerRespMsg(common.MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", common.ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		}
		length = len(item.Value().(string))
	}
	response := common.GetServerRespMsg(common.MESSAGE_SUCCESS, length, nil, client)
	response.DataType = common.DATA_TYPE_NUMBER
	return response
}

//SetNx命令处理
func HandleSetNxCommnd(body string, client *common.Client) common.ServerRespMsg {
	//请求体校验
	args, resp, check := initParam(body, 2, 2)
	if !check {
		return resp
	}
	item := client.CacheTable.Get(args[0])
	flag := false
	//不存在，则设置
	if item == nil {
		flag = true
		client.CacheTable.Set(args[0], args[1], 0, common.DATA_TYPE_STRING)
	}
	response := common.GetServerRespMsg(common.MESSAGE_SUCCESS, flag, nil, client)
	response.DataType = common.DATA_TYPE_BOOL
	return response
}
