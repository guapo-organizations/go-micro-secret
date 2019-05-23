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
	//服务名字
	Name string
	//服务发现心跳检测的端口
	CheckPort string
	//consul服务发现的标记，可以用于consul的过滤查询
	Tags []string
}

func CreateGrpcServiceInfo(ip, port, describe, name, check_port string, tags []string) {
	grpc_service_info = ServiceInfo{
		Ip:        ip,
		Port:      port,
		Describe:  describe,
		Name:      name,
		CheckPort: check_port,
		Tags:      tags,
	}
}

func GetGrpcServiceInfo() ServiceInfo {
	return grpc_service_info
}
