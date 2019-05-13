package frame_tool

import (
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"github.com/spf13/viper"
	"log"
)

type LyFrameTool struct {
	//配置文件路径
	ConfigPath string

	//grpc服务器连接信息
	grpc_service_info *service.ServiceInfo

	//grpc网关服务器连接信息
	grpc_gateway_service_info *service.GatewayServiceInfo
}

//初始化数据库,只是尝试的开启而已，如果配置false就不开启，但是如果服务里面想手动连接数据库的话，可以手动调用连接数据库
func (this *LyFrameTool) initMysql() {
	viper.SetConfigName("db")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("数据库db配置文件加载失败", err)
	}
	start := viper.GetBool("start")

	if start == true {
		//开启数据库
		ip := viper.GetString("ip")
		port := viper.GetString("port")
		user := viper.GetString("user")
		passwd := viper.GetString("passwd")
		db := viper.GetString("db")
		//这个方法可以让服务手动连接数据库
		database.CreateMysqlConnection(user, passwd, ip, port, db)
	}
}

//初始化redis，只是尝试开启redis
func (this *LyFrameTool) initRedis() {
	viper.SetConfigName("redis")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("redis配置文件加载失败", err)
	}

	start := viper.GetBool("start")

	if start == true {
		//开启redis
		ip := viper.GetString("ip")
		port := viper.GetString("port")
		db := viper.GetInt("db")
		//这个方法可以手动连接redis
		cache.CreateRedisConnection(ip, port, db)
	}
}

//解析grpc服务信息
func (this *LyFrameTool) initGrpcServiceInfo() {
	viper.SetConfigName("grpc_service")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("grpc服务配置文件加载失败", err)
	}
	ip := viper.GetString("ip")
	port := viper.GetString("port")
	describe := viper.GetString("describe")
	this.grpc_service_info = &service.ServiceInfo{
		Ip:       ip,
		Port:     port,
		Describe: describe,
	}
}

//获取grpc服务信息
func (this *LyFrameTool) GetGrpcServiceInfo() *service.ServiceInfo {
	if this.grpc_service_info == nil {
		this.Run()
	}
	return this.grpc_service_info
}

//解析grpc网关服务信息
func (this *LyFrameTool) initGrpcGatewayServiceInfo() {
	viper.SetConfigName("grpc_gateway_service")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("grpc 网关服务配置文件加载失败", err)
	}
	start := viper.GetBool("start")
	if start == true {
		//开启redis
		ip := viper.GetString("ip")
		port := viper.GetString("port")
		describe := viper.GetString("describe")
		this.grpc_gateway_service_info = &service.GatewayServiceInfo{
			Ip:       ip,
			Port:     port,
			Describe: describe,
		}
	}
}

//获取grpc网关信息
func (this *LyFrameTool) GetGrpcGatewayServiceInfo() *service.GatewayServiceInfo {
	if this.grpc_gateway_service_info == nil {
		this.Run()
	}
	return this.grpc_gateway_service_info
}

//初始化一下服务需要用的中间件
func (this *LyFrameTool) Run() {
	viper.AddConfigPath(this.ConfigPath)
	//初始化数据库
	this.initMysql()
	//初始化redis
	this.initRedis()

	//初始化grpc服务
	this.initGrpcServiceInfo()

	//初始化grpc网关服务
	this.initGrpcGatewayServiceInfo()
}
