package common

import (
	"encoding/json"
	"strconv"

	"github.com/louch2010/dhaiy/log"
)

//响应信息
type CmdResponse struct {
	Code     string      //响应码
	Data     interface{} //响应数据
	DataType string      //数据类型
	Clo      bool        //是否关闭连接
	Err      error       //错误信息
	Client   *Client     //客户端对象
}

//JSON响应
type JsonCmdResponse struct {
	Code     string
	Msg      string
	Data     interface{}
	DataType string //数据类型
}

func GetCmdResponse(code string, data interface{}, err error, c *Client) *CmdResponse {
	resp := CmdResponse{
		Code:     code,
		Data:     data,
		DataType: DATA_TYPE_STRING,
		Err:      err,
		Clo:      false,
		Client:   c,
	}
	return &resp
}

//根据连接协议，将响应内容进行封装
func TransferResponse(response *CmdResponse) string {
	protocol := ""
	if response.Client != nil {
		protocol = response.Client.Protocol
	}
	//终端方式：有错误，则输出错误信息，没有错误，则直接输出响应信息
	if protocol == "" || protocol == PROTOCOL_RESPONSE_TERMINAL {
		if response.Err != nil {
			return response.Err.Error() + FLAG_CHAR_SOCKET_TERMINAL_RESPONSE_END
		}
		if response.DataType == DATA_TYPE_NIL {
			return "nil" + FLAG_CHAR_SOCKET_TERMINAL_RESPONSE_END
		}
		ret := toString(response.Data)
		if ret == "" {
			ret = "OK"
		}
		return ret + FLAG_CHAR_SOCKET_TERMINAL_RESPONSE_END
	}
	//JSON方式：对响应信息进行json封装
	if protocol == PROTOCOL_RESPONSE_JSON {
		msg := MESSAGE_SUCCESS
		if response.Err != nil {
			msg = response.Err.Error()
		}
		obj := JsonCmdResponse{
			Code:     response.Code,
			Msg:      msg,
			Data:     response.Data,
			DataType: response.DataType,
		}
		j, _ := json.Marshal(obj)
		return string(j) + FLAG_CHAR_SOCKET_JSON_RESPONSE_END
	}
	return ""
}

//转string
func toString(v interface{}) string {
	response := ""
	switch conv := v.(type) {
	case string:
		response = conv
		break
	case int:
		response = strconv.Itoa(conv)
		break
	case bool:
		response = strconv.FormatBool(conv)
		break
	case float64:
		response = strconv.FormatFloat(conv, 'E', -1, 64)
		break
	case []string:
		for _, v := range conv {
			response += v + "\r\n"
		}
		break
	case *CacheItem:
		if conv != nil {
			tmp, _ := json.Marshal(conv.Value())
			response = string(tmp)
		}
		break
	default:
		log.Error("类型转换异常")
	}
	return response
}
