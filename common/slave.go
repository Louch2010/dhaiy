package common

import (
	"io/ioutil"
	"net"
	"strconv"

	"github.com/louch2010/dhaiy/log"
)

func SyncData(host string, serverId string, localPort string, position int) (string, error) {
	//syncdata 本机服务ID、本机ip、本机端口、同步偏移量
	localIp := GetLocalIp()
	msg := "syncdata " + serverId + " " + localIp + " " + localPort + " " + strconv.Itoa(position)
	return SendSocketMessage(host, msg)
}

//发送socket信息
func SendSocketMessage(host, msg string) (string, error) {
	addr, err := net.ResolveTCPAddr("tcp4", host)
	if err != nil {
		log.Error("地址解析失败！", host, err)
		return "", err
	}
	//连接服务器
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Error("连接服务器失败！", err)
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
