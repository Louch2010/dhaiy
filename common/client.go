package common

//客户端
type Client struct {
	ServerConfig *ServerConfig                    //服务器配置
	Host         string                           //地址
	Port         int                              //端口
	Table        string                           //表名
	CacheTable   *CacheTable                      //表指针
	ListenEvent  []string                         //侦听事件
	Protocol     string                           //通讯协议
	Token        string                           //令牌
	Reqest       []string                         //请求参数
	ReqestString string                           //请求命令原文
	Response     *CmdResponse                     //响应信息
	IsLogin      bool                             //是否登录
	Handler      func(client *Client) CmdResponse //处理函数
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

//服务器信息
type ServerInfo struct {
	ServerId     string   //服务器唯一标识
	Host         string   //服务器IP、端口
	Port         int      //端口号
	Status       int      //状态：1首次同步中、2同步中、3同步异常
	SyncPosition int      //同步偏移量
	CmdQueue     []string //命令集合
}
