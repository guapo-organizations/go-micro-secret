package service

var grpc_gateway_service_info GatewayServiceInfo
//grpc网关服务的连接信息

type GatewayServiceInfo struct {
	//服务所在的地址
	Ip string
	//服务所在端口
	Port string
	//服务描述
	Describe string
}

func CreateGrpcGatewayServiceInfo(ip, port, describe string) {
	grpc_gateway_service_info = GatewayServiceInfo{
		Ip:       ip,
		Port:     port,
		Describe: describe,
	}
}

func GetGrpcGateWayServiceInfo() GatewayServiceInfo {
	return grpc_gateway_service_info
}
