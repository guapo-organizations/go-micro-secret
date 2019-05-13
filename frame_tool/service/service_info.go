package service

var grpc_service_info ServiceInfo
//grpc服务的连接信息
type ServiceInfo struct {
	//服务所在的地址
	Ip string
	//服务所在端口
	Port string
	//服务描述
	Describe string
}

func CreateGrpcServiceInfo(ip, port, describe string) {
	grpc_service_info = ServiceInfo{
		Ip:       ip,
		Port:     port,
		Describe: describe,
	}
}


func GetGrpcServiceInfo() ServiceInfo {
	return grpc_service_info
}
