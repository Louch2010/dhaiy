package common

//默认常量
const (
	//默认表名
	DEFAULT_TABLE_NAME = "default"
	//配置文件路径
	CONFIG_SYSTEM_FILE = "sys.ini"
)

//事件
const (
	//缓存项新增
	EVENT_ITEM_ADD = "EVENT_ITEM_ADD"
	//缓存项修改
	EVENT_ITEM_MODIFY = "EVENT_ITEM_MODIFY"
	//缓存项删除
	EVENT_ITEM_DELETE = "EVENT_ITEM_DELETE"
	//缓存表新增
	EVENT_TABLE_ADD = "EVENT_TABLE_ADD"
	//缓存表删除
	EVENT_TABLE_DELETE = "EVENT_TABLE_DELETE"
)

//请求命令
const (
	REQUEST_TYPE_PING    = "PING"    //心跳检测
	REQUEST_TYPE_CONNECT = "CONNECT" //连接
	REQUEST_TYPE_EXIT    = "EXIT"    //断开连接
	REQUEST_TYPE_DELETE  = "DELETE"  //删除
	REQUEST_TYPE_EXIST   = "EXIST"   //存在
	REQUEST_TYPE_EVENT   = "EVENT"   //事件
	REQUEST_TYPE_USE     = "USE"     //切换表
	REQUEST_TYPE_SHOWT   = "SHOWT"   //显示表信息
	REQUEST_TYPE_SHOWI   = "SHOWI"   //显示项信息
	REQUEST_TYPE_INFO    = "INFO"    //显示系统信息
	REQUEST_TYPE_HELP    = "HELP"    //帮助
	REQUEST_TYPE_BGSAVE  = "BGSAVE"  //后台保存gdb文件

	REQUEST_TYPE_SET    = "SET"    //添加string
	REQUEST_TYPE_GET    = "GET"    //获取string
	REQUEST_TYPE_APPEND = "APPEND" //追加string
	REQUEST_TYPE_STRLEN = "STRLEN" //值的长度string
	REQUEST_TYPE_SETNX  = "SETNX"  //不存在则设置string

	REQUEST_TYPE_NSET   = "NSET"   //添加number
	REQUEST_TYPE_NGET   = "NGET"   //获取number
	REQUEST_TYPE_INCR   = "INCR"   //增加1 number
	REQUEST_TYPE_INCRBY = "INCRBY" //增加指定值 number

)

//数据类型
const (
	DATA_TYPE_STRING = "string" //字符
	DATA_TYPE_BOOL   = "bool"   //布尔
	DATA_TYPE_NUMBER = "number" //数字
	DATA_TYPE_MAP    = "map"    //Map
	DATA_TYPE_SET    = "set"    //集合
	DATA_TYPE_LIST   = "list"   //列表
	DATA_TYPE_ZSET   = "zset"   //有序集合
	DATA_TYPE_OBJECT = "object" //对象
)

//协议类型
const (
	//默认使用终端
	PROTOCOL_RESPONSE_DEFAULT = "TERMINAL"
	//JSON，一般用于各语言客户端
	PROTOCOL_RESPONSE_JSON = "JSON"
	//终端，用于telnet等方式
	PROTOCOL_RESPONSE_TERMINAL = "TERMINAL"
)

//标识字符
const (
	FLAG_CHAR_SOCKET_COMMND_END            = '\n'
	FLAG_CHAR_SOCKET_TERMINAL_RESPONSE_END = "\r\n ->"
	FLAG_CHAR_SOCKET_JSON_RESPONSE_END     = "\r\n!--!>"
)

//帮助
const CONFIG_HELP_CONTENT_EN = `
[connect]
Desc=connect to the server
Format=connect [-t'table'] [-a'pwd'] [-i'ip'] [-p'port'] [-r'protocol'] [-e'e1,e2...']
[exit]
Desc=close connect and exit
Format=exit
[ping]
Desc=use for check the server is running
Format=ping
[help]
Desc=show commnd manual
Format=help
[set]
Desc=set key-value
Format=set key value [time]
[get]
Desc=get key-value
Format=get key
[delete]
Desc=delete key-value
Format=delete key
[exist]
Desc=check the key-value is exist
Format=exist key
[info]
Desc=show system info
Format=info
[use]
Desc=change cache table
Format=use [table name]
[showt]
Desc=show table info
Format=showt [table name]
[showi]
Desc=show item info
Format=showi 'item key'
`

//系统默认配置
const CONFIG_SYSTEM_DEFAULT = `
appname=gocache
version=1.0
author=luocihang@126.com

[server]
#server port
port=1334
#connect password
password=
#max size for connect pool
maxPoolSize=10
#core size for connect pool
corePoolSize=5
#connect type: long|short
connectType = long
#connect alive time, unit is second
aliveTime=3000
#system table use to cache connection info
sysTable=sys
#the commnds which could be execute without login
anonymCommnd=ping,connect,exit,help,info

[table]
#default table name. if client connect server without assign table name or 'openSession' is false, 
#use this name as default table name
default=default

[client]
#open session: true|false
openSession=true

[dump]
#use gdb dump: true|false
dump.on=true
#dump file path
filePath=./data/dump.gdb
#dump trigger: time1 update1,time2 update3,time3 update4,……
trigger=10 2,30 5

[log]
#log level: trace|debug|info|critical|warn|error
level = trace
#use console log: true|false
console.on=true
#log format
format=%Date %Time [%LEV] %Msg%n
#log file path
path=./tmp/dhaiy.log
#roll
roll=2006-01-02
`
