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
	Response    *ServerRespMsg                     //响应信息
	IsLogin     bool                               //是否登录
	Handler     func(client *Client) ServerRespMsg //处理函数
}
