package cmd

import (
	//"strconv"

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
	_, err := SyncData(master, client.ServerConfig.ServerId, slavePort, 0)
	if err != nil {
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
	//slaveIp := client.Reqest[2]
	//slavePort := client.Reqest[3]
	//syncPosition, err := strconv.Atoi(client.Reqest[3])
	//	if err != nil {

	//		return
	//	}
	//判断是否为同一台服务器
	if slaveId == client.ServerConfig.ServerId {

		return
	}
	//生成dump文件并发送

	log.Debug("")
}
