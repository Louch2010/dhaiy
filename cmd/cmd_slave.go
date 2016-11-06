package cmd

import (
	"strconv"

	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//slaveof命令处理
func HandleSlaveOfCommand(client *Client) {
	master := client.Reqest[1]
	//取消复制
	if master == SLAVE_NONE {
		client.ServerConfig.Master = nil
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
		return
	}
	info := ServerInfo{
		Host:         master,
		Status:       1,
		SyncPosition: 0,
	}
	client.ServerConfig.Master = &info
	//向master发送复制请求
	slavePort := "9527"
	masterId, err := SlaveSendSyncRequest(master, client.ServerConfig.ServerId, slavePort, 0)
	if err != nil {
		client.ServerConfig.Master.ServerId = masterId
		client.Response = GetCmdResponse(MESSAGE_SLAVEOF_REQ_ERROR, "", err, client)
	} else {
		client.ServerConfig.Master.Status = 3 //同步异常
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
	}
}

//syncdata命令处理
func HandleSyncDataCommand(client *Client) {
	//syncdata 本机服务ID、本机ip、本机端口、同步偏移量
	slaveId := client.Reqest[1]
	slaveIp := client.Reqest[2]
	slavePort, err := strconv.Atoi(client.Reqest[3])
	if err != nil {
		log.Error("syncdata命令参数错误，slave端口号：", client.Reqest[3], err)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, err.Error(), ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	syncPosition, err := strconv.Atoi(client.Reqest[4])
	if err != nil {
		log.Error("syncdata命令参数错误，slave同步偏移量：", client.Reqest[4], err)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, err.Error(), ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	//判断是否为同一台服务器
	if slaveId == client.ServerConfig.ServerId {
		log.Error("syncdata命令参数错误，slave和master为同一台主机，serverId：", slaveId)
		client.Response = GetCmdResponse(MESSAGE_COMMAND_PARAM_ERROR, err.Error(), ERROR_COMMAND_PARAM_ERROR, client)
		return
	}
	//记录到slave列表中
	slaveList := client.ServerConfig.SlaveList
	slave := &ServerInfo{
		Host:         slaveIp,
		Port:         slavePort,
		SyncPosition: syncPosition,
		Status:       1,
	}
	slaveList[len(slaveList)] = slave
	log.Debug("启动线程进行后台同步，slave IP：", slaveIp, "，slave 端口：", slavePort, "，同步偏移量：", syncPosition)
	//生成dump文件并发送
	go MasterHandleSyncRequest(client.ServerConfig)
	//返回master的serverId
	client.Response = GetCmdResponse(MESSAGE_SUCCESS, client.ServerConfig.ServerId, nil, client)
}
