package frame_tool

import (
	"github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type LyGatewayFrameTool struct {
	//配置文件路径
	ConfigPath string
}

//解析grpc服务信息
func (this *LyGatewayFrameTool) initGrpcServiceInfo() {
	viper.SetConfigName("grpc_service")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("grpc服务配置文件加载失败", err)
	}
	//grpc服务信息
	//服务所在ip
	ip := viper.GetString("ip")
	//服务所在端口
	port := viper.GetString("port")
	//服务描述
	describe := viper.GetString("describe")
	//服务名字
	name := viper.GetString("name")
	//服务发现的心跳检测端口
	check_port := viper.GetString("checkPort")
	//服务发现的tags，用户过滤之类的
	tags := viper.GetString("tags")
	//注册grpc服务
	service.CreateGrpcServiceInfo(ip, port, describe, name, check_port, strings.Split(tags, " "))

}

//解析grpc网关服务信息
func (this *LyGatewayFrameTool) initGrpcGatewayServiceInfo() {
	viper.SetConfigName("grpc_gateway_service")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("grpc 网关服务配置文件加载失败", err)
	}
	start := viper.GetBool("start")
	if start == true {
		//开启redis
		port := viper.GetString("gateway_port")
		grpc_service := service.GetGrpcServiceInfo()
		service.CreateGrpcGatewayServiceInfo(grpc_service, port)
	}
}

func (this *LyGatewayFrameTool) Run() {
	viper.AddConfigPath(this.ConfigPath)
	//初始化服务发现
	this.initGrpcServiceInfo()
	//初始化grpc网关服务
	this.initGrpcGatewayServiceInfo()
}
