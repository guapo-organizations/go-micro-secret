package service

//grpc网关服务的连接信息

type GatewayServiceInfo struct {
	//服务所在的地址
	Ip string
	//服务所在端口
	Port string
	//服务描述
	Describe string
}
