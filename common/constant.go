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
	DATA_TYPE_NIL    = "nil"    //空
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

const (
	GDB_GOCACHE         = "GOCACHE"
	GDB_VERSION         = "0001"
	GDB_DATABASE        = "DATABASE"
	GDB_TABLE           = "TABLE"
	GDB_EOF             = "EOF"
	GDB_LIVETIME_ALWAYS = "A"
)

const (
	//魔数长度
	GDB_LEN_GOCACHE = 7
	//版本号长度
	GDB_LEN_VERSION = 4
	//库标识长度
	GDB_LEN_DATABASE = 8
	//表标识长度
	GDB_LEN_TABLE = 5
	//结束符长度
	GDB_LEN_EOF = 3
	//校验码长度
	GDB_LEN_CHECK_SUM = 32
	//库大小长度
	GDB_LEN_DATABASE_SIZE = 4
	//表大小长度
	GDB_LEN_TABLE_SIZE = 8
	//键长度，包括表名、键名
	GDB_LEN_KEY = 3
	//值长度
	LEN_VALUE = 6
	//数据类型长度
	GDB_LEN_DATATYPE = 1
	//存活时间长度
	GDB_LEN_LIVETIME = 14
	//永久存活时间长度
	GDB_LEN_LIVETIME_ALWAYS = 1
)

//数据类型
const (
	GDB_TYPE_STRING = "1" //字符
	GDB_TYPE_BOOL   = "2" //布尔
	GDB_TYPE_NUMBER = "3" //数字
	GDB_TYPE_MAP    = "4" //Map
	GDB_TYPE_SET    = "5" //集合
	GDB_TYPE_LIST   = "6" //列表
	GDB_TYPE_ZSET   = "7" //有序集合
	GDB_TYPE_OBJECT = "8" //对象
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
