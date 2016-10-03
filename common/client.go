package common

//客户端
type Client struct {
	Host        string                             //地址
	Port        int                                //端口
	Table       string                             //表名
	CacheTable  *CacheTable                        //表指针
	ListenEvent []string                           //侦听事件
	Protocol    string                             //通讯协议
	Token       string                             //令牌
	Reqest      []string                           //请求参数
	Response    *CmdResponse                     //响应信息
	IsLogin     bool                               //是否登录
	Handler     func(client *Client) CmdResponse //处理函数
}

//命令
type Cmd struct {
	Name        string               //命令名
	MinParam    int                  //最小参数个数
	MaxParam    int                  //最大参数个数
	InvokeCount int64                //调用次数
	IsWrite     bool                 //是否为写命令
	HandlerFunc func(client *Client) //处理函数
}
