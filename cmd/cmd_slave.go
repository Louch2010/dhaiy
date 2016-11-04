package cmd

import (
	. "github.com/louch2010/dhaiy/common"
	"github.com/louch2010/dhaiy/log"
)

//slaveof命令处理
func HandleSlaveOfCommand(client *Client) {
	master := client.Reqest[1]
	//取消复制
	if master == SLAVE_NONE {
		client.Master = nil
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
		return
	}
	info := ServerInfo{
		Host:         master,
		Status:       1,
		SyncPosition: 0,
	}
	client.Master = &info
	//向master发送复制请求
	slavePort := "9527"
	_, err := SyncData(master, client.ServerId, slavePort, 0)
	if err != nil {
		client.Response = GetCmdResponse(MESSAGE_SLAVEOF_REQ_ERROR, "", err, client)
	} else {
		client.Master.Status = 3 //同步异常
		client.Response = GetCmdResponse(MESSAGE_SUCCESS, "", nil, client)
	}
}

//syncdata命令处理
func HandleSyncDataCommand(client *Client) {
	//syncdata 本机服务ID、本机ip、本机端口、同步偏移量
	log.Debug("")
}
