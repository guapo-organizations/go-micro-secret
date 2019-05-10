package frame_tool

import (
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/database"
	"github.com/spf13/viper"
	"log"
)

type LyFrameTool struct {
	//配置文件路径
	ConfigPath string
}

//初始化数据库
func (this *LyFrameTool) initMysql() {
	viper.SetConfigName("db")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("数据库db配置文件加载失败", err)
	}
	ip := viper.GetString("ip")
	port := viper.GetString("port")
	user := viper.GetString("user")
	passwd := viper.GetString("passwd")
	db := viper.GetString("db")
	log.Println( fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, passwd, ip, port, db))
	database.CreateMysqlConnection(user, passwd, ip, port, db)
}

//从配置文件中获取端口
func (this *LyFrameTool) getConfigProt() string {
	viper.SetConfigName("service")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("grpc服务d配置文件加载失败", err)
	}
	port := viper.GetString("port")
	return port;
}

func (this *LyFrameTool) Run() (port string) {
	viper.AddConfigPath(this.ConfigPath)
	//初始化数据库
	this.initMysql()

	return this.getConfigProt()
}
