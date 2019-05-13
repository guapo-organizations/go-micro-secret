package service

var grpc_gateway_service_info GatewayServiceInfo
//grpc网关服务的连接信息

type GatewayServiceInfo struct {
	ServiceInfo
	GatewayPort string
}

//创建grpc网关
func CreateGrpcGatewayServiceInfo(grpc_service_info ServiceInfo, gateway_port string) {
	grpc_gateway_service_info = GatewayServiceInfo{
		ServiceInfo: ServiceInfo{
			Ip:       grpc_service_info.Ip,
			Port:     grpc_service_info.Port,
			Describe: grpc_service_info.Describe,
		},
		GatewayPort: gateway_port,
	}
}


func GetGrpcGateWayServiceInfo() GatewayServiceInfo {
	return grpc_gateway_service_info
}
