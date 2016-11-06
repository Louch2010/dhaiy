package common

import (
	"io/ioutil"
	"net"
	"strconv"

	"github.com/louch2010/dhaiy/log"
)

//savle发起同步请求
func SlaveSendSyncRequest(host string, serverId string, localPort string, position int) (string, error) {
	//syncdata 本机服务ID、本机ip、本机端口、同步偏移量
	localIp := GetLocalIp()
	msg := "syncdata " + serverId + " " + localIp + " " + localPort + " " + strconv.Itoa(position)
	return SendSocketMessage(host, msg)
}

//master处理同步请求
func MasterHandleSyncRequest(config *ServerConfig) {
	//缓存dummp文件

	//发送dump文件

	//发送队列中的命令

}

//master分发写命令
func MasterCopyCommandToSlave(slaves []*ServerInfo, cmd string) {
	for _, slave := range slaves {
		if slave.Status == 1 || slave.Status == 3 { //正在首次同步或同步异常时，将命令缓存到列表中

		} else if slave.Status == 2 { //同步中时，则直接发送
			resp, err := SendSocketMessage(slave.Host+":"+strconv.Itoa(slave.Port), cmd)
			if err != nil && resp != MESSAGE_SUCCESS { //没有同步成功，则也将命令缓存到列表中
				log.Error("命令同步失败，命令原文：", cmd, "，错误信息：", resp, err)
			}
		}
	}
}

//发送dump文件
func SendDumpFile() {

}

//接收dump文件
func ReceiveDumpFile() {

}

//socket发送信息
func SendSocketMessage(host, msg string) (string, error) {
	conn, err := createSocketConnection(host)
	if err != nil {
		return "", err
	}
	//向服务端写数据
	log.Debug("向服务器发送socket信息，host：", host, "，内容：", msg)
	_, err = conn.Write([]byte(msg))
	if err != nil {
		return "", err
	}
	//读取服务端的响应
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		return "", err
	}
	resp := string(result)
	conn.Close()
	log.Debug("向服务器发送socket请求完成，响应内容为：", resp)
	return resp, nil
}

//socket发送文件
func SendSocketFile(host, filePath string) error {
	conn, err := createSocketConnection(host)
	if err != nil {
		return err
	}
	//发送文件

	conn.Close()
	return nil
}

//创建socket连接
func createSocketConnection(host string) (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Error("地址解析失败！", host, err)
		return nil, err
	}
	//连接服务器
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Error("连接服务器失败！", err)
		return nil, err
	}
	return conn, nil
}
