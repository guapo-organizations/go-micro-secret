package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

var config *api.Config

//创建连接服务发现的config配置对象
func CreateConfig(ip string, port string) (config *api.Config) {
	config = api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", ip, port)
	return config
}

func GetConfig() (*api.Config, error) {
	if config == nil {
		return nil, fmt.Errorf("consul链接信息没有加载到缓存中")
	}
	return config, nil
}
