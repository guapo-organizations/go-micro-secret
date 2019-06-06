package frame_tool

import (
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/consul"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"github.com/spf13/viper"
	"log"
	"strings"
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
		passwd := viper.GetString("passwd")
		//这个方法可以手动连接redis
		cache.CreateRedisConnection(ip, port, passwd, db)
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
	//服务发现的心跳检测端口
	check_port := viper.GetString("checkPort")
	//服务发现的tags，用户过滤之类的
	tags := viper.GetString("tags")
	//注册grpc服务
	service.CreateGrpcServiceInfo(ip, port, describe, name, check_port, strings.Split(tags, " "))

}

//consul服务发现注册
func (this *LyFrameTool) ininConsul() {
	viper.SetConfigName("consul")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("consul服务发现配置文件加载失败", err)
	}
	//是否开启服务发现
	start := viper.GetBool("start")
	if start == true {
		//服务发现连接地址
		ip := viper.GetString("ip")
		//服务发现连接端口
		port := viper.GetString("port")

		//创建连接服务发现的配置
		consul.CreateConfig(ip, port)
		consul_config, err := consul.GetConfig()
		if err != nil {
			log.Fatalln("consul配置获取不到")
		}
		//异步去注册服务发现，如果失败，程序终止,checkPort是心跳检测的端口
		go consul.RegisterServer(consul_config, service.GetGrpcServiceInfo())
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
	//初始化服务发现
	this.ininConsul()
}
