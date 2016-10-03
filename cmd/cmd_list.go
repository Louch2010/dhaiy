package cmd

import (
	"container/list"
	"strconv"

	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//LPush处理
func HandleLPushCommand(client *Client) {
	doPush(client, true, false)
}

//LPushX处理
func HandleLPushXCommand(client *Client) {
	doPush(client, true, true)
}

//LRange处理
func HandleLRangeCommand(client *Client) {
	//index转换
	st := client.Reqest[2]
	start, err_st := strconv.Atoi(st)
	en := client.Reqest[3]
	end, err_en := strconv.Atoi(en)
	if err_st != nil || err_en != nil {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		log.Error("LRange处理时数据转换异常，start:", st, "，end：", en, err_st, err_en)
		return
	}
	//获取项
	l := getList(client)
	if l == nil {
		return
	}
	//参数调整
	if start < 0 {
		start = 0
	}
	if end > l.Len() {
		end = l.Len()
	}
	if start > l.Len() || start >= end {
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, nil, nil, client)
		client.Response.DataType = DATA_TYPE_NIL
		return
	}
	//遍历
	ret := make([]string, end-start)
	i := 0
	count := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i >= start && i < end {
			ret[count] = e.Value.(string)
			count++
		}
		i++
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, ret, nil, client)
	client.Response.DataType = DATA_TYPE_LIST
}

//RPush处理
func HandleRPushCommand(client *Client) {
	doPush(client, false, false)
}

//RPushX处理
func HandleRPushXCommand(client *Client) {
	doPush(client, false, true)
}

//LLen处理
func HandleLLenCommand(client *Client) {
	l := getList(client)
	if l == nil {
		return
	}
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, l.Len(), nil, client)
	client.Response.DataType = DATA_TYPE_NUMBER
}

//LPop处理
func HandleLPopCommand(client *Client) {
	doPop(client, true)
}

//RPop处理
func HandleRPopXCommand(client *Client) {
	doPop(client, false)
}

//LIndex处理
func HandleLIndexCommand(client *Client) {
	//index转换
	ind := client.Reqest[2]
	index, err := strconv.Atoi(ind)
	if err != nil {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		log.Error("LIndex处理时数据转换异常，index:", ind, err)
		return
	}
	//获取项
	l := getList(client)
	if l == nil {
		return
	}
	if index >= l.Len() {
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, nil, nil, client)
		client.Response.DataType = DATA_TYPE_NIL
		return
	}
	//遍历
	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		if i == index {
			client.Response = GetCmdResponse(MESSAGE_SUCCESS, e.Value.(string), nil, client)
		}
		i++
	}
}

//LRem处理
func HandleLRemCommand(client *Client) {
	key := client.Reqest[1]
	c := client.Reqest[2]
	value := client.Reqest[3]
	//count转换
	count, err := strconv.Atoi(c)
	if err != nil {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, "", ERROR_COMMAND_PARAM_ERROR, client)
		log.Error("LRem处理时数据转换异常，count:", c, err)
		return
	}
	//获取列表
	l := getList(client)
	if l == nil {
		return
	}
	//	for {
	//		ret := l.Remove(value)
	//		count--
	//		if ret == nil || count == 0 {
	//			break
	//		}
	//	}
	log.Debug(value, count)
	client.CacheTable.Set(key, *l, 0, DATA_TYPE_LIST)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
}

//获取列表信息
func getList(client *Client) *list.List {
	key := client.Reqest[1]
	o := client.CacheTable.Get(key)
	//不存在则直接返回
	if o == nil {
		client.Response = GetCmdResponse(MESSAGE_ITEM_NOT_EXIST, "", ERROR_ITEM_NOT_EXIST, client)
		return nil
	}
	//存在则校验类型并进行数据类型转换
	if o.DataType() != DATA_TYPE_LIST {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		log.Debug("数据类型错误，", o.DataType(), " 不支持使用list方式操作")
		return nil
	}
	l, flag := o.Value().(list.List)
	if !flag {
		client.Response = GetCmdResponse(MESSAGE_COMMAND_NOT_SUPPORT_DATA, "", ERROR_COMMAND_NOT_SUPPORT_DATA, client)
		log.Error("强制类型转换错误！")
		return nil
	}
	return &l
}

//执行插入操作
//left:是否为左插入
//nx:是否判断存在性
func doPush(client *Client, left bool, nx bool) {
	key := client.Reqest[1]
	value := client.Reqest[2]
	l := getList(client)
	if l == nil {
		if client.Response.Code == MESSAGE_ITEM_NOT_EXIST {
			if nx {
				client.Response = GetCmdResponse(MESSAGE_SUCCESS, false, nil, client)
				client.Response.DataType = DATA_TYPE_BOOL
				return
			} else {
				l = list.New()
			}
		} else {
			return
		}
	}
	if left {
		l.PushFront(value)
	} else {
		l.PushBack(value)
	}
	client.CacheTable.Set(key, *l, 0, DATA_TYPE_LIST)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, true, nil, client)
	client.Response.DataType = DATA_TYPE_BOOL
}

//执行pop操作
func doPop(client *Client, left bool) {
	key := client.Reqest[1]
	l := getList(client)
	if l == nil {
		return
	}
	if l.Len() == 0 {
		client.Response = GetCmdResponse(MESSAGE_ITEM_IS_EMPTY, "", ERROR_ITEM_IS_EMPTY, client)
		return
	}
	var e *list.Element
	if left {
		e = l.Front()
	} else {
		e = l.Back()
	}
	l.Remove(e)
	client.CacheTable.Set(key, *l, 0, DATA_TYPE_LIST)
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, e.Value.(string), nil, client)
}
