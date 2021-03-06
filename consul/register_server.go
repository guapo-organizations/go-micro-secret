package consul

import (
	"fmt"
	"github.com/guapo-organizations/go-micro-secret/frame_tool/service"
	"github.com/hashicorp/consul/api"
	"log"
	"net/http"
	"strconv"
)

//心跳检测逻辑
func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consul心跳检测")
}

//注册一个服务
//config 是连接服务发现的配置
//checkPort是服务发现心跳检测的端口
//name是服务的名字
//address 服务的ip地址
//port服务的端口
//tags 服务标记用于过滤用的，很重要的
func RegisterServer(config *api.Config, service_info service.ServiceInfo) {
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatalln("consul 客户端RegisterService错误 :", err)
	}

	//服务
	registeration := new(api.AgentServiceRegistration)
	registeration.ID = CreateAgentServiceUniqueID(service_info.Name)
	registeration.Name = service_info.Name
	port64, _ := strconv.ParseInt(service_info.Port, 10, 64)
	registeration.Port = int(port64)
	registeration.Tags = service_info.Tags
	registeration.Address = service_info.Ip
	registeration.Check = &api.AgentServiceCheck{
		HTTP:                           fmt.Sprintf("http://%s:%s%s", registeration.Address, service_info.CheckPort, "/check"),
		Timeout:                        "6s",
		Interval:                       "40s",
		DeregisterCriticalServiceAfter: "120s", //check失败后120秒删除本服务
	}

	err = client.Agent().ServiceRegister(registeration)

	if err != nil {
		log.Fatalln("注册服务失败:", err)
	}

	http.HandleFunc("/check", consulCheck)
	http.ListenAndServe(fmt.Sprintf(":%s", service_info.CheckPort), nil)
}
