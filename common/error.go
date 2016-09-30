package common

import (
	"errors"
)

//异常
var (
	ERROR_TABLE_NAME               = errors.New("ERROR_TABLE_NAME")
	ERROR_SERVER_ALREADY_START     = errors.New("server already start")
	ERROR_SERVER_CONNECT_TYPE      = errors.New("connect type is not support")
	ERROR_COMMAND_NOT_FOUND        = errors.New("command not found")
	ERROR_COMMAND_PARAM_ERROR      = errors.New("command param error")
	ERROR_COMMAND_NO_LOGIN         = errors.New("you have no connect")
	ERROR_ITEM_NOT_EXIST           = errors.New("item not exist")
	ERROR_TABLE_NOT_EXIST          = errors.New("table not exist")
	ERROR_AUTHORITY_NO_PWD         = errors.New("system use password")
	ERROR_AUTHORITY_PWD_ERROR      = errors.New("password error")
	ERROR_PORT_ERROR               = errors.New("port error")
	ERROR_PROTOCOL_ERROR           = errors.New("protocol error")
	ERROR_SYSTEM                   = errors.New("system error")
	ERROR_COMMAND_NOT_SUPPORT_DATA = errors.New("command not support for the data type")

	GDB_FILE_INVALID       = errors.New("invalid gdb file")
	GDB_FILE_VERSION_ERROR = errors.New("gdb file version is not support")
	GDB_FILE_CHECK_ERROR   = errors.New("gdb file is broken")
	GDB_FILE_FORMAT_ERROR  = errors.New("error gdb file format")
)
