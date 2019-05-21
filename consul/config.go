package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

//创建连接服务发现的config配置对象
func CreateConfig(ip string, port int64) (config *api.Config) {
	config = api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", ip, port)
	return config
}
