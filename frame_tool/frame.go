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
	viper.AddConfigPath(this.ConfigPath)
	viper.SetConfigName("db")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("数据库配置文件加载失败", err)
	}
	user := viper.GetString("user")
	passwd := viper.GetString("passwd")
	db := viper.GetString("db")
	log.Println(fmt.Sprintf("连接数据库的信息;user:%s,passwd:%s,db:%s", user, passwd, db))
	database.CreateMysqlConnection(user, passwd, db)
}

//从配置文件中获取端口
func (this *LyFrameTool) getConfigProt() string {
	viper.AddConfigPath(this.ConfigPath)
	return "";
}

func (this *LyFrameTool) Run() (port string) {

	//初始化数据库
	this.initMysql()

	return this.getConfigProt()
}
