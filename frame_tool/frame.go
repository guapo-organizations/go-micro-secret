package frame_tool

import (
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/consul"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"github.com/spf13/viper"
	"log"
)

type LyFrameTool struct {
	//配置文件路径
	ConfigPath string
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
	//grpc服务信息
	//服务所在ip
	ip := viper.GetString("ip")
	//服务所在端口
	port := viper.GetString("port")
	//服务描述
	describe := viper.GetString("describe")
	//服务名字
	name := viper.GetString("name")
	//注册grpc服务
	service.CreateGrpcServiceInfo(ip, port, describe)

	//consul服务发现信息
	consul_start := viper.GetStringMap("consul")
	if consul_start["start"].(bool) {
		
		consul_info := viper.GetStringMapString("consul")
		//访问consul客户端的配置
		consul_config := consul.CreateConfig(consul_info["ip"], consul_info["port"])
		//异步去注册服务发现，如果失败，程序终止,checkPort是心跳检测的端口
		go consul.RegisterServer(consul_config, consul_info["checkPort"], name, ip, port, nil)
	}
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
		port := viper.GetString("gateway_port")
		grpc_service := service.GetGrpcServiceInfo()
		service.CreateGrpcGatewayServiceInfo(grpc_service, port)
	}
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
