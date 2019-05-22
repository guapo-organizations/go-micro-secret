package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

//consul过滤的语法参考这里
//https://www.consul.io/api/agent/service.html
func FindService(config *api.Config, service_id string) (*api.AgentService, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("consul 客户端FindService错误 :", err)
	}

	//找到服务ID 等于 server_id的服务
	services, err := client.Agent().ServicesWithFilter(service_id + " Equal ID")
	if err != nil {
		return nil, err
	}

	service, ok := services[service_id]

	if !ok {
		return nil, fmt.Errorf("找不到id为:%s的服务", service_id)
	}

	return service, nil
}
