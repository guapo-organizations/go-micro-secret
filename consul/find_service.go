package consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
)

//consul过滤的语法参考这里
//https://www.consul.io/api/agent/service.html
func FindService(config *api.Config, service_name, tag string) (*api.AgentService, error) {
	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("consul 客户端FindService错误 :%s", err)
	}

	var services map[string]*api.AgentService
	if tag == "" {
		//查询所有的
		services, err = client.Agent().Services()
	} else {
		//过滤查询
		services, err = client.Agent().ServicesWithFilter(tag + " in Tags")
	}

	if err != nil {
		return nil, err
	}

	//通过id去找service
	service, ok := services[CreateAgentServiceUniqueID(service_name)]

	if !ok {
		return nil, fmt.Errorf("找不到%s的服务", service_name)
	}

	return service, nil
}
