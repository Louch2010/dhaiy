package cmd

import (
	"strconv"

	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//ZAdd命令处理
func HandleZAddCommand(client *Client) {
	key := client.Reqest[1]
	value := client.Reqest[3]
	score, err := strconv.Atoi(client.Reqest[2])
	if err != nil {
		log.Info("参数转换错误，score：", client.Reqest[2], err)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	//获取set
	s := getSet(client)
	if s == nil {
		if client.Response.Code != MESSAGE_ITEM_NOT_EXIST {
			return
		}
		s = NewSortSet()
	}
	s.Add(score, value)
	client.CacheTable.Set(key, *s, 0, DATA_TYPE_ZSET)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, true, nil, client)
	client.Response.DataType = DATA_TYPE_BOOL
}

//ZCard命令处理
func HandleZCardCommand(client *Client) {
	//获取set
	s := getSet(client)
	if s == nil {
		return
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, s.Len(), nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//ZCount命令处理
func HandleZCountCommand(client *Client) {
	min, err := strconv.Atoi(client.Reqest[2])
	if err != nil {
		log.Info("参数转换错误，min：", client.Reqest[2], err)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	max, err := strconv.Atoi(client.Reqest[3])
	if err != nil {
		log.Info("参数转换错误，max：", client.Reqest[3], err)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	//获取set
	s := getSet(client)
	if s == nil {
		return
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, s.Count(min, max), nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//ZScore命令处理
func HandleZScoreCommand(client *Client) {
	//获取set
	s := getSet(client)
	if s == nil {
		return
	}
	value := client.Reqest[2]
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, s.Score(value), nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//ZRank命令处理
func HandleZRankCommand(client *Client) {
	//获取set
	s := getSet(client)
	if s == nil {
		return
	}
	value := client.Reqest[2]
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, s.Rank(value), nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//ZRem命令处理
func HandleZRemkCommand(client *Client) {
	key := client.Reqest[1]
	//获取set
	s := getSet(client)
	if s == nil {
		return
	}
	value := client.Reqest[2]
	s.RemoveData(value)
	client.CacheTable.Set(key, *s, 0, DATA_TYPE_ZSET)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, true, nil, client)
	client.Response.DataType = DATA_TYPE_BOOL
}

//ZIncrBy命令处理
func HandleZIncrByCommand(client *Client) {
	key := client.Reqest[1]
	//获取set
	s := getSet(client)
	if s == nil {
		return
	}
	value := client.Reqest[2]
	s.RemoveData(value)
	client.CacheTable.Set(key, *s, 0, DATA_TYPE_ZSET)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, true, nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//获取sort信息
func getSet(client *Client) *SortSet {
	key := client.Reqest[1]
	o := client.CacheTable.Get(key)
	//不存在则直接返回
	if o == nil {
		client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
		return nil
	}
	//存在则校验类型并进行数据类型转换
	if o.DataType() != DATA_TYPE_ZSET {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		log.Debug("数据类型错误，", o.DataType(), " 不支持使用set方式操作")
		return nil
	}
	l, flag := o.Value().(SortSet)
	if !flag {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		log.Error("强制类型转换错误！")
		return nil
	}
	return &l
}
